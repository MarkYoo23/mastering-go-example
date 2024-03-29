// Code generated by ent, DO NOT EDIT.

package entry

import (
	"time"
	"white-page/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Entry {
	return predicate.Entry(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Entry {
	return predicate.Entry(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Entry {
	return predicate.Entry(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Entry {
	return predicate.Entry(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Entry {
	return predicate.Entry(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Entry {
	return predicate.Entry(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Entry {
	return predicate.Entry(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldName, v))
}

// Surname applies equality check predicate on the "surname" field. It's identical to SurnameEQ.
func Surname(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldSurname, v))
}

// Tel applies equality check predicate on the "tel" field. It's identical to TelEQ.
func Tel(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldTel, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Entry {
	return predicate.Entry(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Entry {
	return predicate.Entry(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Entry {
	return predicate.Entry(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Entry {
	return predicate.Entry(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Entry {
	return predicate.Entry(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Entry {
	return predicate.Entry(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Entry {
	return predicate.Entry(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Entry {
	return predicate.Entry(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Entry {
	return predicate.Entry(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Entry {
	return predicate.Entry(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Entry {
	return predicate.Entry(sql.FieldContainsFold(FieldName, v))
}

// SurnameEQ applies the EQ predicate on the "surname" field.
func SurnameEQ(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldSurname, v))
}

// SurnameNEQ applies the NEQ predicate on the "surname" field.
func SurnameNEQ(v string) predicate.Entry {
	return predicate.Entry(sql.FieldNEQ(FieldSurname, v))
}

// SurnameIn applies the In predicate on the "surname" field.
func SurnameIn(vs ...string) predicate.Entry {
	return predicate.Entry(sql.FieldIn(FieldSurname, vs...))
}

// SurnameNotIn applies the NotIn predicate on the "surname" field.
func SurnameNotIn(vs ...string) predicate.Entry {
	return predicate.Entry(sql.FieldNotIn(FieldSurname, vs...))
}

// SurnameGT applies the GT predicate on the "surname" field.
func SurnameGT(v string) predicate.Entry {
	return predicate.Entry(sql.FieldGT(FieldSurname, v))
}

// SurnameGTE applies the GTE predicate on the "surname" field.
func SurnameGTE(v string) predicate.Entry {
	return predicate.Entry(sql.FieldGTE(FieldSurname, v))
}

// SurnameLT applies the LT predicate on the "surname" field.
func SurnameLT(v string) predicate.Entry {
	return predicate.Entry(sql.FieldLT(FieldSurname, v))
}

// SurnameLTE applies the LTE predicate on the "surname" field.
func SurnameLTE(v string) predicate.Entry {
	return predicate.Entry(sql.FieldLTE(FieldSurname, v))
}

// SurnameContains applies the Contains predicate on the "surname" field.
func SurnameContains(v string) predicate.Entry {
	return predicate.Entry(sql.FieldContains(FieldSurname, v))
}

// SurnameHasPrefix applies the HasPrefix predicate on the "surname" field.
func SurnameHasPrefix(v string) predicate.Entry {
	return predicate.Entry(sql.FieldHasPrefix(FieldSurname, v))
}

// SurnameHasSuffix applies the HasSuffix predicate on the "surname" field.
func SurnameHasSuffix(v string) predicate.Entry {
	return predicate.Entry(sql.FieldHasSuffix(FieldSurname, v))
}

// SurnameEqualFold applies the EqualFold predicate on the "surname" field.
func SurnameEqualFold(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEqualFold(FieldSurname, v))
}

// SurnameContainsFold applies the ContainsFold predicate on the "surname" field.
func SurnameContainsFold(v string) predicate.Entry {
	return predicate.Entry(sql.FieldContainsFold(FieldSurname, v))
}

// TelEQ applies the EQ predicate on the "tel" field.
func TelEQ(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldTel, v))
}

// TelNEQ applies the NEQ predicate on the "tel" field.
func TelNEQ(v string) predicate.Entry {
	return predicate.Entry(sql.FieldNEQ(FieldTel, v))
}

// TelIn applies the In predicate on the "tel" field.
func TelIn(vs ...string) predicate.Entry {
	return predicate.Entry(sql.FieldIn(FieldTel, vs...))
}

// TelNotIn applies the NotIn predicate on the "tel" field.
func TelNotIn(vs ...string) predicate.Entry {
	return predicate.Entry(sql.FieldNotIn(FieldTel, vs...))
}

// TelGT applies the GT predicate on the "tel" field.
func TelGT(v string) predicate.Entry {
	return predicate.Entry(sql.FieldGT(FieldTel, v))
}

// TelGTE applies the GTE predicate on the "tel" field.
func TelGTE(v string) predicate.Entry {
	return predicate.Entry(sql.FieldGTE(FieldTel, v))
}

// TelLT applies the LT predicate on the "tel" field.
func TelLT(v string) predicate.Entry {
	return predicate.Entry(sql.FieldLT(FieldTel, v))
}

// TelLTE applies the LTE predicate on the "tel" field.
func TelLTE(v string) predicate.Entry {
	return predicate.Entry(sql.FieldLTE(FieldTel, v))
}

// TelContains applies the Contains predicate on the "tel" field.
func TelContains(v string) predicate.Entry {
	return predicate.Entry(sql.FieldContains(FieldTel, v))
}

// TelHasPrefix applies the HasPrefix predicate on the "tel" field.
func TelHasPrefix(v string) predicate.Entry {
	return predicate.Entry(sql.FieldHasPrefix(FieldTel, v))
}

// TelHasSuffix applies the HasSuffix predicate on the "tel" field.
func TelHasSuffix(v string) predicate.Entry {
	return predicate.Entry(sql.FieldHasSuffix(FieldTel, v))
}

// TelEqualFold applies the EqualFold predicate on the "tel" field.
func TelEqualFold(v string) predicate.Entry {
	return predicate.Entry(sql.FieldEqualFold(FieldTel, v))
}

// TelContainsFold applies the ContainsFold predicate on the "tel" field.
func TelContainsFold(v string) predicate.Entry {
	return predicate.Entry(sql.FieldContainsFold(FieldTel, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Entry {
	return predicate.Entry(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Entry) predicate.Entry {
	return predicate.Entry(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Entry) predicate.Entry {
	return predicate.Entry(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Entry) predicate.Entry {
	return predicate.Entry(sql.NotPredicates(p))
}
