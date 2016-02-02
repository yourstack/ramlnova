// Package main provides types
package main

import (
	"github.com/buddhamagnet/raml"
)

type Template struct {
	Uid     string
	Lang    string
	Path    string
	Name    string
	Content string
}

// Route represents a Resourse in RAML.
type Route struct {
	Uri           string
	UriParameters map[string]raml.NamedParameter
	Methods       []*raml.Method
	Description   string
	DisplayName   string
}

// ControllerInfo contains controller information.
// RAML's Advance Resourse.Method()
type ControllerInfo struct {
	Name, Verb, Path, Doc string
}
