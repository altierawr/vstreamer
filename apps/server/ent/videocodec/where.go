// Code generated by ent, DO NOT EDIT.

package videocodec

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/altierawr/vstreamer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldName, v))
}

// Mime applies equality check predicate on the "mime" field. It's identical to MimeEQ.
func Mime(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldMime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldContainsFold(FieldName, v))
}

// MimeEQ applies the EQ predicate on the "mime" field.
func MimeEQ(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldMime, v))
}

// MimeNEQ applies the NEQ predicate on the "mime" field.
func MimeNEQ(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNEQ(FieldMime, v))
}

// MimeIn applies the In predicate on the "mime" field.
func MimeIn(vs ...string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldIn(FieldMime, vs...))
}

// MimeNotIn applies the NotIn predicate on the "mime" field.
func MimeNotIn(vs ...string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNotIn(FieldMime, vs...))
}

// MimeGT applies the GT predicate on the "mime" field.
func MimeGT(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldGT(FieldMime, v))
}

// MimeGTE applies the GTE predicate on the "mime" field.
func MimeGTE(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldGTE(FieldMime, v))
}

// MimeLT applies the LT predicate on the "mime" field.
func MimeLT(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldLT(FieldMime, v))
}

// MimeLTE applies the LTE predicate on the "mime" field.
func MimeLTE(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldLTE(FieldMime, v))
}

// MimeContains applies the Contains predicate on the "mime" field.
func MimeContains(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldContains(FieldMime, v))
}

// MimeHasPrefix applies the HasPrefix predicate on the "mime" field.
func MimeHasPrefix(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldHasPrefix(FieldMime, v))
}

// MimeHasSuffix applies the HasSuffix predicate on the "mime" field.
func MimeHasSuffix(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldHasSuffix(FieldMime, v))
}

// MimeEqualFold applies the EqualFold predicate on the "mime" field.
func MimeEqualFold(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEqualFold(FieldMime, v))
}

// MimeContainsFold applies the ContainsFold predicate on the "mime" field.
func MimeContainsFold(v string) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldContainsFold(FieldMime, v))
}

// DynamicRangeEQ applies the EQ predicate on the "dynamic_range" field.
func DynamicRangeEQ(v DynamicRange) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldEQ(FieldDynamicRange, v))
}

// DynamicRangeNEQ applies the NEQ predicate on the "dynamic_range" field.
func DynamicRangeNEQ(v DynamicRange) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNEQ(FieldDynamicRange, v))
}

// DynamicRangeIn applies the In predicate on the "dynamic_range" field.
func DynamicRangeIn(vs ...DynamicRange) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldIn(FieldDynamicRange, vs...))
}

// DynamicRangeNotIn applies the NotIn predicate on the "dynamic_range" field.
func DynamicRangeNotIn(vs ...DynamicRange) predicate.VideoCodec {
	return predicate.VideoCodec(sql.FieldNotIn(FieldDynamicRange, vs...))
}

// HasStreams applies the HasEdge predicate on the "streams" edge.
func HasStreams() predicate.VideoCodec {
	return predicate.VideoCodec(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StreamsTable, StreamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStreamsWith applies the HasEdge predicate on the "streams" edge with a given conditions (other predicates).
func HasStreamsWith(preds ...predicate.Stream) predicate.VideoCodec {
	return predicate.VideoCodec(func(s *sql.Selector) {
		step := newStreamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMedia applies the HasEdge predicate on the "media" edge.
func HasMedia() predicate.VideoCodec {
	return predicate.VideoCodec(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MediaTable, MediaColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMediaWith applies the HasEdge predicate on the "media" edge with a given conditions (other predicates).
func HasMediaWith(preds ...predicate.PlaySessionMedia) predicate.VideoCodec {
	return predicate.VideoCodec(func(s *sql.Selector) {
		step := newMediaStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VideoCodec) predicate.VideoCodec {
	return predicate.VideoCodec(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VideoCodec) predicate.VideoCodec {
	return predicate.VideoCodec(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VideoCodec) predicate.VideoCodec {
	return predicate.VideoCodec(sql.NotPredicates(p))
}
