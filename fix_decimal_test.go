package quickfix

import (
	"testing"

	"git.cryptology.com/lib/go/fixed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFIXDecimalWrite(t *testing.T) {
	var tests = []struct {
		decimal  FIXDecimal
		expected string
	}{
		{decimal: FIXDecimal{Fixed: fixed.NewFromString("-124.3456", -4), Scale: 4}, expected: "-124.3456"},
		{decimal: FIXDecimal{Fixed: fixed.NewFromString("-124.3456", -4), Scale: 0}, expected: "-124"},
	}

	for _, test := range tests {
		b := test.decimal.Write()
		assert.Equal(t, test.expected, string(b))
	}
}

func TestFIXDecimalRead(t *testing.T) {
	var tests = []struct {
		bytes       string
		expected    fixed.Fixed
		expectError bool
	}{
		{bytes: "15", expected: fixed.RequireFromString("15")},
		{bytes: "15.000", expected: fixed.RequireFromString("15")},
		{bytes: "15.001", expected: fixed.RequireFromString("15.001")},
		{bytes: "-15.001", expected: fixed.RequireFromString("-15.001")},
		{bytes: "blah", expectError: true},
		{bytes: "+200.00", expected: expected: fixed.RequireFromString("200")},
	}

	for _, test := range tests {
		var field FIXDecimal

		err := field.Read([]byte(test.bytes))
		require.Equal(t, test.expectError, err != nil)

		if !test.expectError {
			assert.True(t, test.expected.Equals(field.Fixed), "Expected %s got %s", test.expected, field.Fixed)
		}
	}
}
