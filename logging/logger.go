// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package logging

import (
	"github.com/go-kit/log"
	"github.com/prometheus/common/promlog"
)

// NewLogger returns log.Logger instance.
//
// logfmt is output log format.
// loglevel is output log level.
// More details,Plase check the following URL.
// https://godoc.org/github.com/prometheus/common/promlog/flag
func NewLogger(logfmt string, loglevel string) log.Logger {
	promLogLevel := &promlog.AllowedLevel{}
	promLogLevel.Set(loglevel)
	promLogFormat := &promlog.AllowedFormat{}
	promLogFormat.Set(logfmt)
	promlogConfig := &promlog.Config{
		Level:  promLogLevel,
		Format: promLogFormat,
	}
	return promlog.New(promlogConfig)
}
