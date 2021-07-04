package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SeckillOrder holds the schema definition for the SeckillOrder entity.
type SeckillOrder struct {
	ent.Schema
}

// Fields of the SeckillOrder.
func (SeckillOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.Int64("order_id"),
		field.Int64("goods_id"),
	}
}

// Edges of the SeckillOrder.
func (SeckillOrder) Edges() []ent.Edge {
	return nil
}
