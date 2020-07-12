package model

import "github.com/sirupsen/logrus"

type Site struct {
	Site    string  `json:"site"`
	Entries []Entry `json:"entries"`
}

func (s *Site) AddItem(username, password string) {
	found := false
	var i int

	for row, next := range s.Entries {
		if next.Username == username {
			i = row
			found = true
			break
		}
	}

	if found {
		logrus.Println("set old")
		s.Entries[i].Password = password
		return
	}

	logrus.Println("set new")
	s.Entries = append(s.Entries, Entry{
		Username: username,
		Password: password,
	})

}

type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
