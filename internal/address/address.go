package address

import (
	"crypto/rand"
	"errors"
	"math"
)

type MACAddress [6]byte

type MACAddressOptions struct {
	local      bool
	individual bool
	identifier []byte
}

const (
	localScopeBitmask     byte = 2
	universalScopeBitmask byte = 253
	unicastBitmask        byte = 254
	multicastBitmask      byte = 1
)

func GenerateIdentifier(lengthBits uint8) []byte {
	lengthBytes := uint8(math.Ceil(float64(lengthBits) / 8))
	var identifier []byte
	identifier = make([]byte, lengthBytes)
	_, err := rand.Read(identifier)
	if err != nil {
		panic(err)
	}
	if lengthBits%8 > 0 {
		remainingBits := 8 - (lengthBits % 8)
		bitmask := GenerateBitmask(remainingBits, false)
		identifier[len(identifier)-1] &= bitmask
	}
	return identifier
}

func GenerateNICSpecificBytes(lengthBits uint8) []byte {
	lengthBytes := uint8(math.Ceil(float64(lengthBits) / 8))
	var nicSpecificBytes []byte
	nicSpecificBytes = make([]byte, lengthBytes)
	_, err := rand.Read(nicSpecificBytes)
	if err != nil {
		panic(err)
	}
	if lengthBits%8 > 0 {
		remainingBits := 8 - (lengthBits % 8)
		bitmask := GenerateBitmask(remainingBits, true)
		nicSpecificBytes[0] &= bitmask
	}
	return nicSpecificBytes
}

func GenerateMACAddress(options *MACAddressOptions) *MACAddress {
	var identifierLength uint8
	if options.identifier != nil {
		identifierLength = uint8(len(options.identifier))
	} else {
		identifierLength = 24
	}
	identifier := GenerateIdentifier(identifierLength)
	if options.local == true {
		identifier[0] |= localScopeBitmask
	} else {
		identifier[0] &= universalScopeBitmask
	}
	if options.individual == true {
		identifier[0] &= unicastBitmask
	} else {
		identifier[0] |= multicastBitmask
	}
	nicSpecificBytes := GenerateNICSpecificBytes(48 - identifierLength)
	address := MACAddress{identifier[0], identifier[1], identifier[2], nicSpecificBytes[0], nicSpecificBytes[1], nicSpecificBytes[2]}
	return &address
}

func GenerateBitmask(shifts uint8, reverse bool) byte {
	var bitmask byte = 255
	if reverse {
		bitmask >>= shifts
	} else {
		bitmask <<= shifts
	}
	return bitmask
}

func NewMACAddressOptions(local bool, individual bool, identifier []byte) *MACAddressOptions {
	options := MACAddressOptions{
		local:      local,
		individual: individual,
		identifier: identifier,
	}
	return &options
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
