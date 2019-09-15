package core

import (
	"strings"
)

// IsValidSMSCode returns true iff the code is a valid SMSCode
func IsValidSMSCode(code string) bool {
	//NNN-NNN
	if code == "" {
		return false
	}

	if len(code) != 7 {
		return false
	}

	parts := strings.Split(code, "-")
	if len(parts) != 2 {
		return false
	}

	expectedLengths := [2]int{3, 3}
	for index, part := range parts {
		if expectedLengths[index] != len(part) {
			return false
		}
	}

	return true
}

// IsValidUUID returns true iff the key is a valid UUID
func IsValidUUID(key string) bool {
	//E621E1F8-C36C-495A-93FC-0C247A3E6E5F
	if key == "" {
		return false
	}

	if len(key) != 36 {
		return false
	}

	parts := strings.Split(key, "-")
	if len(parts) != 5 {
		return false
	}

	expectedLengths := [5]int{8, 4, 4, 4, 12}
	for index, part := range parts {
		if expectedLengths[index] != len(part) {
			// log.Printf("IsValidUUID: invalid part[%d]=\"%v\" -> expected length:%d\n", index, part, len(part))
			return false
		}
	}
	return true
}
