package command

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/czuhajster/macspawn/internal/address"
	"github.com/czuhajster/macspawn/internal/format"
	"github.com/czuhajster/macspawn/internal/registry"
)

var (
	separator        string
	scope            string
	addressType      string
	identifierString string
	macAddress       string
	registryFile     string
	name             string
	rootCmd          = &cobra.Command{
		Use:   "macspawn",
		Short: "MACSpawn is a MAC address generator.",
		RunE: func(cmd *cobra.Command, args []string) error {
			var identifier []byte = nil
			var identifierError error
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

			identifierSet := cmd.Flags().Lookup("identifier").Changed
			if identifierSet {
				validIdentifier, validationErr := validateIdentifierFlag(identifierString)
				if validationErr != nil {
					return validationErr
				}
				if !validIdentifier {
					return errors.New("Invalid identifier.")
				}
				identifier, identifierError = processIdentifier(identifierString)
				if identifierError != nil {
					return errors.New("Invalid identifier.")
				}
			}
			if len(registryFile) > 0 && len(name) > 0 {
				record := registry.FindRecord(registryFile, name)
				if record != nil {
					var err error
					identifier, err = processIdentifier(record[1])
					if err != nil {
						return err
					}
				}
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
	rootCmd.Flags().StringVarP(&registryFile, "file", "f", "", "Registry file.")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Company name.")
}

func Execute() error {
	return rootCmd.Execute()
}
