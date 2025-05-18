package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Zyprush18/github-user-activity/models"
)

func GetDataFromApi(name string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respdata, nil
}

func ParseDataToSlice(dataByte []byte) ([]models.Datas, error){
	var data []models.Datas

	if err := json.Unmarshal(dataByte, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func Activity(data []models.Datas) {
	isExistRepo := map[string]bool{}
	countCommitRepo := map[string]int{}

	for _, d := range data {
		if d.Payload.Commits != nil && d.Type == "PushEvent" {
			if !isExistRepo[d.Repo.Name] {
				isExistRepo[d.Repo.Name] = true
			}
			countCommitRepo[d.Repo.Name] += len(d.Payload.Commits)
		}
	}

	for i, v := range countCommitRepo {
		fmt.Printf("- Pushed %d to %s \n",v,i)
	}

}
