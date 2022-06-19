package defs

// Requests
type UserCredential struct {
	UserName string `json:"user_name"`
	PWD      string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type Session struct {
	Username string
	TTL      int64
}
