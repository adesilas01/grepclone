
# Mini Grep CLI

A small command-line search tool written in Go, inspired by grep.

This project searches text files for matching content and supports a few common search flags.

## Features

* `-i` — case-insensitive search
* `-n` — show line numbers
* `-c` — print only the number of matches
* `-w` — match whole words only

## Usage

```bash
go run . [query] [file] [flags]
```

### Example

```bash
go run . hello file.txt -i -n
```

## Example Input

`file.txt`

```text
Hello world
golang is fun
say hello again
```

## Example Output

```text
1: Hello world
3: say hello again
```

## Why I built this

I built this project to practice:

* command-line argument parsing
* working with files
* strings and slices
* building small but useful CLI tools in Go

## Future Improvements

Possible next steps:

* support regular expressions
* highlight matched text
* read from standard input
* support writing output to a file