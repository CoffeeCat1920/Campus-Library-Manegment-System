package modals

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
)

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
