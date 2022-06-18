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
