package typechecker

import "flint/internal/lexer"

type BinOpSig struct {
	Left  Type
	Right Type
	Out   Type
}

type UnaryOpSig struct {
	Arg Type
	Out Type
}

var binOps = map[lexer.TokenKind][]BinOpSig{
	lexer.Plus:    {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyInt}}},
	lexer.Minus:   {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyInt}}},
	lexer.Star:    {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyInt}}},
	lexer.Slash:   {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyInt}}},
	lexer.Percent: {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyInt}}},

	lexer.PlusDot:  {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyFloat}}},
	lexer.MinusDot: {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyFloat}}},
	lexer.StarDot:  {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyFloat}}},
	lexer.SlashDot: {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyFloat}}},

	lexer.Less:         {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyBool}}},
	lexer.LessEqual:    {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyBool}}},
	lexer.Greater:      {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyBool}}},
	lexer.GreaterEqual: {{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyBool}}},

	lexer.LessDot:         {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyBool}}},
	lexer.LessEqualDot:    {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyBool}}},
	lexer.GreaterDot:      {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyBool}}},
	lexer.GreaterEqualDot: {{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyBool}}},

	lexer.AmperAmper: {{Type{TKind: TyBool}, Type{TKind: TyBool}, Type{TKind: TyBool}}},
	lexer.VbarVbar:   {{Type{TKind: TyBool}, Type{TKind: TyBool}, Type{TKind: TyBool}}},

	lexer.LtGt: {{Type{TKind: TyString}, Type{TKind: TyString}, Type{TKind: TyString}}},

	lexer.EqualEqual: {
		{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyBool}},
		{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyBool}},
		{Type{TKind: TyBool}, Type{TKind: TyBool}, Type{TKind: TyBool}},
		{Type{TKind: TyString}, Type{TKind: TyString}, Type{TKind: TyBool}},
		{Type{TKind: TyByte}, Type{TKind: TyByte}, Type{TKind: TyBool}},
	},
	lexer.NotEqual: {
		{Type{TKind: TyInt}, Type{TKind: TyInt}, Type{TKind: TyBool}},
		{Type{TKind: TyFloat}, Type{TKind: TyFloat}, Type{TKind: TyBool}},
		{Type{TKind: TyBool}, Type{TKind: TyBool}, Type{TKind: TyBool}},
		{Type{TKind: TyString}, Type{TKind: TyString}, Type{TKind: TyBool}},
		{Type{TKind: TyByte}, Type{TKind: TyByte}, Type{TKind: TyBool}},
	},
}

var unaryOps = map[lexer.TokenKind]UnaryOpSig{
	lexer.Minus:    {Type{TKind: TyInt}, Type{TKind: TyInt}},
	lexer.MinusDot: {Type{TKind: TyFloat}, Type{TKind: TyFloat}},
	lexer.Bang:     {Type{TKind: TyBool}, Type{TKind: TyBool}},
}
