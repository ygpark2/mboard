// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/board/proto/entities/entities.proto

package entities

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate is disabled for Board. This method will always return nil.
func (m *Board) Validate() error {
	return nil
}

// BoardValidationError is the validation error returned by Board.Validate if
// the designated constraints aren't met.
type BoardValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BoardValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BoardValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BoardValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BoardValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BoardValidationError) ErrorName() string { return "BoardValidationError" }

// Error satisfies the builtin error interface
func (e BoardValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBoard.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BoardValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BoardValidationError{}
