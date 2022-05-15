// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package logging

import (
	"testing"

	"github.com/go-kit/log/level"
	"github.com/stretchr/testify/assert"
)

func TestNewLoggerText(t *testing.T) {
	logger := NewLogger("logfmt", "info")

	err := level.Error(logger).Log("msg", "TestNewLogger-INFO")
	assert.Nil(t, err)
}

func TestNewLoggerJson(t *testing.T) {
	logger := NewLogger("json", "info")

	err := level.Error(logger).Log("msg", "TestNewLogger-INFO")
	assert.Nil(t, err)
}
