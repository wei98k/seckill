// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/helloMJW/seckill/app/order/service/internal/data/ent/predicate"
	"github.com/helloMJW/seckill/app/order/service/internal/data/ent/seckillorder"
)

// SeckillOrderQuery is the builder for querying SeckillOrder entities.
type SeckillOrderQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SeckillOrder
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SeckillOrderQuery builder.
func (soq *SeckillOrderQuery) Where(ps ...predicate.SeckillOrder) *SeckillOrderQuery {
	soq.predicates = append(soq.predicates, ps...)
	return soq
}

// Limit adds a limit step to the query.
func (soq *SeckillOrderQuery) Limit(limit int) *SeckillOrderQuery {
	soq.limit = &limit
	return soq
}

// Offset adds an offset step to the query.
func (soq *SeckillOrderQuery) Offset(offset int) *SeckillOrderQuery {
	soq.offset = &offset
	return soq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (soq *SeckillOrderQuery) Unique(unique bool) *SeckillOrderQuery {
	soq.unique = &unique
	return soq
}

// Order adds an order step to the query.
func (soq *SeckillOrderQuery) Order(o ...OrderFunc) *SeckillOrderQuery {
	soq.order = append(soq.order, o...)
	return soq
}

// First returns the first SeckillOrder entity from the query.
// Returns a *NotFoundError when no SeckillOrder was found.
func (soq *SeckillOrderQuery) First(ctx context.Context) (*SeckillOrder, error) {
	nodes, err := soq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{seckillorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (soq *SeckillOrderQuery) FirstX(ctx context.Context) *SeckillOrder {
	node, err := soq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SeckillOrder ID from the query.
// Returns a *NotFoundError when no SeckillOrder ID was found.
func (soq *SeckillOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = soq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{seckillorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (soq *SeckillOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := soq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SeckillOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SeckillOrder entity is not found.
// Returns a *NotFoundError when no SeckillOrder entities are found.
func (soq *SeckillOrderQuery) Only(ctx context.Context) (*SeckillOrder, error) {
	nodes, err := soq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{seckillorder.Label}
	default:
		return nil, &NotSingularError{seckillorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (soq *SeckillOrderQuery) OnlyX(ctx context.Context) *SeckillOrder {
	node, err := soq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SeckillOrder ID in the query.
// Returns a *NotSingularError when exactly one SeckillOrder ID is not found.
// Returns a *NotFoundError when no entities are found.
func (soq *SeckillOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = soq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = &NotSingularError{seckillorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (soq *SeckillOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := soq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SeckillOrders.
func (soq *SeckillOrderQuery) All(ctx context.Context) ([]*SeckillOrder, error) {
	if err := soq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return soq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (soq *SeckillOrderQuery) AllX(ctx context.Context) []*SeckillOrder {
	nodes, err := soq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SeckillOrder IDs.
func (soq *SeckillOrderQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := soq.Select(seckillorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (soq *SeckillOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := soq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (soq *SeckillOrderQuery) Count(ctx context.Context) (int, error) {
	if err := soq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return soq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (soq *SeckillOrderQuery) CountX(ctx context.Context) int {
	count, err := soq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (soq *SeckillOrderQuery) Exist(ctx context.Context) (bool, error) {
	if err := soq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return soq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (soq *SeckillOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := soq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SeckillOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (soq *SeckillOrderQuery) Clone() *SeckillOrderQuery {
	if soq == nil {
		return nil
	}
	return &SeckillOrderQuery{
		config:     soq.config,
		limit:      soq.limit,
		offset:     soq.offset,
		order:      append([]OrderFunc{}, soq.order...),
		predicates: append([]predicate.SeckillOrder{}, soq.predicates...),
		// clone intermediate query.
		sql:  soq.sql.Clone(),
		path: soq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SeckillOrder.Query().
//		GroupBy(seckillorder.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (soq *SeckillOrderQuery) GroupBy(field string, fields ...string) *SeckillOrderGroupBy {
	group := &SeckillOrderGroupBy{config: soq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := soq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return soq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.SeckillOrder.Query().
//		Select(seckillorder.FieldUserID).
//		Scan(ctx, &v)
//
func (soq *SeckillOrderQuery) Select(field string, fields ...string) *SeckillOrderSelect {
	soq.fields = append([]string{field}, fields...)
	return &SeckillOrderSelect{SeckillOrderQuery: soq}
}

func (soq *SeckillOrderQuery) prepareQuery(ctx context.Context) error {
	for _, f := range soq.fields {
		if !seckillorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if soq.path != nil {
		prev, err := soq.path(ctx)
		if err != nil {
			return err
		}
		soq.sql = prev
	}
	return nil
}

func (soq *SeckillOrderQuery) sqlAll(ctx context.Context) ([]*SeckillOrder, error) {
	var (
		nodes = []*SeckillOrder{}
		_spec = soq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SeckillOrder{config: soq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, soq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (soq *SeckillOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := soq.querySpec()
	return sqlgraph.CountNodes(ctx, soq.driver, _spec)
}

func (soq *SeckillOrderQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := soq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (soq *SeckillOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   seckillorder.Table,
			Columns: seckillorder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: seckillorder.FieldID,
			},
		},
		From:   soq.sql,
		Unique: true,
	}
	if unique := soq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := soq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seckillorder.FieldID)
		for i := range fields {
			if fields[i] != seckillorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := soq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := soq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := soq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := soq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (soq *SeckillOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(soq.driver.Dialect())
	t1 := builder.Table(seckillorder.Table)
	selector := builder.Select(t1.Columns(seckillorder.Columns...)...).From(t1)
	if soq.sql != nil {
		selector = soq.sql
		selector.Select(selector.Columns(seckillorder.Columns...)...)
	}
	for _, p := range soq.predicates {
		p(selector)
	}
	for _, p := range soq.order {
		p(selector)
	}
	if offset := soq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := soq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SeckillOrderGroupBy is the group-by builder for SeckillOrder entities.
type SeckillOrderGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sogb *SeckillOrderGroupBy) Aggregate(fns ...AggregateFunc) *SeckillOrderGroupBy {
	sogb.fns = append(sogb.fns, fns...)
	return sogb
}

// Scan applies the group-by query and scans the result into the given value.
func (sogb *SeckillOrderGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sogb.path(ctx)
	if err != nil {
		return err
	}
	sogb.sql = query
	return sogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) StringsX(ctx context.Context) []string {
	v, err := sogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) StringX(ctx context.Context) string {
	v, err := sogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) IntsX(ctx context.Context) []int {
	v, err := sogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) IntX(ctx context.Context) int {
	v, err := sogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sogb.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sogb *SeckillOrderGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sogb *SeckillOrderGroupBy) BoolX(ctx context.Context) bool {
	v, err := sogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sogb *SeckillOrderGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sogb.fields {
		if !seckillorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sogb *SeckillOrderGroupBy) sqlQuery() *sql.Selector {
	selector := sogb.sql
	columns := make([]string, 0, len(sogb.fields)+len(sogb.fns))
	columns = append(columns, sogb.fields...)
	for _, fn := range sogb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sogb.fields...)
}

// SeckillOrderSelect is the builder for selecting fields of SeckillOrder entities.
type SeckillOrderSelect struct {
	*SeckillOrderQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sos *SeckillOrderSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sos.prepareQuery(ctx); err != nil {
		return err
	}
	sos.sql = sos.SeckillOrderQuery.sqlQuery(ctx)
	return sos.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sos *SeckillOrderSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sos.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sos *SeckillOrderSelect) StringsX(ctx context.Context) []string {
	v, err := sos.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sos.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sos *SeckillOrderSelect) StringX(ctx context.Context) string {
	v, err := sos.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sos *SeckillOrderSelect) IntsX(ctx context.Context) []int {
	v, err := sos.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sos.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sos *SeckillOrderSelect) IntX(ctx context.Context) int {
	v, err := sos.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sos *SeckillOrderSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sos.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sos.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sos *SeckillOrderSelect) Float64X(ctx context.Context) float64 {
	v, err := sos.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sos.fields) > 1 {
		return nil, errors.New("ent: SeckillOrderSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sos *SeckillOrderSelect) BoolsX(ctx context.Context) []bool {
	v, err := sos.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sos *SeckillOrderSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sos.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{seckillorder.Label}
	default:
		err = fmt.Errorf("ent: SeckillOrderSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sos *SeckillOrderSelect) BoolX(ctx context.Context) bool {
	v, err := sos.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sos *SeckillOrderSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sos.sqlQuery().Query()
	if err := sos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sos *SeckillOrderSelect) sqlQuery() sql.Querier {
	selector := sos.sql
	selector.Select(selector.Columns(sos.fields...)...)
	return selector
}
