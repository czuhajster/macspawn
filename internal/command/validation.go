package command

import (
	"regexp"
)

func validateIdentifierFlag(identifier string) (bool, error) {
	re, err := regexp.Compile(`^[0-9a-fA-F]{2}[\-:]?[0-9a-fA-F]{2}[\-:]?[0-9a-fA-F]{2}`)
	if err != nil {
		return false, err
	}
	valid := re.MatchString(identifier)
	return valid, nil
}
