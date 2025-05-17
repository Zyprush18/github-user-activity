package main

import (
	"os"

	"github.com/Zyprush18/github-user-activity/service"
)

func main() {
	argument := os.Args
	service.Activity(argument[1])

}