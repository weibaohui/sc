// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHook(t *testing.T) {
	path := tempPath()
	h := &Hook{
		name:     HookPreReceive,
		path:     path,
		isSample: false,
		content:  "test content",
	}

	assert.Equal(t, HookPreReceive, h.Name())
	assert.Equal(t, path, h.Path())
	assert.False(t, h.IsSample())
	assert.Equal(t, "test content", h.Content())
}

func TestHook_Update(t *testing.T) {
	path := tempPath()
	defer func() {
		_ = os.Remove(path)
	}()

	h := &Hook{
		name:     HookPreReceive,
		path:     path,
		isSample: false,
	}
	err := h.Update("test content")
	if err != nil {
		t.Fatal(err)
	}

	p, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "test content", string(p))
}
