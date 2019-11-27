/*
Copyright 2019 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package db

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMembership(t *testing.T) {
	topic, closer := makeTopic()
	defer closer()

	for n := uint64(0); n <= 64; n++ {
		key := []byte(strconv.FormatUint(n, 10))
		index, err := topic.Set(key, key)
		assert.NoError(t, err, "n=%d", n)
		assert.Equal(t, n, index, "n=%d", n)
	}

	index := uint64(5)
	at := uint64(64)

	topic.store.WaitUntil(at)

	proof, err := topic.MembershipProof(index)
	assert.NoError(t, err)
	assert.Equal(t, proof.Index, index)
	assert.Equal(t, proof.At, at)
	assert.Equal(t, proof.Root, root64th)
	assert.Equal(t, proof.Hash, *topic.store.Get(0, index))
	assert.True(t, proof.Verify())
}
