// Code generated by entc, DO NOT EDIT.

package ent

import (
	"suiibell/ent/schema"
	"suiibell/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func([]byte) error)
	// userDescFailedLoginAttempts is the schema descriptor for failed_login_attempts field.
	userDescFailedLoginAttempts := userFields[3].Descriptor()
	// user.DefaultFailedLoginAttempts holds the default value on creation for the failed_login_attempts field.
	user.DefaultFailedLoginAttempts = userDescFailedLoginAttempts.Default.(int)
	// userDescIsBlocked is the schema descriptor for is_blocked field.
	userDescIsBlocked := userFields[4].Descriptor()
	// user.DefaultIsBlocked holds the default value on creation for the is_blocked field.
	user.DefaultIsBlocked = userDescIsBlocked.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// user.UpdateDefaultCreatedAt holds the default value on update for the created_at field.
	user.UpdateDefaultCreatedAt = userDescCreatedAt.UpdateDefault.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userFields[7].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(int)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
