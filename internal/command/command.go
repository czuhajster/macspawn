package command

import (
    "errors"

    "github.com/spf13/cobra"

    "github.com/czuhajster/macspawn/internal/address"
    "github.com/czuhajster/macspawn/internal/format"
)

var (
    separator string
    rootCmd = &cobra.Command{
      Use:   "macspawn",
      Short: "MACSpawn is a MAC address generator.",
      RunE: func(cmd *cobra.Command, args []string) error {
          switch separator {
          case ":", "-":
          default:
              return errors.New("Invalid separator.")
          }
          address := address.GenerateMACAddress()
          format.PrintMAC(address, format.ColonFormat)
          return nil
      },
    }
)

func init() {
    rootCmd.Flags().StringVar(&separator, "separator", ":", "Separator of the address bytes.")
}

func Execute() error {
    return rootCmd.Execute()
}
