package typechecker

// Testing for module type checking

import "strings"

var modules = map[string]*Env{}

func registerModule(path []string, env *Env) {
	modules[strings.Join(path, "/")] = env
}

func getModule(path []string) (*Env, bool) {
	env, ok := modules[strings.Join(path, "/")]
	return env, ok
}

func init() {
	ioEnv := NewEnv(nil)
	ioEnv.Set("print", &Type{
		kind:   TyFunc,
		Params: []*Type{{kind: TyString}},
		Ret:    &Type{kind: TyNil},
	})
	ioEnv.Set("println", &Type{
		kind:   TyFunc,
		Params: []*Type{{kind: TyString}},
		Ret:    &Type{kind: TyNil},
	})
	registerModule([]string{"flint", "io"}, ioEnv)

	strEnv := NewEnv(nil)
	strEnv.Set("to_string", &Type{
		kind:   TyFunc,
		Params: []*Type{{kind: TyInt}},
		Ret:    &Type{kind: TyString},
	})
	registerModule([]string{"flint", "string"}, strEnv)
}
