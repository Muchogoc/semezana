// Code generated by ent, DO NOT EDIT.

package channel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Muchogoc/semezana/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateAt applies equality check predicate on the "state_at" field. It's identical to StateAtEQ.
func StateAt(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStateAt), v))
	})
}

// Sequence applies equality check predicate on the "sequence" field. It's identical to SequenceEQ.
func Sequence(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequence), v))
	})
}

// Touched applies equality check predicate on the "touched" field. It's identical to TouchedEQ.
func Touched(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTouched), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// StateAtEQ applies the EQ predicate on the "state_at" field.
func StateAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStateAt), v))
	})
}

// StateAtNEQ applies the NEQ predicate on the "state_at" field.
func StateAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStateAt), v))
	})
}

// StateAtIn applies the In predicate on the "state_at" field.
func StateAtIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStateAt), v...))
	})
}

// StateAtNotIn applies the NotIn predicate on the "state_at" field.
func StateAtNotIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStateAt), v...))
	})
}

// StateAtGT applies the GT predicate on the "state_at" field.
func StateAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStateAt), v))
	})
}

// StateAtGTE applies the GTE predicate on the "state_at" field.
func StateAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStateAt), v))
	})
}

// StateAtLT applies the LT predicate on the "state_at" field.
func StateAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStateAt), v))
	})
}

// StateAtLTE applies the LTE predicate on the "state_at" field.
func StateAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStateAt), v))
	})
}

// SequenceEQ applies the EQ predicate on the "sequence" field.
func SequenceEQ(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequence), v))
	})
}

// SequenceNEQ applies the NEQ predicate on the "sequence" field.
func SequenceNEQ(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSequence), v))
	})
}

// SequenceIn applies the In predicate on the "sequence" field.
func SequenceIn(vs ...int) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSequence), v...))
	})
}

// SequenceNotIn applies the NotIn predicate on the "sequence" field.
func SequenceNotIn(vs ...int) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSequence), v...))
	})
}

// SequenceGT applies the GT predicate on the "sequence" field.
func SequenceGT(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSequence), v))
	})
}

// SequenceGTE applies the GTE predicate on the "sequence" field.
func SequenceGTE(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSequence), v))
	})
}

// SequenceLT applies the LT predicate on the "sequence" field.
func SequenceLT(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSequence), v))
	})
}

// SequenceLTE applies the LTE predicate on the "sequence" field.
func SequenceLTE(v int) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSequence), v))
	})
}

// TouchedEQ applies the EQ predicate on the "touched" field.
func TouchedEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTouched), v))
	})
}

// TouchedNEQ applies the NEQ predicate on the "touched" field.
func TouchedNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTouched), v))
	})
}

// TouchedIn applies the In predicate on the "touched" field.
func TouchedIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTouched), v...))
	})
}

// TouchedNotIn applies the NotIn predicate on the "touched" field.
func TouchedNotIn(vs ...time.Time) predicate.Channel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTouched), v...))
	})
}

// TouchedGT applies the GT predicate on the "touched" field.
func TouchedGT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTouched), v))
	})
}

// TouchedGTE applies the GTE predicate on the "touched" field.
func TouchedGTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTouched), v))
	})
}

// TouchedLT applies the LT predicate on the "touched" field.
func TouchedLT(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTouched), v))
	})
}

// TouchedLTE applies the LTE predicate on the "touched" field.
func TouchedLTE(v time.Time) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTouched), v))
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.User) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscriptions applies the HasEdge predicate on the "subscriptions" edge.
func HasSubscriptions() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubscriptionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SubscriptionsTable, SubscriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscriptionsWith applies the HasEdge predicate on the "subscriptions" edge with a given conditions (other predicates).
func HasSubscriptionsWith(preds ...predicate.Subscription) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubscriptionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SubscriptionsTable, SubscriptionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		p(s.Not())
	})
}