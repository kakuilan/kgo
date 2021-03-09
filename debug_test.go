package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebug_DumpPrint(t *testing.T) {
	defer func() {
		r := recover()
		assert.Empty(t, r)
	}()

	KDbug.DumpPrint(Version)
}

func TestDebug_DumpStacks(t *testing.T) {
	defer func() {
		r := recover()
		assert.Empty(t, r)
	}()

	KDbug.DumpStacks()
}
