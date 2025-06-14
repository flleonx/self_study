package object

import (
	"bytes"
	"fmt"
	"strings"

	"monkey/ast"
)

const (
	INTERGER_OBJ     = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
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

	Function struct {
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Env        *Environment
	}

	String struct {
		Value string
	}

	BuiltinFunction func(args ...Object) Object

	Builtin struct {
		Fn BuiltinFunction
	}

	Array struct {
		Elements []Object
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

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return fmt.Sprintf("ERROR: %s", e.Message) }

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
