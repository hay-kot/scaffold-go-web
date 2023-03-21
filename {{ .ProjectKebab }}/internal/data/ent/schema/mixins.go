package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// BaseMixin defines a mixin for the id, created_at, and updated_at fields.
type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// TenantMixin defines a mixin for the tenant_id field and edge. Schemas
// that use this mixin will have a tenant_id field and a
// tenant edge that is required.
type TenantMixin struct {
	ref string
	mixin.Schema
}

func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("tenant_id", uuid.UUID{}),
	}
}

func (o TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Unique().
			Field("tenant_id").
			Ref(o.ref).
			Required(),
	}
}