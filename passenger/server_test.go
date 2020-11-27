// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package passenger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFactory(t *testing.T) {
	c := Context{}
	o := CreateFactory(c)
	f, ok := o.(*Passenger6ServerFactory)
	assert.True(t, ok)
	assert.NotNil(t, f)
}
