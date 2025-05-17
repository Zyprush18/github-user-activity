package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Zyprush18/github-user-activity/models"
)

type Repsitory struct {
	Name string
}


func Activity(name string) {
	var data []models.Datas

	url := fmt.Sprintf("https://api.github.com/users/%s/events", name)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := json.Unmarshal(respdata, &data); err != nil {
		fmt.Println(err.Error())
	}

	var datas []Repsitory
	var isidata Repsitory

	for _, g := range data {
		if g.Type == "PushEvent" && g.Payload.Commits != nil {
			nama_repodia := g.Repo.Name
			if nama_repodia != isidata.Name {
				isidata.Name = nama_repodia
				datas = append(datas, isidata)
			}
		}
	}

	fmt.Println(datas)
}
