// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package passenger

// Server is in that Passenger Application
type Server interface {
	// Metrics returns Passenger Application Status.
	Metrics() (*PoolInfo, error)
	// IsEnabled returns Server Instance status.
	//
	// If returns false, this instance is not valid.
	IsEnabled() bool
}

// ServerFactory is factory for Server instance.
type ServerFactory interface {
	// NewInstance creates new Server instance.
	NewInstance(string) Server
	// FindInstance search and create new Server instance.
	FindInstance() Server
}

// Context is Passenger Application Context.
// This is for expansion in future.
type Context struct {
}

// CreateFactory returns the appropriate sever instance factory.
//
// This is for expansion in future.
//
// Depending on which version of Passenger you want to monitor,
// you can change connectivity, and metric collection.
func CreateFactory(c Context) ServerFactory {
	return &Passenger6ServerFactory{}
}
