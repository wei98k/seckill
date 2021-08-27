package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// OrderGoods holds the schema definition for the OrderGoods entity.
type OrderGoods struct {
	ent.Schema
}

// Fields of the OrderGoods.
func (OrderGoods) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("order_id"),
		field.Int64("goods_id"),
		field.String("goods_title"),
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

// Edges of the OrderGoods.
func (OrderGoods) Edges() []ent.Edge {
	return nil
}
