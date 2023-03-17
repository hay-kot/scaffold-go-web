{{- if .Scaffold.use_database -}}
// Package schema defines the schema for the entities in the database.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{ref: "users"},
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			NotEmpty().
			MaxLen(100),
		field.String("last_name").
			NotEmpty().
			MaxLen(100),
		field.String("email").
			NotEmpty().
			MaxLen(255). // supposed max length of an email address is 254 characters
			Unique(),
		field.String("password_hash").
			NotEmpty().
			MaxLen(512),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
{{- end -}}