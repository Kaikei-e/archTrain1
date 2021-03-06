// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFailedLoginAttempts holds the string denoting the failed_login_attempts field in the database.
	FieldFailedLoginAttempts = "failed_login_attempts"
	// FieldIsBlocked holds the string denoting the is_blocked field in the database.
	FieldIsBlocked = "is_blocked"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldFailedLoginAttempts,
	FieldIsBlocked,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultFailedLoginAttempts holds the default value on creation for the "failed_login_attempts" field.
	DefaultFailedLoginAttempts int
	// DefaultIsBlocked holds the default value on creation for the "is_blocked" field.
	DefaultIsBlocked bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultCreatedAt holds the default value on update for the "created_at" field.
	UpdateDefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
