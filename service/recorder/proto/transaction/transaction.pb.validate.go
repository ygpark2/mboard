// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/recorder/proto/transaction/transaction.proto

package transaction

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate is disabled for ReadRequest. This method will always return nil.
func (m *ReadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll is disabled for ReadRequest. This method will always return nil.
func (m *ReadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadRequest) validate(all bool) error {
	return nil
}

// ReadRequestMultiError is an error wrapping multiple validation errors
// returned by ReadRequest.ValidateAll() if the designated constraints aren't met.
type ReadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadRequestMultiError) AllErrors() []error { return m }

// ReadRequestValidationError is the validation error returned by
// ReadRequest.Validate if the designated constraints aren't met.
type ReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadRequestValidationError) ErrorName() string { return "ReadRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadRequestValidationError{}

// Validate is disabled for WriteRequest. This method will always return nil.
func (m *WriteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll is disabled for WriteRequest. This method will always return nil.
func (m *WriteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteRequest) validate(all bool) error {
	return nil
}

// WriteRequestMultiError is an error wrapping multiple validation errors
// returned by WriteRequest.ValidateAll() if the designated constraints aren't met.
type WriteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteRequestMultiError) AllErrors() []error { return m }

// WriteRequestValidationError is the validation error returned by
// WriteRequest.Validate if the designated constraints aren't met.
type WriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteRequestValidationError) ErrorName() string { return "WriteRequestValidationError" }

// Error satisfies the builtin error interface
func (e WriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteRequestValidationError{}

// Validate checks the field values on TransactionEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TransactionEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransactionEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TransactionEventMultiError, or nil if none found.
func (m *TransactionEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *TransactionEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Req

	// no validation rules for Rsp

	if len(errors) > 0 {
		return TransactionEventMultiError(errors)
	}
	return nil
}

// TransactionEventMultiError is an error wrapping multiple validation errors
// returned by TransactionEvent.ValidateAll() if the designated constraints
// aren't met.
type TransactionEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactionEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactionEventMultiError) AllErrors() []error { return m }

// TransactionEventValidationError is the validation error returned by
// TransactionEvent.Validate if the designated constraints aren't met.
type TransactionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionEventValidationError) ErrorName() string { return "TransactionEventValidationError" }

// Error satisfies the builtin error interface
func (e TransactionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionEventValidationError{}
