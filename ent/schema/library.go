package schema

import "entgo.io/ent"

// Library holds the schema definition for the Library entity.
type Library struct {
	ent.Schema
}

// Fields of the Library.
func (Library) Fields() []ent.Field {
	return nil
}

// Edges of the Library.
func (Library) Edges() []ent.Edge {
	return nil
}
