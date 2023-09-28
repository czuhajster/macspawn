package command

import (
	"encoding/hex"
	"regexp"
)

func processIdentifier(identifierString string) ([]byte, error) {
	re, err := regexp.Compile(`[\-:]`)
	if err != nil {
		return nil, err
	}
	identifierString = re.ReplaceAllString(identifierString, "")
	identifier, identifierError := hex.DecodeString(identifierString)
	if identifierError != nil {
		return identifier, identifierError
	}
	return identifier, nil
}
