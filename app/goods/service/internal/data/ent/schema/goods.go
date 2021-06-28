package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Goods holds the schema definition for the Goods entity.
type Goods struct {
	ent.Schema
}

// Fields of the Goods.
func (Goods) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Goods.
func (Goods) Edges() []ent.Edge {
	return nil
}
