package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/xlab/treeprint"
)

const (
	formulaeURL = "https://formulae.brew.sh/api/formula.json"
)

var (
	deps map[string][]string
)

func main() {
	var formulae []Formula
	var err error
	if len(os.Args) == 2 {
		formulae, err = readFormulae(os.Args[1])
	} else {
		formulae, err = getFormulae()
	}
	if err != nil {
		fatal(err)
	}

	deps = make(map[string][]string)
	for _, f := range formulae {
		if len(f.Deps) > 0 {
			deps[f.Name] = f.Deps
		}
	}

	var pkgs []string
	pkgs, err = getInstalledPkgs()
	if err != nil {
		fatal(err)
	}

	tree := treeprint.New()
	for _, pkg := range pkgs {
		plant(tree, pkg)
	}
	fmt.Println(tree.String())
}

func plant(tree treeprint.Tree, pkg string) {
	branch := tree.AddBranch(pkg)
	for _, dep := range deps[pkg] {
		plant(branch, dep)
	}
}

// A Formula for a homebrew package
type Formula struct {
	Name string   `json:"name"`
	Deps []string `json:"dependencies"`
}

func getFormulae() (formulae []Formula, err error) {
	var resp *http.Response
	resp, err = http.Get(formulaeURL)
	if err != nil {
		return
	}

	var respData []byte
	respData, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	err = json.Unmarshal(respData, &formulae)
	return
}

func readFormulae(filename string) (formulae []Formula, err error) {
	var fileData []byte
	fileData, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileData, &formulae)
	return
}

func getInstalledPkgs() (installed []string, err error) {
	cmd := exec.Command("brew", "list", "--formula")

	var stdout io.ReadCloser
	stdout, err = cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		return
	}

	var stderr io.ReadCloser
	stderr, err = cmd.StderrPipe()
	defer stderr.Close()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	var errResp []byte
	errResp, err = ioutil.ReadAll(stderr)
	if err != nil {
		return
	}
	if len(errResp) != 0 {
		err = fmt.Errorf(string(errResp))
		return
	}

	var resp []byte
	resp, err = ioutil.ReadAll(stdout)
	if err != nil {
		return
	}

	installed = strings.Split(strings.TrimRight(string(resp), "\n"), "\n")
	return
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
}

func fatal(err error) {
	printErr(err)
	os.Exit(1)
}
