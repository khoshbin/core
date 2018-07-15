package main

// User represents an item from the "users" table, column
// names are mapped to Go values.
type User struct {
	// Map the "id" column to the ID field. Only exported
	// fields can be mapped to database columns.
	ID            uint    `db:"id"`
	CreatedAt     *string `db:"created_at"`
	DeletedAt     *string `db:"deleted_at"`
	Name          string  `db:"name" json:"name"`
	Email         string  `db:"email" json:"email"`
	Mobile        string  `db:"mobile" json:"mobile"`
	Password      string  `db:"password"`
	Status        string  `db:"status"`
	Type          string  `db:"type"`
	RememberToken string  `db:"remember_token"`
	SmsToken      uint    `db:"sms_token"`
}
