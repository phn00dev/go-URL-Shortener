package generateshorturl

import (
	"math/rand"
	"time"
)

func GenerateShortUrl(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	shortUrl := make([]rune, length)
	for i := range shortUrl {
		shortUrl[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortUrl)
}
