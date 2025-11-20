package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Foo holds the schema definition for the Foo entity.
type Foo struct {
	ent.Schema
}

// Fields of the Foo.
func (Foo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
	}
}

// Edges of the Foo.
func (Foo) Edges() []ent.Edge {
	return nil
}
