# brew-deps

Display homebrew dependencies as a tree

## Usage

```bash
$ brew-deps
.
├── ack
├── automake
│   └── autoconf
├── bat
├── cairo
│   ├── fontconfig
│   │   └── freetype
│   │       └── libpng
│   ├── freetype
│   │   └── libpng
│   ├── glib
│   │   ├── gettext
│   │   ├── libffi
│   │   ├── pcre
│   │   └── python
│   │       ├── gdbm
│   │       ├── openssl@1.1
│   │       ├── readline
│   │       ├── sqlite
│   │       │   └── readline
│   │       └── xz
│   ├── libpng
│   ├── lzo
│   └── pixman
├── curl
.
.
.
└── zlib
```

## Installation

```bash
$ go get github.com/quells/brew-deps
$ go install github.com/quells/brew-deps
```
