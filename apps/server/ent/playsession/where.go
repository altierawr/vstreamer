// Code generated by ent, DO NOT EDIT.

package playsession

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/altierawr/vstreamer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PlaySession {
	return predicate.PlaySession(sql.FieldLTE(FieldID, id))
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.PlaySession {
	return predicate.PlaySession(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.PlaySession {
	return predicate.PlaySession(func(s *sql.Selector) {
		step := newVideoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlaySession) predicate.PlaySession {
	return predicate.PlaySession(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlaySession) predicate.PlaySession {
	return predicate.PlaySession(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlaySession) predicate.PlaySession {
	return predicate.PlaySession(sql.NotPredicates(p))
}