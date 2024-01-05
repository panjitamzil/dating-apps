package model

import "time"

type User struct {
	Id           string     `json:"id"`
	Email        string     `json:"email"`
	Password     string     `json:"password,omitempty"`
	Fullname     *string    `json:"fullname"`
	DOB          *time.Time `json:"dob" time_format:"2006-01-02" time_utc:"1"`
	Occupation   *string    `json:"occupation"`
	Subscription string     `json:"subscription"`
	CreatedAt    time.Time  `json:"created_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
