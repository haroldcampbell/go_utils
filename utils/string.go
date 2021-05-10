package utils

import (
	"encoding/json"
	"regexp"
	"testing"
)

// ToString pretty printf for objects
//
// eg.
//
// 		jsonBody := response["jsonBody"].(map[string]interface{})
// 		fmt.Printf("jsonBody: %v\n", utils.ToString(jsonBody))
//
// Output:
//
// 		jsonBody: {
// 		  "UUID": "W793S5W5-C8E2-619F-B3V4-Q1496S1B3G6M",
// 		  "conversationID": "P348N6U1-P8Z2-458J-D6H2-H3825P6V8F5U",
// 		  "partnerProxyID": "VCTEUEKKZQTHCNAT",
// 		  "proxyID": "HHPPHPMCYFTYQRUN"
// 		}
func ToString(v interface{}, offset ...string) string {
	if v == nil {
		return ""
	}

	var prefix = ""
	var indent = "   "
	if len(offset) > 0 {
		prefix = offset[0]
	}

	b, _ := json.MarshalIndent(v, prefix, indent)

	str := string(b)
	return str
}

const exprIsCommaSeparatedString = `^(\d{1,3}\,){0,}\d{1,3}(\.\d+)?$`

var regIsCommaSeparatedString, _ = regexp.Compile(exprIsCommaSeparatedString)

func IsCommaSeparatedString(s string) bool {
	return regIsCommaSeparatedString.MatchString(s)
}

func TestIsCommaSeparatedString(t *testing.T) {
	validCommaSeparatedNumbers := []string{
		"10",
		"1,000",
		"0",
		"12.3441",
		"0.1223",
		"5,555.00",
		"100,000.00",
		"2,000",
		"10,000.00",
		"10,000.0",
		"1,000,000.00",
		"123,123,123,000.0",
	}

	invalidCommaSeparatedNumbers := []string{
		"aaaa",
		"1000", //No comma
		"1000.00",
		"1,0,0",
	}

	for _, numStr := range validCommaSeparatedNumbers {
		isValid := IsCommaSeparatedString(numStr)
		if !isValid {
			t.Errorf("Expected IsCommaSeparatedString(%s) => true\n", numStr)
		}
	}

	for _, numStr := range invalidCommaSeparatedNumbers {
		isValid := IsCommaSeparatedString(numStr)
		if isValid {
			t.Errorf("Expected IsCommaSeparatedString(%s) => false\n", numStr)
		}
	}
}
