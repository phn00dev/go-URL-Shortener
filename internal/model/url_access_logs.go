package model

import "time"

type UrlAccessLog struct {
	ID          int       `json:"id"`
	UrlID       int       `json:"url_id"`      // Baglanşykly URL
	Accessed_At time.Time `json:"accessed_at"` // Girilen wagt
	IpAddress   string    `json:"ip_address"`  // IP Salgy
	UserAgent   string    `json:"user_agent"`  // Brauzer maglumatlary
}
