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

// NewCSL generates a new CSL interpretor.
func NewCSL(registerDefaults bool) *CSL {
	csl := new(CSL)
	csl.builtins = make(map[string]operation)
	if registerDefaults {
		csl.RegisterDefaults()
	}
	return csl
}

func (csl *CSL) ResolveDepsAndCall(cslpm *CSLPackageManager, script Script, args ...interface{}) (result interface{}, err error, trace Trace) {
	if err = csl.PreprocessScript(script.Code, script.ImportPath, cslpm); err != nil {
		return
	}
	var sexprs []*SExpr
	sexprs, err = csl.Parse(script.Code)
	if err != nil {
		return
	}
	result, trace = csl.Invoke(sexprs, args...)
	return
}
