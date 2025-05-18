package main

import (
	"fmt"
	"os"

	"github.com/Zyprush18/github-user-activity/service"
)

func main() {
	argument := os.Args

	if len(argument) > 1 {
		fetch, err := service.GetDataFromApi(argument[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data, err := service.ParseDataToSlice(fetch)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		service.Activity(data)

	} else {
		fmt.Println("please enter ./github-activity <your_github_username>")
	}

}
