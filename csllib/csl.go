package csllib

type CSL struct {
	builtins map[string]operation
}

type CSLPackageManager struct {
	scriptFetcher ScriptFetcher
	allowInsecureRequests bool
}

func NewCSL(registerDefaults bool) *CSL {
	csl := new(CSL)
	csl.builtins = make(map[string]operation)
	if registerDefaults {
		csl.RegisterDefaults()
	}
	return csl
}

func NewCSLPackageManager(scriptFetcher ScriptFetcher, allowInsecureRequests bool) *CSLPackageManager {
	cslpm := new(CSLPackageManager)
	if scriptFetcher == nil {
		cslpm.scriptFetcher = DefaultFetcher
	} else {
		cslpm.scriptFetcher = scriptFetcher
	}
	cslpm.allowInsecureRequests = allowInsecureRequests
	return cslpm
}
