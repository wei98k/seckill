// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/peter-wow/seckill/app/goods/service/internal/data/ent/goods"
	"github.com/peter-wow/seckill/app/goods/service/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	goodsFields := schema.Goods{}.Fields()
	_ = goodsFields
	// goodsDescCreatedAt is the schema descriptor for created_at field.
	goodsDescCreatedAt := goodsFields[3].Descriptor()
	// goods.DefaultCreatedAt holds the default value on creation for the created_at field.
	goods.DefaultCreatedAt = goodsDescCreatedAt.Default.(func() time.Time)
	// goodsDescUpdatedAt is the schema descriptor for updated_at field.
	goodsDescUpdatedAt := goodsFields[4].Descriptor()
	// goods.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	goods.DefaultUpdatedAt = goodsDescUpdatedAt.Default.(func() time.Time)
}
