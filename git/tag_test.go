// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTag(t *testing.T) {
	tag, err := testrepo.Tag("v1.1.0")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, ObjectTag, tag.Type())
	assert.Equal(t, "b39c8508bbc4b00ad2e24d358012ea123bcafd8d", tag.ID().String())
	assert.Equal(t, "0eedd79eba4394bbef888c804e899731644367fe", tag.CommitID().String())
	assert.Equal(t, "refs/tags/v1.1.0", tag.Refspec())

	t.Run("Tagger", func(t *testing.T) {
		assert.Equal(t, "Joe Chen", tag.Tagger().Name)
		assert.Equal(t, "joe@sourcegraph.com", tag.Tagger().Email)
		assert.Equal(t, int64(1581602099), tag.Tagger().When.Unix())
	})

	assert.Equal(t, "The version 1.1.0\n", tag.Message())
}

func TestTag_Commit(t *testing.T) {
	tag, err := testrepo.Tag("v1.1.0")
	if err != nil {
		t.Fatal(err)
	}

	c, err := tag.Commit()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "0eedd79eba4394bbef888c804e899731644367fe", c.ID.String())
}
