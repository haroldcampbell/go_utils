package utils

import (
	"strconv"
	"strings"
)

// StrInRange returns true if the string is between min and max (inclusively)
func StrInRange(str string, min int, max int) bool {
	length := len(str)

	return length >= min && length <= max
}

// IsValidMobileNumber returns true if the mobile number seems valid
func IsValidMobileNumber(mobileNumber string) bool {
	length := len(mobileNumber)

	mobileNumber = strings.TrimSpace(mobileNumber)

	if length < 10 || length > 15 {
		return false
	}

	if mobileNumber[0] != '+' {
		return false
	}

	// TODO: Make this more efficient
	mobileNumber = strings.Replace(mobileNumber, "+", "", -1)
	mobileNumber = strings.Replace(mobileNumber, "-", "", -1)
	mobileNumber = strings.Replace(mobileNumber, " ", "", -1)

	_, err := strconv.ParseInt(mobileNumber, 10, 64)
	// fmt.Printf("X: [%s]->%v, %v\n", mobileNumber, i, err)

	return err == nil
}

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
