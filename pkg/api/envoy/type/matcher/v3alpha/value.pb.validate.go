// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/v3alpha/value.proto

package envoy_type_matcher_v3alpha

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on ValueMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ValueMatcher) Validate() error {
	if m == nil {
		return nil
	}

	switch m.MatchPattern.(type) {

	case *ValueMatcher_NullMatch_:

		{
			tmp := m.GetNullMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ValueMatcherValidationError{
						field:  "NullMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *ValueMatcher_DoubleMatch:

		{
			tmp := m.GetDoubleMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ValueMatcherValidationError{
						field:  "DoubleMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *ValueMatcher_StringMatch:

		{
			tmp := m.GetStringMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ValueMatcherValidationError{
						field:  "StringMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *ValueMatcher_BoolMatch:
		// no validation rules for BoolMatch

	case *ValueMatcher_PresentMatch:
		// no validation rules for PresentMatch

	case *ValueMatcher_ListMatch:

		{
			tmp := m.GetListMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ValueMatcherValidationError{
						field:  "ListMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return ValueMatcherValidationError{
			field:  "MatchPattern",
			reason: "value is required",
		}

	}

	return nil
}

// ValueMatcherValidationError is the validation error returned by
// ValueMatcher.Validate if the designated constraints aren't met.
type ValueMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValueMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValueMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValueMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValueMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValueMatcherValidationError) ErrorName() string { return "ValueMatcherValidationError" }

// Error satisfies the builtin error interface
func (e ValueMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValueMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValueMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValueMatcherValidationError{}

// Validate checks the field values on ListMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListMatcher) Validate() error {
	if m == nil {
		return nil
	}

	switch m.MatchPattern.(type) {

	case *ListMatcher_OneOf:

		{
			tmp := m.GetOneOf()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListMatcherValidationError{
						field:  "OneOf",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return ListMatcherValidationError{
			field:  "MatchPattern",
			reason: "value is required",
		}

	}

	return nil
}

// ListMatcherValidationError is the validation error returned by
// ListMatcher.Validate if the designated constraints aren't met.
type ListMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMatcherValidationError) ErrorName() string { return "ListMatcherValidationError" }

// Error satisfies the builtin error interface
func (e ListMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMatcherValidationError{}

// Validate checks the field values on ValueMatcher_NullMatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ValueMatcher_NullMatch) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ValueMatcher_NullMatchValidationError is the validation error returned by
// ValueMatcher_NullMatch.Validate if the designated constraints aren't met.
type ValueMatcher_NullMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValueMatcher_NullMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValueMatcher_NullMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValueMatcher_NullMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValueMatcher_NullMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValueMatcher_NullMatchValidationError) ErrorName() string {
	return "ValueMatcher_NullMatchValidationError"
}

// Error satisfies the builtin error interface
func (e ValueMatcher_NullMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValueMatcher_NullMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValueMatcher_NullMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValueMatcher_NullMatchValidationError{}
