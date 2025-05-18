package models

import "time"

type Datas struct {
	Id string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Actor Actor `json:"actor,omitempty"`
	Repo Repo `json:"repo,omitempty"`
	Payload Payload `json:"payload,omitempty"`
	Public bool `json:"public,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Actor struct {
	Id int `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Display string `json:"display_login,omitempty"`
	GravatarId string `json:"gravatar_id,omitempty"`
	Url string `json:"url,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type Repo struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
}

type Payload struct {
	Action string `json:"action"`
	RepositiryId int `json:"repository_id,omitempty"`
	PushId int `json:"push_id,omitempty"`
	Size int `json:"size,omitempty"`
	DistinctSize int `json:"distinct_size,omitempty"`
	Ref string `json:"ref,omitempty"`
	RefType string `json:"ref_type,omitempty"`
	MasterBranch string `json:"master_branch,omitempty"`
	Description string `json:"description,omitempty"`
	PusherType string `json:"pusher_type,omitempty"`
	Head string `json:"head,omitempty"`
	Before string `json:"before,omitempty"`
	Commits []Commit `json:"commits,omitempty"`
}

type Commit struct {
	Sha string `json:"sha,omitempty"`
	Author Author `json:"author,omitempty"`
	Message string `json:"message,omitempty"`
	Distinct bool `json:"distinct,omitempty"`
	Url string `json:"url,omitempty"`
}

type Author struct {
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}