package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// SeckillGoods holds the schema definition for the SeckillGoods entity.
type SeckillGoods struct {
	ent.Schema
}

// Fields of the SeckillGoods.
func (SeckillGoods) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("goods_id"),
		field.Float("seckill_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}),
		field.Int64("stock_count"),
		field.Time("start_date"),
		field.Time("end_date"),
	}
}

// Edges of the SeckillGoods.
func (SeckillGoods) Edges() []ent.Edge {
	return nil
}
