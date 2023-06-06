package user_model

import "time"

type User struct {
	Id           int64     `json:"id" db:"id"`
	FirstName    string    `json:"firstName" db:"first_name"`
	LastName     string    `json:"lastName" db:"last_name"`
	Username     string    `json:"username" db:"username"`
	EmailAddress string    `json:"emailAddress" db:"email_address"`
	Password     string    `json:"password" db:"password"`
	BirthDate    time.Time `json:"birthDate" db:"birth_date"`
	Gender       string    `json:"gender" db:"gender"`
	CreatedAt    time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	DeletedAt    time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type RegisterReq struct {
	FirstName    string    `json:"firstName" db:"first_name"`
	LastName     string    `json:"lastName" db:"last_name"`
	Username     string    `json:"username" db:"username"`
	EmailAddress string    `json:"emailAddress" db:"email_address"`
	Password     string    `json:"password" db:"password"`
	BirthDate    time.Time `json:"birthDate" db:"birth_date"`
	Gender       string    `json:"gender" db:"gender"`
}

type LoginReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Resp struct {
	Id           int64     `json:"id"  db:"id"`
	FirstName    string    `json:"firstName" db:"first_name"`
	LastName     string    `json:"lastName" db:"last_name"`
	Username     string    `json:"username" db:"username"`
	EmailAddress string    `json:"emailAddress" db:"email_address"`
	BirthDate    time.Time `json:"birthDate" db:"birth_date"`
	Gender       string    `json:"gender" db:"gender"`
}
