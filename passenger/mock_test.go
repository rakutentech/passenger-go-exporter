// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package passenger

// MockServer is mock that gets metrics.
type MockServer struct {
	isEnabled bool
}

// MockServer implements the passenger.ServerFactory interface.
type MockServerFactory struct {
	isEnabled bool
}

func (s MockServer) Metrics() (*PoolInfo, error) {
	pool := &PoolInfo{
		SuperGroups: []SuperGroup{
			{
				Group: Group{
					Name: "a/b/c",
					Processes: []Process{
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
	return s.isEnabled
}

// NewInstance implements the passenger.ServerFactory interface.
func (f MockServerFactory) NewInstance(homedir string) Server {
	return MockServer{isEnabled: f.isEnabled}
}

// SearchInstance implements the passenger.ServerFactory interface.
func (f MockServerFactory) SearchInstance() Server {
	return MockServer{isEnabled: f.isEnabled}
}
