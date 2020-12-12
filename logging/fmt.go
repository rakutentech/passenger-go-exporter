// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIT License.
// License that can be found in the LICENSE file.

package logging

import "fmt"

const (
	// MessageLabel is message label key for logger.
	MessageLabel = "msg"
)

// Msg returns MESSAGE_LABEL and message of argument.
func Msg(message string) (string, string) {
	return MessageLabel, message
}

// Msgf returns MessageLabel and formatted message of argument.
func Msgf(format string, args ...interface{}) (string, string) {
	return MessageLabel, fmt.Sprintf(format, args...)
}

// Err returns key and error message for logging
func Err(err *error) (string, string) {
	return MessageLabel, fmt.Sprintf("%+v", *err)
}
