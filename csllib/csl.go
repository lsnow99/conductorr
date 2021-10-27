package csllib

type CSL struct {
	builtins map[string]operation
}

type CSLPackageManager struct {
	scriptFetcher         ScriptFetcher
	allowInsecureRequests bool
}

type Script struct {
	Code       string
	ImportPath string
}

func NewCSL(registerDefaults bool) *CSL {
	csl := new(CSL)
	csl.builtins = make(map[string]operation)
	if registerDefaults {
		csl.RegisterDefaults()
	}
	return csl
}

func ResolveDepsAndCall(cslpm *CSLPackageManager, script Script, args ...interface{}) (interface{}, error) {
	csl := NewCSL(true)
	if err := csl.PreprocessScript(script.Code, script.ImportPath, cslpm); err != nil {
		return nil, err
	}
	sexprs, err := csl.Parse(script.Code)
	if err != nil {
		return nil, err
	}
	env := make(map[string]interface{})
	var argsList List
	for _, arg := range args {
		argsList = append(argsList, arg)
	}
	env["args"] = argsList
	result, trace := csl.Eval(sexprs, env)
	if trace.Err != nil {
		return nil, trace.Err
	}
	return result, nil
}
