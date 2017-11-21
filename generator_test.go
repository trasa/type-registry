package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyPrefix(t *testing.T) {
	g := NewGenerator()
	g.typenamePrefix = "zzz"
	g.typeNames["zzzfoo"] = true

	assert.Equal(t, "zzzfoo", g.transformTypeName("foo"))
}

func TestApplyPrefix_AlreadyThere(t *testing.T) {
	g := NewGenerator()
	g.typenamePrefix = "zzz"
	g.typeNames["zzzfoo"] = true

	assert.Equal(t, "zzzfoo", g.transformTypeName("zzzfoo"))
}

func TestApplyPrefix_CreatesInvalidType(t *testing.T) {
	g := NewGenerator()
	g.typenamePrefix = "zzz"
	//	g.typeNames["zzzfoo"] = true

	assert.Equal(t, "foo", g.transformTypeName("foo"))
}
