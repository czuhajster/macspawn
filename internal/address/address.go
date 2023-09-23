package address

import (
    "math/rand"
)

type MACAddress [6]uint8

func GenerateOUI() *[3]uint8 {
    var OUI [3]uint8
    for i := 0; i < 3; i++ {
        OUI[i] = uint8(rand.Intn(255))
    }
    return &OUI
}

func GenerateNICSpecificBytes() *[3]uint8 {
    var nicSpecificBytes [3]uint8
    for i := 0; i < 3; i++ {
        nicSpecificBytes[i] = uint8(rand.Intn(255))
    }
    return &nicSpecificBytes
}

func GenerateMACAddress() *MACAddress {
    oui := GenerateOUI()
    nicSpecificBytes := GenerateNICSpecificBytes()
    address := MACAddress{oui[0], oui[1], oui[2], nicSpecificBytes[0], nicSpecificBytes[1], nicSpecificBytes[2]}
    return &address
}
