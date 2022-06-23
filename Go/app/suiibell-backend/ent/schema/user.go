package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Unique().NotEmpty().StructTag(`gqlgen:"email",required:"true"`),
		field.String("password").NotEmpty().StructTag(`gqlgen:"password",required:"true"`),
		field.Int("failed_login_attempts").Default(0),
		field.Bool("is_blocked").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now),
		field.Int("deleted_at").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
