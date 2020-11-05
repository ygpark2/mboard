// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/posts/proto/post/post.proto

package posts

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
var _post_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Post with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Post) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Slug

	// no validation rules for Content

	// no validation rules for Timestamp

	return nil
}

// PostValidationError is the validation error returned by Post.Validate if the
// designated constraints aren't met.
type PostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostValidationError) ErrorName() string { return "PostValidationError" }

// Error satisfies the builtin error interface
func (e PostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostValidationError{}

// Validate checks the field values on QueryRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *QueryRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Slug

	// no validation rules for Offset

	// no validation rules for Limit

	return nil
}

// QueryRequestValidationError is the validation error returned by
// QueryRequest.Validate if the designated constraints aren't met.
type QueryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryRequestValidationError) ErrorName() string { return "QueryRequestValidationError" }

// Error satisfies the builtin error interface
func (e QueryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryRequestValidationError{}

// Validate checks the field values on QueryResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *QueryResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryResponseValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// QueryResponseValidationError is the validation error returned by
// QueryResponse.Validate if the designated constraints aren't met.
type QueryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryResponseValidationError) ErrorName() string { return "QueryResponseValidationError" }

// Error satisfies the builtin error interface
func (e QueryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryResponseValidationError{}

// Validate checks the field values on SaveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SaveRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveRequestValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SaveRequestValidationError is the validation error returned by
// SaveRequest.Validate if the designated constraints aren't met.
type SaveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveRequestValidationError) ErrorName() string { return "SaveRequestValidationError" }

// Error satisfies the builtin error interface
func (e SaveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveRequestValidationError{}

// Validate checks the field values on SaveResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SaveResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveResponseValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SaveResponseValidationError is the validation error returned by
// SaveResponse.Validate if the designated constraints aren't met.
type SaveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveResponseValidationError) ErrorName() string { return "SaveResponseValidationError" }

// Error satisfies the builtin error interface
func (e SaveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveResponseValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}

// Validate checks the field values on DeleteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteResponseValidationError is the validation error returned by
// DeleteResponse.Validate if the designated constraints aren't met.
type DeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResponseValidationError) ErrorName() string { return "DeleteResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResponseValidationError{}
