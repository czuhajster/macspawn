package address

import (
    "crypto/rand"
    "errors"
)

type MACAddress [6]byte

const (
    localScopeBitmask byte = 2
    universalScopeBitmask byte = 253
    unicastBitmask byte = 254
    multicastBitmask byte = 1
)

func GenerateOUI() *[3]byte {
    var oui [3]byte
    _, err := rand.Read(oui[:])
    if err != nil {
        panic(err)
    }
    return &oui
}

func GenerateNICSpecificBytes() *[3]byte {
    var nicSpecificBytes [3]byte
    _, err := rand.Read(nicSpecificBytes[:])
    if err != nil {
        panic(err)
    }
    return &nicSpecificBytes
}

func GenerateMACAddress(local bool, individual bool) *MACAddress {
    oui := GenerateOUI()
    if local == true {
        oui[0] = oui[0] | localScopeBitmask
    } else {
        oui[0] = oui[0] & universalScopeBitmask
    }
    if individual == true {
        oui[0] = oui[0] & unicastBitmask
    } else {
        oui[0] = oui[0] | multicastBitmask
    }
    nicSpecificBytes := GenerateNICSpecificBytes()
    address := MACAddress{oui[0], oui[1], oui[2], nicSpecificBytes[0], nicSpecificBytes[1], nicSpecificBytes[2]}
    return &address
}

func CheckAddressType(addressType string) (bool, error) {
    switch addressType {
    case "individual", "i":
        return true, nil
    case "group", "g":
        return false, nil
    default:
        return false, errors.New("Unrecognised address type.")
    }
}

func CheckAddressScope(scope string) (bool, error) {
    switch scope {
    case "local", "l":
        return true, nil
    case "universal", "u":
        return false, nil
    default:
        return false, errors.New("Unrecognised scope.")
    }
}
