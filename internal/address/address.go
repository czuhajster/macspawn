package address

import (
    "math/rand"
    "errors"
)

type MACAddress [6]byte

const (
    localScopeBitmask byte = 2
    universalScopeBitmask byte = 253
)

func GenerateOUI() *[3]byte {
    var OUI [3]byte
    for i := 0; i < 3; i++ {
        OUI[i] = byte(rand.Intn(255))
    }
    return &OUI
}

func GenerateNICSpecificBytes() *[3]byte {
    var nicSpecificBytes [3]byte
    for i := 0; i < 3; i++ {
        nicSpecificBytes[i] = byte(rand.Intn(255))
    }
    return &nicSpecificBytes
}

func GenerateMACAddress(local bool) *MACAddress {
    oui := GenerateOUI()
    if local == true {
        oui[0] = oui[0] | localScopeBitmask
    } else {
        oui[0] = oui[0] & universalScopeBitmask
    }
    nicSpecificBytes := GenerateNICSpecificBytes()
    address := MACAddress{oui[0], oui[1], oui[2], nicSpecificBytes[0], nicSpecificBytes[1], nicSpecificBytes[2]}
    return &address
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
