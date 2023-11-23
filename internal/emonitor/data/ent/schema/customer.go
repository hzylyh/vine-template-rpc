package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_customer"},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return nil
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
