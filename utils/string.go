package utils

import "encoding/json"

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
