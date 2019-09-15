package core

import "go_utils/util"

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
