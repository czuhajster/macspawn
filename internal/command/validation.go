package command

import (
	"regexp"
)

// Matches one of the following:
// - XX:XX:XX
// - XX:XX:XX:X
// - XX:XX:XX:XX:X
// where `X` is a hexadecimal symbol and `:` can be replaced either with `-` or
// with the empty string.
const identifierRegexp string = `^[0-9a-fA-F]{2}[\-:]?[0-9a-fA-F]{2}[\-:]?[0-9a-fA-F]{2}(|([\-:]?[0-9a-fA-F]{1})|([\-:]?[0-9a-fA-F]{2}){2}[\-:]?[0-9a-fA-F]{1})$`

func validateIdentifierFlag(identifier string) (bool, error) {
	re, err := regexp.Compile(identifierRegexp)
	if err != nil {
		return false, err
	}
	valid := re.MatchString(identifier)
	return valid, nil
}
