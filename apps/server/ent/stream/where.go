// Code generated by ent, DO NOT EDIT.

package stream

import (
	"entgo.io/ent/dialect/sql"
	"github.com/altierawr/vstreamer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldID, id))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldWidth, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldHeight, v))
}

// Container applies equality check predicate on the "container" field. It's identical to ContainerEQ.
func Container(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldContainer, v))
}

// VideoCodec applies equality check predicate on the "video_codec" field. It's identical to VideoCodecEQ.
func VideoCodec(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldVideoCodec, v))
}

// AudioCodec applies equality check predicate on the "audio_codec" field. It's identical to AudioCodecEQ.
func AudioCodec(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldAudioCodec, v))
}

// SegmentDuration applies equality check predicate on the "segment_duration" field. It's identical to SegmentDurationEQ.
func SegmentDuration(v int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldSegmentDuration, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldWidth, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldHeight, v))
}

// ContainerEQ applies the EQ predicate on the "container" field.
func ContainerEQ(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldContainer, v))
}

// ContainerNEQ applies the NEQ predicate on the "container" field.
func ContainerNEQ(v string) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldContainer, v))
}

// ContainerIn applies the In predicate on the "container" field.
func ContainerIn(vs ...string) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldContainer, vs...))
}

// ContainerNotIn applies the NotIn predicate on the "container" field.
func ContainerNotIn(vs ...string) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldContainer, vs...))
}

// ContainerGT applies the GT predicate on the "container" field.
func ContainerGT(v string) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldContainer, v))
}

// ContainerGTE applies the GTE predicate on the "container" field.
func ContainerGTE(v string) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldContainer, v))
}

// ContainerLT applies the LT predicate on the "container" field.
func ContainerLT(v string) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldContainer, v))
}

// ContainerLTE applies the LTE predicate on the "container" field.
func ContainerLTE(v string) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldContainer, v))
}

// ContainerContains applies the Contains predicate on the "container" field.
func ContainerContains(v string) predicate.Stream {
	return predicate.Stream(sql.FieldContains(FieldContainer, v))
}

// ContainerHasPrefix applies the HasPrefix predicate on the "container" field.
func ContainerHasPrefix(v string) predicate.Stream {
	return predicate.Stream(sql.FieldHasPrefix(FieldContainer, v))
}

// ContainerHasSuffix applies the HasSuffix predicate on the "container" field.
func ContainerHasSuffix(v string) predicate.Stream {
	return predicate.Stream(sql.FieldHasSuffix(FieldContainer, v))
}

// ContainerEqualFold applies the EqualFold predicate on the "container" field.
func ContainerEqualFold(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEqualFold(FieldContainer, v))
}

// ContainerContainsFold applies the ContainsFold predicate on the "container" field.
func ContainerContainsFold(v string) predicate.Stream {
	return predicate.Stream(sql.FieldContainsFold(FieldContainer, v))
}

// VideoCodecEQ applies the EQ predicate on the "video_codec" field.
func VideoCodecEQ(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldVideoCodec, v))
}

// VideoCodecNEQ applies the NEQ predicate on the "video_codec" field.
func VideoCodecNEQ(v string) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldVideoCodec, v))
}

// VideoCodecIn applies the In predicate on the "video_codec" field.
func VideoCodecIn(vs ...string) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldVideoCodec, vs...))
}

// VideoCodecNotIn applies the NotIn predicate on the "video_codec" field.
func VideoCodecNotIn(vs ...string) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldVideoCodec, vs...))
}

// VideoCodecGT applies the GT predicate on the "video_codec" field.
func VideoCodecGT(v string) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldVideoCodec, v))
}

// VideoCodecGTE applies the GTE predicate on the "video_codec" field.
func VideoCodecGTE(v string) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldVideoCodec, v))
}

// VideoCodecLT applies the LT predicate on the "video_codec" field.
func VideoCodecLT(v string) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldVideoCodec, v))
}

// VideoCodecLTE applies the LTE predicate on the "video_codec" field.
func VideoCodecLTE(v string) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldVideoCodec, v))
}

// VideoCodecContains applies the Contains predicate on the "video_codec" field.
func VideoCodecContains(v string) predicate.Stream {
	return predicate.Stream(sql.FieldContains(FieldVideoCodec, v))
}

// VideoCodecHasPrefix applies the HasPrefix predicate on the "video_codec" field.
func VideoCodecHasPrefix(v string) predicate.Stream {
	return predicate.Stream(sql.FieldHasPrefix(FieldVideoCodec, v))
}

// VideoCodecHasSuffix applies the HasSuffix predicate on the "video_codec" field.
func VideoCodecHasSuffix(v string) predicate.Stream {
	return predicate.Stream(sql.FieldHasSuffix(FieldVideoCodec, v))
}

// VideoCodecEqualFold applies the EqualFold predicate on the "video_codec" field.
func VideoCodecEqualFold(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEqualFold(FieldVideoCodec, v))
}

// VideoCodecContainsFold applies the ContainsFold predicate on the "video_codec" field.
func VideoCodecContainsFold(v string) predicate.Stream {
	return predicate.Stream(sql.FieldContainsFold(FieldVideoCodec, v))
}

// AudioCodecEQ applies the EQ predicate on the "audio_codec" field.
func AudioCodecEQ(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldAudioCodec, v))
}

// AudioCodecNEQ applies the NEQ predicate on the "audio_codec" field.
func AudioCodecNEQ(v string) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldAudioCodec, v))
}

// AudioCodecIn applies the In predicate on the "audio_codec" field.
func AudioCodecIn(vs ...string) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldAudioCodec, vs...))
}

// AudioCodecNotIn applies the NotIn predicate on the "audio_codec" field.
func AudioCodecNotIn(vs ...string) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldAudioCodec, vs...))
}

// AudioCodecGT applies the GT predicate on the "audio_codec" field.
func AudioCodecGT(v string) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldAudioCodec, v))
}

// AudioCodecGTE applies the GTE predicate on the "audio_codec" field.
func AudioCodecGTE(v string) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldAudioCodec, v))
}

// AudioCodecLT applies the LT predicate on the "audio_codec" field.
func AudioCodecLT(v string) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldAudioCodec, v))
}

// AudioCodecLTE applies the LTE predicate on the "audio_codec" field.
func AudioCodecLTE(v string) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldAudioCodec, v))
}

// AudioCodecContains applies the Contains predicate on the "audio_codec" field.
func AudioCodecContains(v string) predicate.Stream {
	return predicate.Stream(sql.FieldContains(FieldAudioCodec, v))
}

// AudioCodecHasPrefix applies the HasPrefix predicate on the "audio_codec" field.
func AudioCodecHasPrefix(v string) predicate.Stream {
	return predicate.Stream(sql.FieldHasPrefix(FieldAudioCodec, v))
}

// AudioCodecHasSuffix applies the HasSuffix predicate on the "audio_codec" field.
func AudioCodecHasSuffix(v string) predicate.Stream {
	return predicate.Stream(sql.FieldHasSuffix(FieldAudioCodec, v))
}

// AudioCodecEqualFold applies the EqualFold predicate on the "audio_codec" field.
func AudioCodecEqualFold(v string) predicate.Stream {
	return predicate.Stream(sql.FieldEqualFold(FieldAudioCodec, v))
}

// AudioCodecContainsFold applies the ContainsFold predicate on the "audio_codec" field.
func AudioCodecContainsFold(v string) predicate.Stream {
	return predicate.Stream(sql.FieldContainsFold(FieldAudioCodec, v))
}

// SegmentDurationEQ applies the EQ predicate on the "segment_duration" field.
func SegmentDurationEQ(v int) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldSegmentDuration, v))
}

// SegmentDurationNEQ applies the NEQ predicate on the "segment_duration" field.
func SegmentDurationNEQ(v int) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldSegmentDuration, v))
}

// SegmentDurationIn applies the In predicate on the "segment_duration" field.
func SegmentDurationIn(vs ...int) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldSegmentDuration, vs...))
}

// SegmentDurationNotIn applies the NotIn predicate on the "segment_duration" field.
func SegmentDurationNotIn(vs ...int) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldSegmentDuration, vs...))
}

// SegmentDurationGT applies the GT predicate on the "segment_duration" field.
func SegmentDurationGT(v int) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldSegmentDuration, v))
}

// SegmentDurationGTE applies the GTE predicate on the "segment_duration" field.
func SegmentDurationGTE(v int) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldSegmentDuration, v))
}

// SegmentDurationLT applies the LT predicate on the "segment_duration" field.
func SegmentDurationLT(v int) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldSegmentDuration, v))
}

// SegmentDurationLTE applies the LTE predicate on the "segment_duration" field.
func SegmentDurationLTE(v int) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldSegmentDuration, v))
}

// QualityEQ applies the EQ predicate on the "quality" field.
func QualityEQ(v Quality) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldQuality, v))
}

// QualityNEQ applies the NEQ predicate on the "quality" field.
func QualityNEQ(v Quality) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldQuality, v))
}

// QualityIn applies the In predicate on the "quality" field.
func QualityIn(vs ...Quality) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldQuality, vs...))
}

// QualityNotIn applies the NotIn predicate on the "quality" field.
func QualityNotIn(vs ...Quality) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldQuality, vs...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldType, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stream) predicate.Stream {
	return predicate.Stream(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stream) predicate.Stream {
	return predicate.Stream(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Stream) predicate.Stream {
	return predicate.Stream(sql.NotPredicates(p))
}