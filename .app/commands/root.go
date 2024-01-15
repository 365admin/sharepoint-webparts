/*---
title: cmd "meeting-infrastructure"
--*/

package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "sharepoint-webparts",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
