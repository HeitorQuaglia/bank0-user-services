package models

import "time"

type User struct {
	ID           *int
	Username     *string
	PasswordHash *string
	PasswordSalt *string
	Deleted      *bool
	DeletedAt    *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time

	// Emails
	Emails        []UserEmail
	PersonalInfos []PersonalInfo
}

type UserEmail struct {
	ID        *int
	UserID    *int
	Primary   *bool
	Email     *string
	Verified  *bool
	Deleted   *bool
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type PersonalInfo struct {
	ID        *int
	UserID    *int
	GivenName *string
	Surname   *string
	FederalID *string
	BirthDate *time.Time
	Relative  *bool
	Deleted   *bool
	DeletedAt *time.Time
	Address   []Address
}

type Address struct {
	ID             *int
	PersonalInfoID *int
	StreetName     *string
	StreetNumber   *string
	Complement     *string
	Neighborhood   *string
	City           *string
	Province       *string
	Country        *string
	PostalCode     *string
	Primary        *bool
	Deleted        *bool
	DeletedAt      *time.Time
}

type UserPasswordReset struct {
	ID        *int
	UserID    *int
	Token     *string
	Used      *bool
	UsedAt    *time.Time
	CreatedAt *time.Time
}

type UserPasswordChangeLog struct {
	ID           int
	UserID       int
	PasswordHash string
	PasswordSalt string
	CreatedAt    time.Time
}

type Filter struct {
	Offset  *int
	Limit   *int
	Search  *string
	From    *time.Time
	To      *time.Time
	Deleted *bool
}
