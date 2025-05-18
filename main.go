package main

import (
	"fmt"
	"os"

	"github.com/Zyprush18/github-user-activity/service"
)

func main() {
	argument := os.Args

	
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
	

}