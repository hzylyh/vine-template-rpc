// Code generated by ent, DO NOT EDIT.

package site

import (
	"vine-template-rpc/internal/emonitor/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldType, v))
}

// Dept applies equality check predicate on the "dept" field. It's identical to DeptEQ.
func Dept(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldDept, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldOwner, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldRemark, v))
}

// Lon applies equality check predicate on the "lon" field. It's identical to LonEQ.
func Lon(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldLon, v))
}

// Lat applies equality check predicate on the "lat" field. It's identical to LatEQ.
func Lat(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldLat, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldStatus, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldAddress, v))
}

// LastServiceTime applies equality check predicate on the "last_service_time" field. It's identical to LastServiceTimeEQ.
func LastServiceTime(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldLastServiceTime, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldCreateTime, v))
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldCreator, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldType, v))
}

// DeptEQ applies the EQ predicate on the "dept" field.
func DeptEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldDept, v))
}

// DeptNEQ applies the NEQ predicate on the "dept" field.
func DeptNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldDept, v))
}

// DeptIn applies the In predicate on the "dept" field.
func DeptIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldDept, vs...))
}

// DeptNotIn applies the NotIn predicate on the "dept" field.
func DeptNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldDept, vs...))
}

// DeptGT applies the GT predicate on the "dept" field.
func DeptGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldDept, v))
}

// DeptGTE applies the GTE predicate on the "dept" field.
func DeptGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldDept, v))
}

// DeptLT applies the LT predicate on the "dept" field.
func DeptLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldDept, v))
}

// DeptLTE applies the LTE predicate on the "dept" field.
func DeptLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldDept, v))
}

// DeptContains applies the Contains predicate on the "dept" field.
func DeptContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldDept, v))
}

// DeptHasPrefix applies the HasPrefix predicate on the "dept" field.
func DeptHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldDept, v))
}

// DeptHasSuffix applies the HasSuffix predicate on the "dept" field.
func DeptHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldDept, v))
}

// DeptIsNil applies the IsNil predicate on the "dept" field.
func DeptIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldDept))
}

// DeptNotNil applies the NotNil predicate on the "dept" field.
func DeptNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldDept))
}

// DeptEqualFold applies the EqualFold predicate on the "dept" field.
func DeptEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldDept, v))
}

// DeptContainsFold applies the ContainsFold predicate on the "dept" field.
func DeptContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldDept, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldOwner, v))
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldOwner, v))
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldOwner, v))
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldOwner, v))
}

// OwnerIsNil applies the IsNil predicate on the "owner" field.
func OwnerIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldOwner))
}

// OwnerNotNil applies the NotNil predicate on the "owner" field.
func OwnerNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldOwner))
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldOwner, v))
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldOwner, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldRemark, v))
}

// LonEQ applies the EQ predicate on the "lon" field.
func LonEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldLon, v))
}

// LonNEQ applies the NEQ predicate on the "lon" field.
func LonNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldLon, v))
}

// LonIn applies the In predicate on the "lon" field.
func LonIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldLon, vs...))
}

// LonNotIn applies the NotIn predicate on the "lon" field.
func LonNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldLon, vs...))
}

// LonGT applies the GT predicate on the "lon" field.
func LonGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldLon, v))
}

// LonGTE applies the GTE predicate on the "lon" field.
func LonGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldLon, v))
}

// LonLT applies the LT predicate on the "lon" field.
func LonLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldLon, v))
}

// LonLTE applies the LTE predicate on the "lon" field.
func LonLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldLon, v))
}

// LonContains applies the Contains predicate on the "lon" field.
func LonContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldLon, v))
}

// LonHasPrefix applies the HasPrefix predicate on the "lon" field.
func LonHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldLon, v))
}

// LonHasSuffix applies the HasSuffix predicate on the "lon" field.
func LonHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldLon, v))
}

// LonIsNil applies the IsNil predicate on the "lon" field.
func LonIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldLon))
}

// LonNotNil applies the NotNil predicate on the "lon" field.
func LonNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldLon))
}

// LonEqualFold applies the EqualFold predicate on the "lon" field.
func LonEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldLon, v))
}

// LonContainsFold applies the ContainsFold predicate on the "lon" field.
func LonContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldLon, v))
}

// LatEQ applies the EQ predicate on the "lat" field.
func LatEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldLat, v))
}

// LatNEQ applies the NEQ predicate on the "lat" field.
func LatNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldLat, v))
}

// LatIn applies the In predicate on the "lat" field.
func LatIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldLat, vs...))
}

// LatNotIn applies the NotIn predicate on the "lat" field.
func LatNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldLat, vs...))
}

// LatGT applies the GT predicate on the "lat" field.
func LatGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldLat, v))
}

// LatGTE applies the GTE predicate on the "lat" field.
func LatGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldLat, v))
}

// LatLT applies the LT predicate on the "lat" field.
func LatLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldLat, v))
}

// LatLTE applies the LTE predicate on the "lat" field.
func LatLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldLat, v))
}

// LatContains applies the Contains predicate on the "lat" field.
func LatContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldLat, v))
}

// LatHasPrefix applies the HasPrefix predicate on the "lat" field.
func LatHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldLat, v))
}

// LatHasSuffix applies the HasSuffix predicate on the "lat" field.
func LatHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldLat, v))
}

// LatIsNil applies the IsNil predicate on the "lat" field.
func LatIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldLat))
}

// LatNotNil applies the NotNil predicate on the "lat" field.
func LatNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldLat))
}

// LatEqualFold applies the EqualFold predicate on the "lat" field.
func LatEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldLat, v))
}

// LatContainsFold applies the ContainsFold predicate on the "lat" field.
func LatContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldLat, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldStatus, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldAddress, v))
}

// LastServiceTimeEQ applies the EQ predicate on the "last_service_time" field.
func LastServiceTimeEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldLastServiceTime, v))
}

// LastServiceTimeNEQ applies the NEQ predicate on the "last_service_time" field.
func LastServiceTimeNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldLastServiceTime, v))
}

// LastServiceTimeIn applies the In predicate on the "last_service_time" field.
func LastServiceTimeIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldLastServiceTime, vs...))
}

// LastServiceTimeNotIn applies the NotIn predicate on the "last_service_time" field.
func LastServiceTimeNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldLastServiceTime, vs...))
}

// LastServiceTimeGT applies the GT predicate on the "last_service_time" field.
func LastServiceTimeGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldLastServiceTime, v))
}

// LastServiceTimeGTE applies the GTE predicate on the "last_service_time" field.
func LastServiceTimeGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldLastServiceTime, v))
}

// LastServiceTimeLT applies the LT predicate on the "last_service_time" field.
func LastServiceTimeLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldLastServiceTime, v))
}

// LastServiceTimeLTE applies the LTE predicate on the "last_service_time" field.
func LastServiceTimeLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldLastServiceTime, v))
}

// LastServiceTimeContains applies the Contains predicate on the "last_service_time" field.
func LastServiceTimeContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldLastServiceTime, v))
}

// LastServiceTimeHasPrefix applies the HasPrefix predicate on the "last_service_time" field.
func LastServiceTimeHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldLastServiceTime, v))
}

// LastServiceTimeHasSuffix applies the HasSuffix predicate on the "last_service_time" field.
func LastServiceTimeHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldLastServiceTime, v))
}

// LastServiceTimeIsNil applies the IsNil predicate on the "last_service_time" field.
func LastServiceTimeIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldLastServiceTime))
}

// LastServiceTimeNotNil applies the NotNil predicate on the "last_service_time" field.
func LastServiceTimeNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldLastServiceTime))
}

// LastServiceTimeEqualFold applies the EqualFold predicate on the "last_service_time" field.
func LastServiceTimeEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldLastServiceTime, v))
}

// LastServiceTimeContainsFold applies the ContainsFold predicate on the "last_service_time" field.
func LastServiceTimeContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldLastServiceTime, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeContains applies the Contains predicate on the "create_time" field.
func CreateTimeContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldCreateTime, v))
}

// CreateTimeHasPrefix applies the HasPrefix predicate on the "create_time" field.
func CreateTimeHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldCreateTime, v))
}

// CreateTimeHasSuffix applies the HasSuffix predicate on the "create_time" field.
func CreateTimeHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldCreateTime))
}

// CreateTimeEqualFold applies the EqualFold predicate on the "create_time" field.
func CreateTimeEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldCreateTime, v))
}

// CreateTimeContainsFold applies the ContainsFold predicate on the "create_time" field.
func CreateTimeContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldCreateTime, v))
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v string) predicate.Site {
	return predicate.Site(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...string) predicate.Site {
	return predicate.Site(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v string) predicate.Site {
	return predicate.Site(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v string) predicate.Site {
	return predicate.Site(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v string) predicate.Site {
	return predicate.Site(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v string) predicate.Site {
	return predicate.Site(sql.FieldLTE(FieldCreator, v))
}

// CreatorContains applies the Contains predicate on the "creator" field.
func CreatorContains(v string) predicate.Site {
	return predicate.Site(sql.FieldContains(FieldCreator, v))
}

// CreatorHasPrefix applies the HasPrefix predicate on the "creator" field.
func CreatorHasPrefix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasPrefix(FieldCreator, v))
}

// CreatorHasSuffix applies the HasSuffix predicate on the "creator" field.
func CreatorHasSuffix(v string) predicate.Site {
	return predicate.Site(sql.FieldHasSuffix(FieldCreator, v))
}

// CreatorIsNil applies the IsNil predicate on the "creator" field.
func CreatorIsNil() predicate.Site {
	return predicate.Site(sql.FieldIsNull(FieldCreator))
}

// CreatorNotNil applies the NotNil predicate on the "creator" field.
func CreatorNotNil() predicate.Site {
	return predicate.Site(sql.FieldNotNull(FieldCreator))
}

// CreatorEqualFold applies the EqualFold predicate on the "creator" field.
func CreatorEqualFold(v string) predicate.Site {
	return predicate.Site(sql.FieldEqualFold(FieldCreator, v))
}

// CreatorContainsFold applies the ContainsFold predicate on the "creator" field.
func CreatorContainsFold(v string) predicate.Site {
	return predicate.Site(sql.FieldContainsFold(FieldCreator, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Site) predicate.Site {
	return predicate.Site(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Site) predicate.Site {
	return predicate.Site(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Site) predicate.Site {
	return predicate.Site(sql.NotPredicates(p))
}
