package object

import "fmt"

const (
	INTERGER_OBJ     = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

type (
	ObjectType string

	Object interface {
		Type() ObjectType
		Inspect() string
	}

	Integer struct {
		Value int64
	}

	Boolean struct {
		Value bool
	}

	Null struct{}

	ReturnValue struct {
		Value Object
	}

	Error struct {
		Message string
	}
)

func (i *Integer) Type() ObjectType { return INTERGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return fmt.Sprintf("ERROR: %s", e.Message) }
