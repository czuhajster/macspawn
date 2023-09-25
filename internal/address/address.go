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

func GenerateIdentifier() *[3]byte {
    var identifier [3]byte
    _, err := rand.Read(identifier[:])
    if err != nil {
        panic(err)
    }
    return &identifier
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
    identifier := GenerateOUI()
    if local == true {
        identifier[0] |= localScopeBitmask
    } else {
        identifier[0] &= universalScopeBitmask
    }
    if individual == true {
        identifier[0] &= unicastBitmask
    } else {
        identifier[0] |= multicastBitmask
    }
    nicSpecificBytes := GenerateNICSpecificBytes()
    address := MACAddress{identifier[0], identifier[1], identifier[2], nicSpecificBytes[0], nicSpecificBytes[1], nicSpecificBytes[2]}
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
