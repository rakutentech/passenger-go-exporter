// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package passenger

import (
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"
)

// PoolInfo is the structure retuened by passenger-status.
type PoolInfo struct {
	XMLName          xml.Name     `xml:"info"`
	PassengerVersion string       `xml:"passenger_version"`
	GroupCount       int          `xml:"group_count"`
	ProcessCount     int          `xml:"process_count"`
	Max              int          `xml:"max"`
	GetWaitListSize  int          `xml:"get_wait_list_size"`
	SuperGroups      []SuperGroup `xml:"supergroups>supergroup"`
}

// SuperGroup is the structure retuened by passenger-status.
type SuperGroup struct {
	Name            string `xml:"name"`
	State           string `xml:"state"`
	GetWaitListSize int    `xml:"get_wait_list_size"`
	CapacityUsed    int    `xml:"capacity_used"`
	Group           Group  `xml:"group"`
}

// Group is the structure retuened by passenger-status.
type Group struct {
	Name            string    `xml:"name"`
	UUID            string    `xml:"uuid"`
	LifeStatus      string    `xml:"life_status"`
	GetWaitListSize int       `xml:"get_wait_list_size"`
	Processes       []Process `xml:"processes>process"`
}

// Process is the structure retuened by passenger-status.
type Process struct {
	PID        string  `xml:"pid"`
	Enabled    string  `xml:"enabled"`
	LifeStatus string  `xml:"life_status"`
	Processed  float64 `xml:"processed"`
	RSS        int64   `xml:"rss"`
	PSS        int64   `xml:"pss"`
	RealMemory int64   `xml:"real_memory"`
	VMSize     int64   `xml:"vmsize"`
}

// ParsePoolInfo returns PoolInfo instance.
func ParsePoolInfo(reader io.Reader) (*PoolInfo, error) {
	var info PoolInfo

	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
