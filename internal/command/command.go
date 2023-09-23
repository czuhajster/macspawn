package command

import (
    "github.com/spf13/cobra"
)

var (
    rootCmd = &cobra.Command{
      Use:   "macspawn",
      Short: "MACSpawn is a MAC address generator.",
    },
)
