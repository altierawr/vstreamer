package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ent/hook"
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

func (Library) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.LibraryFunc(func(ctx context.Context, m *gen.LibraryMutation) (ent.Value, error) {
					fmt.Println("created lib")
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate,
		),
	}
}
