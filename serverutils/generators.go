package serverutils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/haroldcampbell/go_utils/utils"
)

// GenerateSMSCode Returns an SMS code in the format NNN-NNN
func GenerateSMSCode() string {
	return utils.RandDigits(3) + "-" + utils.RandDigits(3)
}

// GenerateUUID generates a random uuid similar to E621E1F8-C36C-495A-93FC-0C247A3E6E5F
func GenerateUUID() string {
	return utils.RandCharacters(1) + utils.RandDigits(3) + utils.RandCharNumPair() + utils.RandCharNumPair() +
		"-" + utils.RandCharNumPair() + utils.RandCharNumPair() +
		"-" + utils.RandDigits(3) + utils.RandCharacters(1) +
		"-" + utils.RandCharNumPair() + utils.RandCharNumPair() +
		"-" + utils.RandCharNumPair() + utils.RandDigits(3) + utils.RandCharNumPair() + utils.RandCharNumPair() + utils.RandCharNumPair() + utils.RandCharacters(1)
}

// GenerateGUID generates a random guid similar to E621E1F8
func GenerateGUID() string {
	uuid := GenerateUUID()
	uuidParts := strings.Split(uuid, "-")

	return fmt.Sprintf("%s", uuidParts[0])
}

// IsValidGUID checks if the specified guid is valid
func IsValidGUID(guid string) bool {
	var validGUID = regexp.MustCompile(`[a-zA-Z][0-9]{3}[a-zA-Z][0-9][a-zA-Z][0-9]`)

	return validGUID.MatchString(guid)
}

// IsValidSLUG generates the first 8 characters of random uuid
func IsValidSLUG(slug string) bool {
	//V987X4U7
	var validSLUG = regexp.MustCompile(`[a-zA-Z]{1}[0-9]{3}([a-zA-Z][0-9]){2}`)

	return validSLUG.MatchString(slug)
}
