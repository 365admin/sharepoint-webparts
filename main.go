package main

import (
	"github.com/365admin/sharepoint-webparts/app/commands"
	"github.com/365admin/sharepoint-webparts/app/config"
)

func main() {

	config.Setup(".env")
	commands.Execute()
}
