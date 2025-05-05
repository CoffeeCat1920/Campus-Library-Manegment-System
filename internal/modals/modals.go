package modals

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
)

type UserType int

const (
	Student = iota
	Admin
)

func (ut UserType) String() string {
	names := [...]string{"Student", "Admin"}
	if ut < Student || ut > Admin {
		return "Invalid"
	} else {
		return names[ut]
	}
}

func generateSessionToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func getTime(input string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, input)
	if err != nil {
		log.Print("Error parsing time:", err)
		return time.Now(), err
	}
	return t, nil
}
