package flag

import (
	"crypto/tls"
	"reflect"
	"testing"
)

func TestStrToUInt16(t *testing.T) {
	tests := []struct {
		flag           []string
		expected       []uint16
		expected_error bool
	}{
		{
			// Happy case
			flag:           []string{"TLS_RSA_WITH_RC4_128_SHA", "TLS_RSA_WITH_AES_128_CBC_SHA", "TLS_ECDHE_RSA_WITH_RC4_128_SHA", "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"},
			expected:       []uint16{tls.TLS_RSA_WITH_RC4_128_SHA, tls.TLS_RSA_WITH_AES_128_CBC_SHA, tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA},
			expected_error: false,
		},
		{
			// One flag only
			flag:           []string{"TLS_RSA_WITH_RC4_128_SHA"},
			expected:       []uint16{tls.TLS_RSA_WITH_RC4_128_SHA},
			expected_error: false,
		},
		{
			// Empty flag
			flag:           []string{},
			expected:       nil,
			expected_error: false,
		},
		{
			// Duplicated flag
			flag:           []string{"TLS_RSA_WITH_RC4_128_SHA", "TLS_RSA_WITH_AES_128_CBC_SHA", "TLS_ECDHE_RSA_WITH_RC4_128_SHA", "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA", "TLS_RSA_WITH_RC4_128_SHA"},
			expected:       []uint16{tls.TLS_RSA_WITH_RC4_128_SHA, tls.TLS_RSA_WITH_AES_128_CBC_SHA, tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, tls.TLS_RSA_WITH_RC4_128_SHA},
			expected_error: false,
		},
		{
			// Invalid flag
			flag:           []string{"foo"},
			expected:       nil,
			expected_error: true,
		},
		{
			// All existing cipher suites
			flag: []string{
				"TLS_RSA_WITH_3DES_EDE_CBC_SHA",
				"TLS_RSA_WITH_AES_128_CBC_SHA",
				"TLS_RSA_WITH_AES_256_CBC_SHA",
				"TLS_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
				"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
				"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
				"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
				"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
				"TLS_AES_128_GCM_SHA256",
				"TLS_CHACHA20_POLY1305_SHA256",
				"TLS_AES_256_GCM_SHA384",

				"TLS_RSA_WITH_RC4_128_SHA",
				"TLS_RSA_WITH_AES_128_CBC_SHA256",
				"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
				"TLS_ECDHE_RSA_WITH_RC4_128_SHA",
				"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
			},
			expected: []uint16{
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_AES_256_GCM_SHA384,

				tls.TLS_RSA_WITH_RC4_128_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			},
		},
	}

	for i, test := range tests {
		uIntFlags, err := TLSCipherSuites(test.flag)
		if !reflect.DeepEqual(uIntFlags, test.expected) {
			t.Errorf("%d: expected %+v, got %+v", i, test.expected, uIntFlags)
		}
		if test.expected_error && err == nil {
			t.Errorf("%d: expecting error, got %+v", i, err)
		}
	}
}
