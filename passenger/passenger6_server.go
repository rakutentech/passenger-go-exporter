// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package passenger

import (
	"bufio"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Passenger6Server implements the passenger.Server interface.
type Passenger6Server struct {
	instanceDir       string
	basicAuthPassword string
}

// Passenger6ServerFactory implements the passenger.ServerFactory interface.
type Passenger6ServerFactory struct {
}

// Metrics implements the passenger.Metrics interface.
func (s *Passenger6Server) Metrics() (*PoolInfo, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	req, err := http.NewRequest("GET", "/pool.xml", nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("ro_admin", s.basicAuthPassword)
	if err = req.Write(conn); err != nil {
		return nil, err
	}
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return nil, err
	}

	return ParsePoolInfo(res.Body)
}

// IsEnabled implements the passenger.IsEnabled interface.
func (s *Passenger6Server) IsEnabled() bool {
	conn, err := s.connect()
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// connect is connecting the passenger application through unix socket.
func (s *Passenger6Server) connect() (net.Conn, error) {
	proto := "unix"
	addr := filepath.Join(s.instanceDir, "/agents.s/core_api")
	conn, err := net.Dial(proto, addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewInstance implements the passenger.ServerFactory interface.
//
// If homedir satisfies the following conditions, instance will be returned.
//  - homedir has creation_finalized file.
//  - homedir has read_only_admin_password.txt file.
//  - can connect unix domain socket with /agents.s/core_api.
func (f *Passenger6ServerFactory) NewInstance(homedir string) Server {
	state := filepath.Join(homedir, "creation_finalized")
	if _, err := os.Stat(state); err != nil {
		return nil
	}
	passFile := filepath.Join(homedir, "read_only_admin_password.txt")
	bytes, err := os.ReadFile(passFile)
	if err != nil {
		return nil
	}

	server := Passenger6Server{}
	server.instanceDir = homedir
	server.basicAuthPassword = string(bytes)

	if server.IsEnabled() {
		return &server
	}
	return nil
}

// FindInstance implements the passenger.ServerFactory interface.
//
// search based on TMPDIR and PASSENGER_INSTANCE_REGISTRY_DIR dir.
func (f *Passenger6ServerFactory) FindInstance() Server {
	baseDirs := [2]string{
		os.Getenv("PASSENGER_INSTANCE_REGISTRY_DIR"),
		os.TempDir(),
	}
	for _, dir := range baseDirs {
		if dir == "" {
			continue
		}
		if server := f.findInstanceFirstOne(dir); server != nil {
			return server
		}
	}
	return nil
}

// findInstanceFirstOne searches a Passenger Instance.
func (f *Passenger6ServerFactory) findInstanceFirstOne(dirname string) Server {
	files, err := os.ReadDir(dirname)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "passenger.") {
			homedir := filepath.Join(dirname, file.Name())
			if server := f.NewInstance(homedir); server != nil {
				return server
			}
		}
	}
	return nil
}
