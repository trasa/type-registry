package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetermineFileName_withOverride(t *testing.T) {
	assert.Equal(t, "override.go", determineFileName("notused", "override.go"))
	assert.Equal(t, "override", determineFileName("notused", "override"))
}

func TestDetermineFileName_WithGoExtension(t *testing.T) {
	assert.Equal(t, "myinput.pb.typeregistry.go", determineFileName("myinput.pb.go", ""))
}

func TestDetermineFileName_WithoutExtension(t *testing.T) {
	assert.Equal(t, "myinput.pb.typeregistry.go", determineFileName("myinput.pb", ""))
}

func TestDetermineFileName_WithPath(t *testing.T) {
	assert.Equal(t, "src/to/somewhere/myinput.pb.typeregistry.go", determineFileName("src/to/somewhere/myinput.pb.go", ""))
}
