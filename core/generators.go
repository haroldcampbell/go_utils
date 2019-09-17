package core

import (
	"fmt"
	"go_utils/util"
	"strings"
)

// GenerateSMSCode Returns an SMS code in the format NNN-NNN
func GenerateSMSCode() string {
	return util.RandDigits(3) + "-" + util.RandDigits(3)
}

// GenerateUUID generates a random uuid similar to E621E1F8-C36C-495A-93FC-0C247A3E6E5F
func GenerateUUID() string {
	return util.RandCharacters(1) + util.RandDigits(3) + util.RandCharNumPair() + util.RandCharNumPair() +
		"-" + util.RandCharNumPair() + util.RandCharNumPair() +
		"-" + util.RandDigits(3) + util.RandCharacters(1) +
		"-" + util.RandCharNumPair() + util.RandCharNumPair() +
		"-" + util.RandCharNumPair() + util.RandDigits(3) + util.RandCharNumPair() + util.RandCharNumPair() + util.RandCharNumPair() + util.RandCharacters(1)
}

// GenerateGUID generates a random guid similar to E621E1F8
func GenerateGUID() string {
	uuid := GenerateUUID()
	uuidParts := strings.Split(uuid, "-")

	return fmt.Sprintf("%s", uuidParts[0])
}
