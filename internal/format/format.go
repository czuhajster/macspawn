package format

import (
    "fmt"
)

type macAddressFormat string

const (
    ColonFormat macAddressFormat = "%02X:%02X:%02X:%02X:%02X:%02X\n"
    HyphenFormat macAddressFormat = "%02X-%02X-%02X-%02X-%02X-%02X\n"
    SimpmeFormat macAddressFormat = "%02X%02X%02X%02X%02X%02X\n"
)

func PrintMAC(x *[6]uint8, format macAddressFormat) {
    fmt.Printf(string(format), x[0], x[1], x[2], x[3], x[4], x[5])
}
