# Go bin-unbin
[![Commits since last release](https://img.shields.io/github/commits-since/nico-castell/go-binme/latest?label=Commits%20since%20last%20release&color=informational&logo=Git&logoColor=white&style=flat-square)](https://github.com/nico-castell/go-binme/commits)
[![Release](https://img.shields.io/github/v/release/nico-castell/go-binme?label=Release&color=informational&logo=GitHub&logoColor=white&style=flat-square)](https://github.com/nico-castell/go-binme/releases)
[![License](https://img.shields.io/github/license/nico-castell/go-binme?label=License&color=informational&logo=Open%20Source%20Initiative&logoColor=white&style=flat-square)](LICENSE)
[![Lines of code](https://img.shields.io/tokei/lines/github/nico-castell/go-binme?label=Lines%20of%20code&color=informational&logo=Go&logoColor=white&style=flat-square)](https://github.com/nico-castell/go-binme)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/nico-castell/go-binme/CodeQL?label=CodeQL&logo=GitHub%20Actions&logoColor=white&style=flat-square)](https://github.com/nico-castell/go-binme/actions/workflows/codeql-analysis.yml)

I've been learning [***Golang***](https://golang.org/) for a while, I made this program to process input into binary, hexadecimal and octal.

Usage:
- To process text input into binary:
    ```shell
    $ go-binme
    ```
- To process text input into hexadecimal:
    ```shell
    $ go-binme -x
    ```
- To process text input into octal:
    ```shell
    $ go-binme -o
    ```
- You can pipe the input:
    ```shell
    $ echo "Hello, World" | go-binme
    ```

The program will print a new line after it's done, but only if the output is a terminal, not a pipe.
