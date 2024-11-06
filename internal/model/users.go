package model

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // Şifreyi dışarıya göstermemek için
	CreatedAt    string `json:"created_at"`
	UserUrls     []Url  `json:"user_urls" `
	UrlCount     int    `json:"url_count"` // Burada gorm:"-" olmamalıdır
}
