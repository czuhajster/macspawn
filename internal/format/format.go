package format

import (
	"errors"
	"fmt"

	"github.com/czuhajster/macspawn/internal/address"
)

type macAddressFormat string

const (
	ColonFormat  macAddressFormat = "%02X:%02X:%02X:%02X:%02X:%02X\n"
	HyphenFormat macAddressFormat = "%02X-%02X-%02X-%02X-%02X-%02X\n"
	SimpmeFormat macAddressFormat = "%02X%02X%02X%02X%02X%02X\n"
)

func PrintMAC(x *address.MACAddress, format macAddressFormat) {
	fmt.Printf(string(format), x[0], x[1], x[2], x[3], x[4], x[5])
}

func GetFormat(separator string) (macAddressFormat, error) {
	switch separator {
	case ":":
		return ColonFormat, nil
	case "-":
		return HyphenFormat, nil
	default:
		return "", errors.New("Invalid separator.")
	}
}
