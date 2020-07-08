// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tls/v4alpha/cert.proto

package envoy_extensions_transport_sockets_tls_v4alpha

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

// define the regex for a UUID once up-front
var _cert_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TlsParameters with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TlsParameters) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := TlsParameters_TlsProtocol_name[int32(m.GetTlsMinimumProtocolVersion())]; !ok {
		return TlsParametersValidationError{
			field:  "TlsMinimumProtocolVersion",
			reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := TlsParameters_TlsProtocol_name[int32(m.GetTlsMaximumProtocolVersion())]; !ok {
		return TlsParametersValidationError{
			field:  "TlsMaximumProtocolVersion",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// TlsParametersValidationError is the validation error returned by
// TlsParameters.Validate if the designated constraints aren't met.
type TlsParametersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlsParametersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlsParametersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlsParametersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlsParametersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlsParametersValidationError) ErrorName() string { return "TlsParametersValidationError" }

// Error satisfies the builtin error interface
func (e TlsParametersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTlsParameters.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlsParametersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlsParametersValidationError{}

// Validate checks the field values on PrivateKeyProvider with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PrivateKeyProvider) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetProviderName()) < 1 {
		return PrivateKeyProviderValidationError{
			field:  "ProviderName",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *PrivateKeyProvider_TypedConfig:

		{
			tmp := m.GetTypedConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return PrivateKeyProviderValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// PrivateKeyProviderValidationError is the validation error returned by
// PrivateKeyProvider.Validate if the designated constraints aren't met.
type PrivateKeyProviderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrivateKeyProviderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrivateKeyProviderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrivateKeyProviderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrivateKeyProviderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrivateKeyProviderValidationError) ErrorName() string {
	return "PrivateKeyProviderValidationError"
}

// Error satisfies the builtin error interface
func (e PrivateKeyProviderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrivateKeyProvider.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrivateKeyProviderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrivateKeyProviderValidationError{}

// Validate checks the field values on TlsCertificate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TlsCertificate) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetCertificateChain()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TlsCertificateValidationError{
					field:  "CertificateChain",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetPrivateKey()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TlsCertificateValidationError{
					field:  "PrivateKey",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetPrivateKeyProvider()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TlsCertificateValidationError{
					field:  "PrivateKeyProvider",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetPassword()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TlsCertificateValidationError{
					field:  "Password",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetOcspStaple()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TlsCertificateValidationError{
					field:  "OcspStaple",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetSignedCertificateTimestamp() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return TlsCertificateValidationError{
						field:  fmt.Sprintf("SignedCertificateTimestamp[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// TlsCertificateValidationError is the validation error returned by
// TlsCertificate.Validate if the designated constraints aren't met.
type TlsCertificateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlsCertificateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlsCertificateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlsCertificateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlsCertificateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlsCertificateValidationError) ErrorName() string { return "TlsCertificateValidationError" }

// Error satisfies the builtin error interface
func (e TlsCertificateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTlsCertificate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlsCertificateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlsCertificateValidationError{}

// Validate checks the field values on TlsSessionTicketKeys with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TlsSessionTicketKeys) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetKeys()) < 1 {
		return TlsSessionTicketKeysValidationError{
			field:  "Keys",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetKeys() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return TlsSessionTicketKeysValidationError{
						field:  fmt.Sprintf("Keys[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// TlsSessionTicketKeysValidationError is the validation error returned by
// TlsSessionTicketKeys.Validate if the designated constraints aren't met.
type TlsSessionTicketKeysValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlsSessionTicketKeysValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlsSessionTicketKeysValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlsSessionTicketKeysValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlsSessionTicketKeysValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlsSessionTicketKeysValidationError) ErrorName() string {
	return "TlsSessionTicketKeysValidationError"
}

// Error satisfies the builtin error interface
func (e TlsSessionTicketKeysValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTlsSessionTicketKeys.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlsSessionTicketKeysValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlsSessionTicketKeysValidationError{}

// Validate checks the field values on CertificateValidationContext with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CertificateValidationContext) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetTrustedCa()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CertificateValidationContextValidationError{
					field:  "TrustedCa",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetVerifyCertificateSpki() {
		_, _ = idx, item

		if len(item) != 44 {
			return CertificateValidationContextValidationError{
				field:  fmt.Sprintf("VerifyCertificateSpki[%v]", idx),
				reason: "value length must be 44 bytes",
			}
		}

	}

	for idx, item := range m.GetVerifyCertificateHash() {
		_, _ = idx, item

		if l := len(item); l < 64 || l > 95 {
			return CertificateValidationContextValidationError{
				field:  fmt.Sprintf("VerifyCertificateHash[%v]", idx),
				reason: "value length must be between 64 and 95 bytes, inclusive",
			}
		}

	}

	for idx, item := range m.GetMatchSubjectAltNames() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CertificateValidationContextValidationError{
						field:  fmt.Sprintf("MatchSubjectAltNames[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetRequireOcspStaple()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CertificateValidationContextValidationError{
					field:  "RequireOcspStaple",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRequireSignedCertificateTimestamp()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CertificateValidationContextValidationError{
					field:  "RequireSignedCertificateTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetCrl()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CertificateValidationContextValidationError{
					field:  "Crl",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for AllowExpiredCertificate

	if _, ok := CertificateValidationContext_TrustChainVerification_name[int32(m.GetTrustChainVerification())]; !ok {
		return CertificateValidationContextValidationError{
			field:  "TrustChainVerification",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// CertificateValidationContextValidationError is the validation error returned
// by CertificateValidationContext.Validate if the designated constraints
// aren't met.
type CertificateValidationContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateValidationContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateValidationContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateValidationContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateValidationContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateValidationContextValidationError) ErrorName() string {
	return "CertificateValidationContextValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateValidationContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateValidationContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateValidationContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateValidationContextValidationError{}

// Validate checks the field values on CommonTlsContext with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CommonTlsContext) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetTlsParams()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "TlsParams",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetTlsCertificates() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CommonTlsContextValidationError{
						field:  fmt.Sprintf("TlsCertificates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	if len(m.GetTlsCertificateSdsSecretConfigs()) > 1 {
		return CommonTlsContextValidationError{
			field:  "TlsCertificateSdsSecretConfigs",
			reason: "value must contain no more than 1 item(s)",
		}
	}

	for idx, item := range m.GetTlsCertificateSdsSecretConfigs() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CommonTlsContextValidationError{
						field:  fmt.Sprintf("TlsCertificateSdsSecretConfigs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	switch m.ValidationContextType.(type) {

	case *CommonTlsContext_ValidationContext:

		{
			tmp := m.GetValidationContext()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CommonTlsContextValidationError{
						field:  "ValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *CommonTlsContext_ValidationContextSdsSecretConfig:

		{
			tmp := m.GetValidationContextSdsSecretConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CommonTlsContextValidationError{
						field:  "ValidationContextSdsSecretConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *CommonTlsContext_CombinedValidationContext:

		{
			tmp := m.GetCombinedValidationContext()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CommonTlsContextValidationError{
						field:  "CombinedValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// CommonTlsContextValidationError is the validation error returned by
// CommonTlsContext.Validate if the designated constraints aren't met.
type CommonTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonTlsContextValidationError) ErrorName() string { return "CommonTlsContextValidationError" }

// Error satisfies the builtin error interface
func (e CommonTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContextValidationError{}

// Validate checks the field values on UpstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpstreamTlsContext) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetCommonTlsContext()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UpstreamTlsContextValidationError{
					field:  "CommonTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if len(m.GetSni()) > 255 {
		return UpstreamTlsContextValidationError{
			field:  "Sni",
			reason: "value length must be at most 255 bytes",
		}
	}

	// no validation rules for AllowRenegotiation

	{
		tmp := m.GetMaxSessionKeys()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UpstreamTlsContextValidationError{
					field:  "MaxSessionKeys",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// UpstreamTlsContextValidationError is the validation error returned by
// UpstreamTlsContext.Validate if the designated constraints aren't met.
type UpstreamTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamTlsContextValidationError) ErrorName() string {
	return "UpstreamTlsContextValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamTlsContextValidationError{}

// Validate checks the field values on DownstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DownstreamTlsContext) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetCommonTlsContext()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "CommonTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRequireClientCertificate()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "RequireClientCertificate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRequireSni()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "RequireSni",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if d := m.GetSessionTimeout(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return DownstreamTlsContextValidationError{
				field:  "SessionTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		lt := time.Duration(4294967296*time.Second + 0*time.Nanosecond)
		gte := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur < gte || dur >= lt {
			return DownstreamTlsContextValidationError{
				field:  "SessionTimeout",
				reason: "value must be inside range [0s, 1193046h28m16s)",
			}
		}

	}

	switch m.SessionTicketKeysType.(type) {

	case *DownstreamTlsContext_SessionTicketKeys:

		{
			tmp := m.GetSessionTicketKeys()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return DownstreamTlsContextValidationError{
						field:  "SessionTicketKeys",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *DownstreamTlsContext_SessionTicketKeysSdsSecretConfig:

		{
			tmp := m.GetSessionTicketKeysSdsSecretConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return DownstreamTlsContextValidationError{
						field:  "SessionTicketKeysSdsSecretConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *DownstreamTlsContext_DisableStatelessSessionResumption:
		// no validation rules for DisableStatelessSessionResumption

	}

	return nil
}

// DownstreamTlsContextValidationError is the validation error returned by
// DownstreamTlsContext.Validate if the designated constraints aren't met.
type DownstreamTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownstreamTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownstreamTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownstreamTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownstreamTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownstreamTlsContextValidationError) ErrorName() string {
	return "DownstreamTlsContextValidationError"
}

// Error satisfies the builtin error interface
func (e DownstreamTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownstreamTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownstreamTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownstreamTlsContextValidationError{}

// Validate checks the field values on GenericSecret with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GenericSecret) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetSecret()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return GenericSecretValidationError{
					field:  "Secret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// GenericSecretValidationError is the validation error returned by
// GenericSecret.Validate if the designated constraints aren't met.
type GenericSecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenericSecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenericSecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenericSecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenericSecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenericSecretValidationError) ErrorName() string { return "GenericSecretValidationError" }

// Error satisfies the builtin error interface
func (e GenericSecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenericSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenericSecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenericSecretValidationError{}

// Validate checks the field values on SdsSecretConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SdsSecretConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	{
		tmp := m.GetSdsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return SdsSecretConfigValidationError{
					field:  "SdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// SdsSecretConfigValidationError is the validation error returned by
// SdsSecretConfig.Validate if the designated constraints aren't met.
type SdsSecretConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SdsSecretConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SdsSecretConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SdsSecretConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SdsSecretConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SdsSecretConfigValidationError) ErrorName() string { return "SdsSecretConfigValidationError" }

// Error satisfies the builtin error interface
func (e SdsSecretConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSdsSecretConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SdsSecretConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SdsSecretConfigValidationError{}

// Validate checks the field values on Secret with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Secret) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	switch m.Type.(type) {

	case *Secret_TlsCertificate:

		{
			tmp := m.GetTlsCertificate()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return SecretValidationError{
						field:  "TlsCertificate",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *Secret_SessionTicketKeys:

		{
			tmp := m.GetSessionTicketKeys()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return SecretValidationError{
						field:  "SessionTicketKeys",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *Secret_ValidationContext:

		{
			tmp := m.GetValidationContext()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return SecretValidationError{
						field:  "ValidationContext",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *Secret_GenericSecret:

		{
			tmp := m.GetGenericSecret()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return SecretValidationError{
						field:  "GenericSecret",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// SecretValidationError is the validation error returned by Secret.Validate if
// the designated constraints aren't met.
type SecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretValidationError) ErrorName() string { return "SecretValidationError" }

// Error satisfies the builtin error interface
func (e SecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretValidationError{}

// Validate checks the field values on
// CommonTlsContext_CombinedCertificateValidationContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CommonTlsContext_CombinedCertificateValidationContext) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDefaultValidationContext() == nil {
		return CommonTlsContext_CombinedCertificateValidationContextValidationError{
			field:  "DefaultValidationContext",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetDefaultValidationContext()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CommonTlsContext_CombinedCertificateValidationContextValidationError{
					field:  "DefaultValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if m.GetValidationContextSdsSecretConfig() == nil {
		return CommonTlsContext_CombinedCertificateValidationContextValidationError{
			field:  "ValidationContextSdsSecretConfig",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetValidationContextSdsSecretConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CommonTlsContext_CombinedCertificateValidationContextValidationError{
					field:  "ValidationContextSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// CommonTlsContext_CombinedCertificateValidationContextValidationError is the
// validation error returned by
// CommonTlsContext_CombinedCertificateValidationContext.Validate if the
// designated constraints aren't met.
type CommonTlsContext_CombinedCertificateValidationContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) ErrorName() string {
	return "CommonTlsContext_CombinedCertificateValidationContextValidationError"
}

// Error satisfies the builtin error interface
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext_CombinedCertificateValidationContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContext_CombinedCertificateValidationContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContext_CombinedCertificateValidationContextValidationError{}
