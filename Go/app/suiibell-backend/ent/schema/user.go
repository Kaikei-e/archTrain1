package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New()).Unique(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.Int("failed_login_attempts").Default(0),
		field.Bool("is_blocked").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
