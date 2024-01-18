/*---
title: cmd "meeting-infrastructure"
--*/

package commands

import (
	"fmt"
	"os"

	"github.com/365admin/sharepoint-webparts/matrix"
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
	var JobsCmd cobra.Command = cobra.Command{
		Use:   "jobs",
		Short: "Jobs commands for SharePoint Webparts",
		Long:  `Here you can find all the commands that are relevant for the management of jobs`,
	}

	var MatrixPageCmd cobra.Command = cobra.Command{
		Use:   "matrixpage",
		Short: "MatrixPage commands for SharePoint Webparts",
		Long:  `Here you can find all the commands that are relevant for the management of matrix pages`,
		Run: func(cmd *cobra.Command, args []string) {
			filepath, err := matrix.WriteMatrix()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Matrix page written to %s\n", filepath)

		},
	}

	JobsCmd.AddCommand(&MatrixPageCmd)
	RootCmd.AddCommand(&JobsCmd)

}
