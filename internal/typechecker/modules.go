package typechecker

import "strings"

var modules = map[string]*Env{}

func RegisterModule(path []string, env *Env) {
	modules[strings.Join(path, "/")] = env
}

func getModule(path []string) (*Env, bool) {
	env, ok := modules[strings.Join(path, "/")]
	return env, ok
}

// Module testing

func init() {
	ioEnv := NewEnv(nil)
	ioEnv.Set("print", &Type{
		TKind:  TyFunc,
		Params: []*Type{{TKind: TyString}},
		Ret:    &Type{TKind: TyNil},
	})
	ioEnv.Set("println", &Type{
		TKind:  TyFunc,
		Params: []*Type{{TKind: TyString}},
		Ret:    &Type{TKind: TyNil},
	})
	RegisterModule([]string{"flint", "io"}, ioEnv)

	strEnv := NewEnv(nil)
	strEnv.Set("to_string", &Type{
		TKind:  TyFunc,
		Params: []*Type{{TKind: TyInt}},
		Ret:    &Type{TKind: TyString},
	})
	RegisterModule([]string{"flint", "string"}, strEnv)
}
