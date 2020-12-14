// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package passenger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInstance(t *testing.T) {
	_, passenger, err := createDummyPassengerDir(os.TempDir())
	if err != nil {
		assert.Error(t, err)
	}
	factory := Passenger6ServerFactory{}
	server := factory.NewInstance(passenger)
	assert.Nil(t, server)
}

func TestNewInstanceCheckCreatetionFinalized(t *testing.T) {
	_, passenger, err := createDummyPassengerDir(os.TempDir())
	if err != nil {
		assert.Error(t, err)
	}
	if err := os.Remove(filepath.Join(passenger, "creation_finalized")); err != nil {
		assert.Error(t, err)
	}
	factory := Passenger6ServerFactory{}
	server := factory.NewInstance(passenger)
	assert.Nil(t, server)
}

func TestNewInstanceCheckPasswordFile(t *testing.T) {
	_, passenger, err := createDummyPassengerDir(os.TempDir())
	if err != nil {
		assert.Error(t, err)
	}
	if err := os.Remove(filepath.Join(passenger, "read_only_admin_password.txt")); err != nil {
		assert.Error(t, err)
	}
	factory := Passenger6ServerFactory{}
	server := factory.NewInstance(passenger)
	assert.Nil(t, server)
}

func TestFindInstanceUsePassenger(t *testing.T) {
	if os.Getenv("USE_PASSENGER") != "true" {
		t.Skip("This test should be execute after build only.")
	}

	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "/sock")
	defer os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "")

	factory := &Passenger6ServerFactory{}
	server := factory.FindInstance()
	assert.NotNil(t, server) //because of not connect.

	pool, err := server.Metrics()
	assert.NotNil(t, pool)
	fmt.Printf("%+v", pool)
	assert.Nil(t, err)
}

func createDummyPassengerDir(dirname string) (string, string, error) {
	home, err := ioutil.TempDir(dirname, "go-unit.*")
	if err != nil {
		return "", "", err
	}
	passengerHome, err := ioutil.TempDir(home, "passenger.*")
	if err != nil {
		return "", "", err
	}
	err = ioutil.WriteFile(filepath.Join(passengerHome, "creation_finalized"), []byte{}, 0644)
	if err != nil {
		return "", "", err
	}

	err = ioutil.WriteFile(filepath.Join(passengerHome, "read_only_admin_password.txt"), []byte("dummy"), 0644)
	if err != nil {
		return "", "", err
	}
	return home, passengerHome, nil
}

func TestFindInstanceTmpDir(t *testing.T) {
	home, _, err := createDummyPassengerDir(os.TempDir())
	if err != nil {
		assert.Error(t, err)
	}
	os.Setenv("TMPDIR", home)
	defer os.Setenv("TMPDIR", "")

	factory := &Passenger6ServerFactory{}
	server := factory.FindInstance()
	assert.Nil(t, server)

}

func TestFindInstanceRegDir(t *testing.T) {
	home, _, err := createDummyPassengerDir(os.TempDir())
	if err != nil {
		assert.Error(t, err)
	}
	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", home)
	defer os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "")

	factory := &Passenger6ServerFactory{}
	server := factory.FindInstance()
	assert.Nil(t, server)
}

func TestFindInstanceNotFound(t *testing.T) {
	os.Setenv("TMPDIR", os.TempDir())
	defer os.Setenv("TMPDIR", "")

	factory := &Passenger6ServerFactory{}
	server := factory.FindInstance()
	assert.Nil(t, server)
}

func TestFindInstanceNotDir(t *testing.T) {
	os.Setenv("TMPDIR", "not-exists")
	defer os.Setenv("TMPDIR", "")

	factory := &Passenger6ServerFactory{}
	server := factory.FindInstance()
	assert.Nil(t, server)
}

func TestFindInstanceXMLNotConnect(t *testing.T) {
	home, _, err := createDummyPassengerDir(os.TempDir())
	if err != nil {
		assert.Error(t, err)
	}
	os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", home)
	defer os.Setenv("PASSENGER_INSTANCE_REGISTRY_DIR", "")

	factory := &Passenger6ServerFactory{}
	server := factory.FindInstance()
	assert.Nil(t, server) //because of not connect.
}
