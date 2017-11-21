package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"sort"
	"strings"
)

// The state of the run.
type Generator struct {
	buf            bytes.Buffer //output
	packageName    string
	innerFieldName string
	typeNames      map[string]bool
	typenamePrefix string
}

func NewGenerator() *Generator {
	return &Generator{
		typeNames: make(map[string]bool),
	}
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// Apply the transform prefix to the type name if it doesn't already
// begin with the type name prefix, AND only apply that transform
// if it results in a valid type name from our typeNames collection.
func (g *Generator) transformTypeName(typeName string) string {
	if strings.HasPrefix(typeName, g.typenamePrefix) {
		return typeName
	}

	// try and apply the prefix, does it give us something valid?
	s := fmt.Sprintf("%s%s", g.typenamePrefix, typeName)
	if _, ok := g.typeNames[s]; ok {
		return s
	}
	// typenameprefix + typename isn't a known type, so just return the basic type
	return typeName
}

func (g *Generator) SortedTypeNames() []string {
	var typeNames []string
	for t := range g.typeNames {
		typeNames = append(typeNames, t)
	}
	sort.Strings(typeNames)
	return typeNames
}
