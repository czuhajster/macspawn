package command

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/czuhajster/macspawn/internal/address"
	"github.com/czuhajster/macspawn/internal/format"
)

var (
	separator        string
	scope            string
	addressType      string
	identifierString string
	macAddress       string
	rootCmd          = &cobra.Command{
		Use:   "macspawn",
		Short: "MACSpawn is a MAC address generator.",
		RunE: func(cmd *cobra.Command, args []string) error {
			addressFormat, e := format.GetFormat(separator)
			if e != nil {
				return e
			}
			local, scopeError := address.CheckAddressScope(scope)
			if scopeError != nil {
				return scopeError
			}
			individual, typeError := address.CheckAddressType(addressType)
			if typeError != nil {
				return typeError
			}
			validIdentifier, validationErr := validateIdentifierFlag(identifierString)
			if validationErr != nil {
				return validationErr
			}
			if !validIdentifier {
				return errors.New("Invalid identifier.")
			}
			identifier, identifierError := processIdentifier(identifierString)
			if identifierError != nil {
				return errors.New("Invalid identifier.")
			}
			options := address.NewMACAddressOptions(local, individual, identifier)
			address := address.GenerateMACAddress(options)
			format.PrintMAC(address, addressFormat)
			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&separator, "separator", "s", ":", "Separator of the address bytes.")
	rootCmd.Flags().StringVarP(&scope, "scope", "c", "local", "Scope of the MAC address: local or universal.")
	rootCmd.Flags().StringVarP(&addressType, "type", "t", "individual", "Type of the MAC address: individual or group.")
	rootCmd.Flags().StringVarP(&identifierString, "identifier", "i", "", "Identifier. A 24- (MA-L), 28- (MA-M), or 36-bit (MA-S) hex number.")
	rootCmd.Flags().StringVarP(&macAddress, "mac-address", "m", "", "MAC Address. Used to generate an EUI-64 identifier.")
}

func Execute() error {
	return rootCmd.Execute()
}
