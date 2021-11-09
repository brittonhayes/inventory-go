// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-farms/inventory/ent/location"
	"github.com/open-farms/inventory/ent/predicate"
	"github.com/open-farms/inventory/ent/vehicle"
)

// VehicleUpdate is the builder for updating Vehicle entities.
type VehicleUpdate struct {
	config
	hooks    []Hook
	mutation *VehicleMutation
}

// Where appends a list predicates to the VehicleUpdate builder.
func (vu *VehicleUpdate) Where(ps ...predicate.Vehicle) *VehicleUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetCreateTime sets the "create_time" field.
func (vu *VehicleUpdate) SetCreateTime(t time.Time) *VehicleUpdate {
	vu.mutation.SetCreateTime(t)
	return vu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vu *VehicleUpdate) SetNillableCreateTime(t *time.Time) *VehicleUpdate {
	if t != nil {
		vu.SetCreateTime(*t)
	}
	return vu
}

// SetUpdateTime sets the "update_time" field.
func (vu *VehicleUpdate) SetUpdateTime(t time.Time) *VehicleUpdate {
	vu.mutation.SetUpdateTime(t)
	return vu
}

// SetMake sets the "make" field.
func (vu *VehicleUpdate) SetMake(s string) *VehicleUpdate {
	vu.mutation.SetMake(s)
	return vu
}

// SetModel sets the "model" field.
func (vu *VehicleUpdate) SetModel(s string) *VehicleUpdate {
	vu.mutation.SetModel(s)
	return vu
}

// SetHours sets the "hours" field.
func (vu *VehicleUpdate) SetHours(i int64) *VehicleUpdate {
	vu.mutation.ResetHours()
	vu.mutation.SetHours(i)
	return vu
}

// SetNillableHours sets the "hours" field if the given value is not nil.
func (vu *VehicleUpdate) SetNillableHours(i *int64) *VehicleUpdate {
	if i != nil {
		vu.SetHours(*i)
	}
	return vu
}

// AddHours adds i to the "hours" field.
func (vu *VehicleUpdate) AddHours(i int64) *VehicleUpdate {
	vu.mutation.AddHours(i)
	return vu
}

// SetYear sets the "year" field.
func (vu *VehicleUpdate) SetYear(i int64) *VehicleUpdate {
	vu.mutation.ResetYear()
	vu.mutation.SetYear(i)
	return vu
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (vu *VehicleUpdate) SetNillableYear(i *int64) *VehicleUpdate {
	if i != nil {
		vu.SetYear(*i)
	}
	return vu
}

// AddYear adds i to the "year" field.
func (vu *VehicleUpdate) AddYear(i int64) *VehicleUpdate {
	vu.mutation.AddYear(i)
	return vu
}

// ClearYear clears the value of the "year" field.
func (vu *VehicleUpdate) ClearYear() *VehicleUpdate {
	vu.mutation.ClearYear()
	return vu
}

// SetActive sets the "active" field.
func (vu *VehicleUpdate) SetActive(b bool) *VehicleUpdate {
	vu.mutation.SetActive(b)
	return vu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (vu *VehicleUpdate) SetNillableActive(b *bool) *VehicleUpdate {
	if b != nil {
		vu.SetActive(*b)
	}
	return vu
}

// SetPower sets the "power" field.
func (vu *VehicleUpdate) SetPower(s string) *VehicleUpdate {
	vu.mutation.SetPower(s)
	return vu
}

// SetNillablePower sets the "power" field if the given value is not nil.
func (vu *VehicleUpdate) SetNillablePower(s *string) *VehicleUpdate {
	if s != nil {
		vu.SetPower(*s)
	}
	return vu
}

// ClearPower clears the value of the "power" field.
func (vu *VehicleUpdate) ClearPower() *VehicleUpdate {
	vu.mutation.ClearPower()
	return vu
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (vu *VehicleUpdate) SetLocationID(id int) *VehicleUpdate {
	vu.mutation.SetLocationID(id)
	return vu
}

// SetLocation sets the "location" edge to the Location entity.
func (vu *VehicleUpdate) SetLocation(l *Location) *VehicleUpdate {
	return vu.SetLocationID(l.ID)
}

// Mutation returns the VehicleMutation object of the builder.
func (vu *VehicleUpdate) Mutation() *VehicleMutation {
	return vu.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (vu *VehicleUpdate) ClearLocation() *VehicleUpdate {
	vu.mutation.ClearLocation()
	return vu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VehicleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	vu.defaults()
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			if vu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VehicleUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VehicleUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VehicleUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VehicleUpdate) defaults() {
	if _, ok := vu.mutation.UpdateTime(); !ok {
		v := vehicle.UpdateDefaultUpdateTime()
		vu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VehicleUpdate) check() error {
	if v, ok := vu.mutation.Hours(); ok {
		if err := vehicle.HoursValidator(v); err != nil {
			return &ValidationError{Name: "hours", err: fmt.Errorf("ent: validator failed for field \"hours\": %w", err)}
		}
	}
	if v, ok := vu.mutation.Year(); ok {
		if err := vehicle.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf("ent: validator failed for field \"year\": %w", err)}
		}
	}
	if v, ok := vu.mutation.Power(); ok {
		if err := vehicle.PowerValidator(v); err != nil {
			return &ValidationError{Name: "power", err: fmt.Errorf("ent: validator failed for field \"power\": %w", err)}
		}
	}
	if _, ok := vu.mutation.LocationID(); vu.mutation.LocationCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"location\"")
	}
	return nil
}

func (vu *VehicleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehicle.Table,
			Columns: vehicle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehicle.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldCreateTime,
		})
	}
	if value, ok := vu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldUpdateTime,
		})
	}
	if value, ok := vu.mutation.Make(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldMake,
		})
	}
	if value, ok := vu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldModel,
		})
	}
	if value, ok := vu.mutation.Hours(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldHours,
		})
	}
	if value, ok := vu.mutation.AddedHours(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldHours,
		})
	}
	if value, ok := vu.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldYear,
		})
	}
	if value, ok := vu.mutation.AddedYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldYear,
		})
	}
	if vu.mutation.YearCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: vehicle.FieldYear,
		})
	}
	if value, ok := vu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vehicle.FieldActive,
		})
	}
	if value, ok := vu.mutation.Power(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldPower,
		})
	}
	if vu.mutation.PowerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehicle.FieldPower,
		})
	}
	if vu.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicle.LocationTable,
			Columns: []string{vehicle.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicle.LocationTable,
			Columns: []string{vehicle.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VehicleUpdateOne is the builder for updating a single Vehicle entity.
type VehicleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VehicleMutation
}

// SetCreateTime sets the "create_time" field.
func (vuo *VehicleUpdateOne) SetCreateTime(t time.Time) *VehicleUpdateOne {
	vuo.mutation.SetCreateTime(t)
	return vuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vuo *VehicleUpdateOne) SetNillableCreateTime(t *time.Time) *VehicleUpdateOne {
	if t != nil {
		vuo.SetCreateTime(*t)
	}
	return vuo
}

// SetUpdateTime sets the "update_time" field.
func (vuo *VehicleUpdateOne) SetUpdateTime(t time.Time) *VehicleUpdateOne {
	vuo.mutation.SetUpdateTime(t)
	return vuo
}

// SetMake sets the "make" field.
func (vuo *VehicleUpdateOne) SetMake(s string) *VehicleUpdateOne {
	vuo.mutation.SetMake(s)
	return vuo
}

// SetModel sets the "model" field.
func (vuo *VehicleUpdateOne) SetModel(s string) *VehicleUpdateOne {
	vuo.mutation.SetModel(s)
	return vuo
}

// SetHours sets the "hours" field.
func (vuo *VehicleUpdateOne) SetHours(i int64) *VehicleUpdateOne {
	vuo.mutation.ResetHours()
	vuo.mutation.SetHours(i)
	return vuo
}

// SetNillableHours sets the "hours" field if the given value is not nil.
func (vuo *VehicleUpdateOne) SetNillableHours(i *int64) *VehicleUpdateOne {
	if i != nil {
		vuo.SetHours(*i)
	}
	return vuo
}

// AddHours adds i to the "hours" field.
func (vuo *VehicleUpdateOne) AddHours(i int64) *VehicleUpdateOne {
	vuo.mutation.AddHours(i)
	return vuo
}

// SetYear sets the "year" field.
func (vuo *VehicleUpdateOne) SetYear(i int64) *VehicleUpdateOne {
	vuo.mutation.ResetYear()
	vuo.mutation.SetYear(i)
	return vuo
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (vuo *VehicleUpdateOne) SetNillableYear(i *int64) *VehicleUpdateOne {
	if i != nil {
		vuo.SetYear(*i)
	}
	return vuo
}

// AddYear adds i to the "year" field.
func (vuo *VehicleUpdateOne) AddYear(i int64) *VehicleUpdateOne {
	vuo.mutation.AddYear(i)
	return vuo
}

// ClearYear clears the value of the "year" field.
func (vuo *VehicleUpdateOne) ClearYear() *VehicleUpdateOne {
	vuo.mutation.ClearYear()
	return vuo
}

// SetActive sets the "active" field.
func (vuo *VehicleUpdateOne) SetActive(b bool) *VehicleUpdateOne {
	vuo.mutation.SetActive(b)
	return vuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (vuo *VehicleUpdateOne) SetNillableActive(b *bool) *VehicleUpdateOne {
	if b != nil {
		vuo.SetActive(*b)
	}
	return vuo
}

// SetPower sets the "power" field.
func (vuo *VehicleUpdateOne) SetPower(s string) *VehicleUpdateOne {
	vuo.mutation.SetPower(s)
	return vuo
}

// SetNillablePower sets the "power" field if the given value is not nil.
func (vuo *VehicleUpdateOne) SetNillablePower(s *string) *VehicleUpdateOne {
	if s != nil {
		vuo.SetPower(*s)
	}
	return vuo
}

// ClearPower clears the value of the "power" field.
func (vuo *VehicleUpdateOne) ClearPower() *VehicleUpdateOne {
	vuo.mutation.ClearPower()
	return vuo
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (vuo *VehicleUpdateOne) SetLocationID(id int) *VehicleUpdateOne {
	vuo.mutation.SetLocationID(id)
	return vuo
}

// SetLocation sets the "location" edge to the Location entity.
func (vuo *VehicleUpdateOne) SetLocation(l *Location) *VehicleUpdateOne {
	return vuo.SetLocationID(l.ID)
}

// Mutation returns the VehicleMutation object of the builder.
func (vuo *VehicleUpdateOne) Mutation() *VehicleMutation {
	return vuo.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (vuo *VehicleUpdateOne) ClearLocation() *VehicleUpdateOne {
	vuo.mutation.ClearLocation()
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VehicleUpdateOne) Select(field string, fields ...string) *VehicleUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vehicle entity.
func (vuo *VehicleUpdateOne) Save(ctx context.Context) (*Vehicle, error) {
	var (
		err  error
		node *Vehicle
	)
	vuo.defaults()
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			if vuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VehicleUpdateOne) SaveX(ctx context.Context) *Vehicle {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VehicleUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VehicleUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VehicleUpdateOne) defaults() {
	if _, ok := vuo.mutation.UpdateTime(); !ok {
		v := vehicle.UpdateDefaultUpdateTime()
		vuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VehicleUpdateOne) check() error {
	if v, ok := vuo.mutation.Hours(); ok {
		if err := vehicle.HoursValidator(v); err != nil {
			return &ValidationError{Name: "hours", err: fmt.Errorf("ent: validator failed for field \"hours\": %w", err)}
		}
	}
	if v, ok := vuo.mutation.Year(); ok {
		if err := vehicle.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf("ent: validator failed for field \"year\": %w", err)}
		}
	}
	if v, ok := vuo.mutation.Power(); ok {
		if err := vehicle.PowerValidator(v); err != nil {
			return &ValidationError{Name: "power", err: fmt.Errorf("ent: validator failed for field \"power\": %w", err)}
		}
	}
	if _, ok := vuo.mutation.LocationID(); vuo.mutation.LocationCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"location\"")
	}
	return nil
}

func (vuo *VehicleUpdateOne) sqlSave(ctx context.Context) (_node *Vehicle, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehicle.Table,
			Columns: vehicle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vehicle.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Vehicle.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicle.FieldID)
		for _, f := range fields {
			if !vehicle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vehicle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldCreateTime,
		})
	}
	if value, ok := vuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldUpdateTime,
		})
	}
	if value, ok := vuo.mutation.Make(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldMake,
		})
	}
	if value, ok := vuo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldModel,
		})
	}
	if value, ok := vuo.mutation.Hours(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldHours,
		})
	}
	if value, ok := vuo.mutation.AddedHours(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldHours,
		})
	}
	if value, ok := vuo.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldYear,
		})
	}
	if value, ok := vuo.mutation.AddedYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vehicle.FieldYear,
		})
	}
	if vuo.mutation.YearCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: vehicle.FieldYear,
		})
	}
	if value, ok := vuo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vehicle.FieldActive,
		})
	}
	if value, ok := vuo.mutation.Power(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldPower,
		})
	}
	if vuo.mutation.PowerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vehicle.FieldPower,
		})
	}
	if vuo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicle.LocationTable,
			Columns: []string{vehicle.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehicle.LocationTable,
			Columns: []string{vehicle.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Vehicle{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
