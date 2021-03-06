// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v3/dynamic_ot.proto

package envoy_config_trace_v3

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

// Validate checks the field values on DynamicOtConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DynamicOtConfig) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetLibrary()) < 1 {
		return DynamicOtConfigValidationError{
			field:  "Library",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamicOtConfigValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DynamicOtConfigValidationError is the validation error returned by
// DynamicOtConfig.Validate if the designated constraints aren't met.
type DynamicOtConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamicOtConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamicOtConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamicOtConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamicOtConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamicOtConfigValidationError) ErrorName() string { return "DynamicOtConfigValidationError" }

// Error satisfies the builtin error interface
func (e DynamicOtConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamicOtConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamicOtConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamicOtConfigValidationError{}
