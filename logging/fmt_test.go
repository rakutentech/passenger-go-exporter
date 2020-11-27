// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package logging

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsg(t *testing.T) {
	msg := "This is test message."
	key, val := Msg(msg)

	assert.Equal(t, MessageLabel, key)
	assert.Equal(t, msg, val)
}

func TestMsgf(t *testing.T) {
	fmt := "key=%s, code=%d"
	key, val := Msgf(fmt, "OK", 200)

	assert.Equal(t, MessageLabel, key)
	assert.Equal(t, "key=OK, code=200", val)
}

func TestErr(t *testing.T) {
	err := errors.New("This is Unit Test")
	key, val := Err(&err)

	assert.Equal(t, MessageLabel, key)
	assert.Equal(t, "This is Unit Test", val)
}
