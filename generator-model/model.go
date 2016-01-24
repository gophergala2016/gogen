package model

import "errors"

// Predefined errors
var (
	ErrFieldNotFound = errors.New("Specified field could not be found")
)

// Model defines entity that can be used to
// populate information needed by the generators.
//
// This includes fields, enums and helper methods.
// It may also restrict usage of these fields based on
// storage or access technology used
type Model struct {
	// Name of the model by which it should be distinguished
	Name string
	// Fields is array of the fields that should the
	// model have. Field may restrict itself to any
	Fields []Field
	// Restrictions array defines a number of restrictions
	// that should be used while generating the model
	Restrictions []Restriction
}

// AddField extends fields definition by the provided fields.
// Returns self.
func (m *Model) AddField(field Field, fields ...Field) *Model {
	m.Fields = append(m.Fields, field)
	m.Fields = append(m.Fields, fields...)
	return m
}

// AddRestriction extends restrictions list by the provided
// restrictions. Returns self.
func (m *Model) AddRestriction(r Restriction, rs ...Restriction) *Model {
	m.Restrictions = append(m.Restrictions, r)
	m.Restrictions = append(m.Restrictions, rs...)
	return m
}

// GetField will search the model for the field with
// the given name. If no field was found, function
// will rise an error.
func (m *Model) GetField(name string) (Field, error) {
	for _, f := range m.Fields {
		if f.Name == name {
			return f, nil
		}
	}

	return Field{}, ErrFieldNotFound
}

// Extend will extend passed model by the current model,
// adding fields and restrictions that are present in the
// current model
//
// Returns new model.
//
// Warning: this method will not extend the current model,
// but create new model and returns it. To extend current model,
// use Add methods
func (m *Model) Extend(by *Model) *Model {
	// copy fields
	for _, f := range m.Fields {
		by.AddField(f)
	}
	// copy restrictions
	for _, r := range m.Restrictions {
		by.AddRestriction(r)
	}

	return by
}

// Field defines model field. Every field may define
// set of data types and restrictions based on the
// storage that will be used
type Field struct {
	// Name of the field in the object notation
	Name string
	// Type must be from the range of the types provided by
	// the base type system or generator-specific
	Type FieldType
	// defines if the field is only helper field (non-exported).
	// Defaults to false
	Helper bool
}

// Restriction defines map of restrictions that may be
// used by the generators to give the set of rules
type Restriction struct {
}

// Relation is the enum variable that will
// have the value of OneToOne to ManyToMany from
// the constants below
type Relation int

// Constants providing indexes for the relations
const (
	OneToOne = iota
	OneToMany
	ManyToOne
	ManyToMany
)

// FieldType is interface that must be implemented
// for any generator type to be accepted.
type FieldType interface {
	// Name method returns the name of the type in
	// the golang.
	// example:
	//		return "time.Time"
	//
	// This name will be then used in the generator
	// to have the right name of the variable type
	Name() string
}

type (
	// StringType is type for the string
	StringType struct{}
	// Int8Type is type for the int8
	Int8Type struct{}
	// Int16Type is type for the int16
	Int16Type struct{}
	// Int32Type is type for the int32
	Int32Type struct{}
	// Float32Type is type for the float32
	Float32Type struct{}
	// Float64Type is type for the float64
	Float64Type struct{}
	// TimeType is type for the time.Time
	TimeType struct{}
	// BoolType is type for the bool type
	BoolType struct{}

	// Link encapsulates relationship made by the
	// field to the model
	Link struct {
		// Model we are linking to
		Model *Model
		// Relationship to the model
		Relation Relation
	}
)

// Name returns the true name of the type
func (t *StringType) Name() string {
	return "string"
}

// Name returns the true name of the type
func (t *Int8Type) Name() string {
	return "int8"
}

// Name returns the true name of the type
func (t *Int16Type) Name() string {
	return "int16"
}

// Name returns the true name of the type
func (t *Int32Type) Name() string {
	return "int32"
}

// Name returns the true name of the type
func (t *Float32Type) Name() string {
	return "float32"
}

// Name returns the true name of the type
func (t *Float64Type) Name() string {
	return "float64"
}

// Name returns the true name of the type
func (t *TimeType) Name() string {
	return "time.Time"
}

// Name returns the true name of the type
func (t *BoolType) Name() string {
	return "bool"
}

// Basic field type variables
var (
	String  = &StringType{}
	Int8    = &Int8Type{}
	Int16   = &Int16Type{}
	Int32   = &Int32Type{}
	Float32 = &Float32Type{}
	Float64 = &Float64Type{}
	Time    = &TimeType{}
	Bool    = &BoolType{}
)
