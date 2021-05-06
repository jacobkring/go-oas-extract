# go-oas-extract

go-oas-extract is a tool for extracting specially tagged comments in Go
source code and transforming it into valid Open API 3.0.3 Specification.

Tagged comments start with a blank line containing `+extract`.  

Both grouped line comments (`//`) and block comments (`/* */`) are supported.

## Installation

Go 1.16:

    go install github.com/jacobkring/go-oas-extract@latest

## Usage

Simply tag comments in your Go source code with `+extract` as the first
line of your comment.  For example,

```go
package main

// +extract
// API Documentation
// =================
//
// (TODO: add documentation)
```

Then run the `doc-extract` command, providing it with a directory of
Go source files and an output file:

    doc-extract ./example example.txt

Source files are processed in lexicographic order, _except_ that a file
named `doc.go` gets highest priority.  Comments within a file are
processed in the order they appear.

This predictable ordering allows you to add, for instance, a header to
the output file by adding it to `doc.go`.

## Example

The `example` directory contains an [example using API
Blueprint](example/README.md).

## Motivation

There weren't any existing libraries that I could find that would extract 
inline comments from a Go API and transform them into a valid Open API 3.0+ 
specification. 

## Contributing

Issues and pull requests are welcome.  When filing a PR, please make
sure the code has been run through `gofmt` and that the tests pass.

## Notes

Initially forked from https://github.com/joeshaw/doc-extract to provide a basis for text extraction.

## License

MIT
