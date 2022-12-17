package entity

import "time"

type AdminUser struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"username,omitempty"`
	FullName  string    `json:"fullName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	IsActive  string    `json:"isActive,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
