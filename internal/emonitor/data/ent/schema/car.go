package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_car"},
	}
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return nil
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return nil
}
