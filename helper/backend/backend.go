package backend

import (
	"github.com/hashicorp/vault/vault"
)

// Backend is an implementation of vault.LogicalBackend that allows
// the implementer to code a backend using a much more programmer-friendly
// framework that handles a lot of the routing and validation for you.
//
// This is recommended over implementing vault.LogicalBackend directly.
type Backend struct {
	Paths []*Path
}

// Path is a single path that the backend responds to.
type Path struct {
	// Pattern is the pattern of the URL that matches this path.
	//
	// This should be a valid regular expression. Named captures will be
	// exposed as fields that should map to a schema in Fields. If a named
	// capture is not a field in the Fields map, then it will be ignored.
	Pattern string

	// Fields is the mapping of data fields to a schema describing that
	// field. Named captures in the Pattern also map to fields. If a named
	// capture name matches a PUT body name, the named capture takes
	// priority.
	//
	// Note that only named capture fields are available in every operation,
	// whereas all fields are avaiable in the Write operation.
	Fields map[string]*FieldSchema

	// Root if not blank, denotes that this path requires root
	// privileges and the path pattern that is the root path. This can't
	// be a regular expression and must be an exact path. It may have a
	// trailing '*' to denote that it is a prefix, and not an exact match.
	Root string

	// Callback is what is called when this path is requested with
	// a valid set of data.
	Callback func(*vault.Request, *FieldData) (*vault.Response, error)
}

// FieldSchema is a basic schema to describe the format of a path field.
type FieldSchema struct {
	Type FieldType
}

func (t FieldType) Zero() interface{} {
	switch t {
	case TypeString:
		return ""
	case TypeInt:
		return 0
	case TypeBool:
		return false
	default:
		panic("unknown type: " + t.String())
	}
}
