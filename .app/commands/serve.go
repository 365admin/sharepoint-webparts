/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"log"

	"github.com/365admin/sharepoint-webparts/app/restapi"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve [[service]]",
	Short: "Starts the Koksmat server",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			restapi.Serve()
			return
		}
		service := args[0]
		switch service {

		default:

			log.Fatalln("Unknown service", service)
			return
		}
		//webserver.Run()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
