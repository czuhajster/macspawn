package address

import (
    "math/rand"
)

type MACAddress [6]byte

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

func GenerateMACAddress() *MACAddress {
    oui := GenerateOUI()
    nicSpecificBytes := GenerateNICSpecificBytes()
    address := MACAddress{oui[0], oui[1], oui[2], nicSpecificBytes[0], nicSpecificBytes[1], nicSpecificBytes[2]}
    return &address
}
