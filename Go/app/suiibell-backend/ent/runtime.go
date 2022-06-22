// Code generated by entc, DO NOT EDIT.

package ent

import (
	"suiibell/ent/schema"
	"suiibell/ent/user"
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
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescFailedLoginAttempts is the schema descriptor for failed_login_attempts field.
	userDescFailedLoginAttempts := userFields[3].Descriptor()
	// user.DefaultFailedLoginAttempts holds the default value on creation for the failed_login_attempts field.
	user.DefaultFailedLoginAttempts = userDescFailedLoginAttempts.Default.(int)
	// userDescIsBlocked is the schema descriptor for is_blocked field.
	userDescIsBlocked := userFields[4].Descriptor()
	// user.DefaultIsBlocked holds the default value on creation for the is_blocked field.
	user.DefaultIsBlocked = userDescIsBlocked.Default.(bool)
}
