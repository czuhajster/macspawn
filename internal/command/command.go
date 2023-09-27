package command

import (
    "github.com/spf13/cobra"

    "github.com/czuhajster/macspawn/internal/address"
    "github.com/czuhajster/macspawn/internal/format"
)

var (
    separator string
    scope string
    addressType string
    rootCmd = &cobra.Command{
      Use:   "macspawn",
      Short: "MACSpawn is a MAC address generator.",
      RunE: func(cmd *cobra.Command, args []string) error {
          addressFormat, e := format.GetFormat(separator)
          if e != nil {
              return e
          }
          local, scopeError  := address.CheckAddressScope(scope)
          if scopeError != nil {
              return scopeError
          }
          individual, typeError := address.CheckAddressType(addressType)
          if typeError != nil {
              return typeError
          }
          var identifier []byte
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
}

func Execute() error {
    return rootCmd.Execute()
}
