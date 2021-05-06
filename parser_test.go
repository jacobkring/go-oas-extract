package main

import (
	"github.com/jacobkring/go-assert"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"testing"
)

func TestSorter(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "testdata", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs) != 1 {
		t.Fatalf("Expected len(pkgs) == 1, got %d", len(pkgs))
	}

	expected := []string{"doc.go", "a.go", "paths.go", "testdata.go", "z.go"}
	sorted := sortedFiles(pkgs["testdata"])

	if len(sorted) != len(expected) {
		t.Fatalf("Expected %d, got %d", len(expected), len(sorted))
	}

	for i := 0; i < len(sorted); i++ {
		if sorted[i].name != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], sorted[i].name)
		}
	}
}

func TestExtractComment(t *testing.T) {
	tests := []struct {
		comment  string
		expected bool
	}{
		{"//+extract", true},
		{"// +extract", true},
		{"//           +extract\n", true},
		{"/*+extract*/", true},
		{"/* +extract */", true},
		{"/*           +extract */", true},
		{"/*\n+extract */", true},
		{"/* +extract\nfoo */", true},
		{"/* foo\n+extract */", false},
		{"// extract", false},
		{"// foo +extract", false},
	}

	for i, test := range tests {
		var c ast.Comment
		c.Text = test.comment

		cgrp := &ast.CommentGroup{List: []*ast.Comment{&c}}
		_, e := extractComment(cgrp)
		if e.Valid() != test.expected {
			t.Fatalf("Iteration %d: Expected %t, got %t: %q", i, test.expected, e.Valid(), test.comment)
		}
	}
}

func TestParsePackage(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "testdata", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs) != 1 {
		t.Fatalf("Expected len(pkgs) == 1, got %d", len(pkgs))
	}

	comments := extractPackageComments(pkgs["testdata"])

	file, err := os.Open("testoutput/doc.yml")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	expected := []string{
		string(b),
	}


	if len(comments) != len(expected) {
		t.Fatalf("Expected %#v, got %#v", expected, comments)
	}

	for i := 0; i < len(comments); i++ {
		assert.Equal(t, expected[i], comments[i])
	}
}
