package tiv

import "reflect"

type ValidationError struct {

	// Tag returns the validation tag that failed. if the
	// validation was an alias, this will return the
	// alias name and not the underlying tag that failed.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "iscolor"
	Tag string

	// ActualTag returns the validation tag that failed, even if an
	// alias the actual tag within the alias will be returned.
	// If an 'or' validation fails the entire or will be returned.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "hexcolor|rgb|rgba|hsl|hsla"
	ActualTag string

	// Namespace returns the namespace for the field error, with the tag
	// name taking precedence over the field's actual name.
	//
	// eg. JSON name "User.fname"
	//
	// See StructNamespace() for a version that returns actual names.
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract it's name
	Namespace string

	// StructNamespace returns the namespace for the field error, with the field's
	// actual name.
	//
	// eq. "User.FirstName" see Namespace for comparison
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract its name
	StructNamespace string

	// Field returns the fields name with the tag name taking precedence over the
	// field's actual name.
	//
	// eq. JSON name "fname"
	// see StructField for comparison
	Field string

	// StructField returns the field's actual name from the struct, when able to determine.
	//
	// eq.  "FirstName"
	// see Field for comparison
	StructField string

	// Value returns the actual field's value in case needed for creating the error
	// message
	Value interface{}

	// Param returns the param value, in string form for comparison; this will also
	// help with generating an error message
	Param string

	// Kind returns the Field's reflect Kind
	//
	// eg. time.Time's kind is a struct
	Kind reflect.Kind

	// Type returns the Field's reflect Type
	//
	// eg. time.Time's type is time.Time
	Type reflect.Type

	// Error returns the FieldError's message
	Error string

	RealError error
}
