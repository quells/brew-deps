package main

import "testing"

func Test_getInstalledPkgs(t *testing.T) {
	installed, err := getInstalledPkgs()
	if err != nil {
		t.Fatal(err)
	}
	if len(installed) == 0 {
		t.Fatal("no installed packages found")
	}
}
