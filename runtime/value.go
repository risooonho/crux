package runtime

import "math/big"

type Value interface{}

type (
	Char  struct{ Value rune }
	Int   struct{ Value big.Int }
	Float struct{ Value float64 }

	Struct struct {
		Index  int32
		Values []Value
	}

	Thunk struct {
		Result Value
		Code   *Code
		Data   []Value
	}
)

type Code struct {
	Kind  CodeKind
	X     int32
	Table []Code
	Value Value
}

type CodeKind int32

const (
	CodeValue CodeKind = iota
	CodeOperator
	CodeMake
	CodeVar
	CodeGlobal
	CodeAbst
	CodeAppl
	CodeSwitch
)
