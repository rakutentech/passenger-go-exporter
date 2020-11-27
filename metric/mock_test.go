// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package metric

import (
	"errors"

	"ghe.rakuten-it.com/SOK/passenger-go-exporter/passenger"
)

// MockServer is mock that gets metrics.
type MockServer struct {
}

func (s MockServer) Metrics() (*passenger.PoolInfo, error) {
	pool := &passenger.PoolInfo{
		SuperGroups: []passenger.SuperGroup{
			{
				Group: passenger.Group{
					Name: "a/b/c",
					Processes: []passenger.Process{
						{
							PID: "1",
						},
					},
				},
			},
		},
	}
	return pool, nil
}

func (s MockServer) IsEnabled() bool {
	return true
}
func (s MockServer) Configuration(homedir string) bool {
	return true
}

// ErrorMockServer is mock causes an error.
type ErrorMockServer struct {
}

func (s ErrorMockServer) Metrics() (*passenger.PoolInfo, error) {
	return nil, errors.New("mock error")
}

func (s ErrorMockServer) IsEnabled() bool {
	return true
}
