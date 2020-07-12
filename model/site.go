package model

type Site struct {
	Site    string  `json:"site"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
