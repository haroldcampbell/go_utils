package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValidMobileNumber(t *testing.T) {
	t.Run("Mauritius", mauritius_test)
	t.Run("Jamaica", jamaica_test)
	t.Run("Length", length_test)
	t.Run("Plus", plus_test)
	t.Run("Edge case", edge_case_test)
}

func edge_case_test(t *testing.T) {
	testNumber := "+aaaaaaaaaa"
	result := IsValidMobileNumber(testNumber)
	assert.Falsef(t, result, "expected false: %v", testNumber)
}

func mauritius_test(t *testing.T) {
	testNumber := "+230 5 909 6012"
	result := IsValidMobileNumber(testNumber)
	assert.Truef(t, result, "expected true valid numbers: %v", testNumber)
}
func jamaica_test(t *testing.T) {
	testNumber := "+1 876 399 3302"
	result := IsValidMobileNumber(testNumber)
	assert.Truef(t, result, "expected true valid numbers: %v", testNumber)
}

func length_test(t *testing.T) {
	testNumber := "+aaaaaaa"
	result := IsValidMobileNumber(testNumber)
	assert.Falsef(t, result, "short garbage should fail: %v", testNumber)

	testNumber = "+aaaaaaaaaaaaaaa"
	result = IsValidMobileNumber(testNumber)
	assert.Falsef(t, result, "long garbage should fail: %v", testNumber)

	testNumber = "000"
	result = IsValidMobileNumber(testNumber)
	assert.Falsef(t, result, "short numbers should fail: %v", testNumber)

	testNumber = "0000000000000001"
	result = IsValidMobileNumber(testNumber)
	assert.Falsef(t, result, "long numbers should fail: %v", testNumber)
}

func plus_test(t *testing.T) {
	testNumber := "230 5 909 6012"
	result := IsValidMobileNumber(testNumber)
	assert.Falsef(t, result, "missing '+' should fail: %v", testNumber)

	testNumber = "+230 5 909 6012"
	result = IsValidMobileNumber(testNumber)
	assert.Truef(t, result, "when '+' present should not fail: %v", testNumber)

}
