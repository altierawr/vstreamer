package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Library holds the schema definition for the Library entity.
type Library struct {
	ent.Schema
}

// Fields of the Library.
func (Library) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Library.
func (Library) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("videos", Video.Type),
	}
}

func (Library) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
