package csllib

import (
	"testing"
)

func TestParseGit1(t *testing.T) {
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("gitlab.com/lsnow/mycsl:main.csl")
	if err != nil {
		t.Fatal(err)
	}
	gs, ok := importedScript.(GitScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	expectedGS := GitScript{
		host:     "gitlab.com",
		repo:     "lsnow/mycsl",
		filePath: "main.csl",
		version:  "HEAD",
	}
	checkGS(t, gs, expectedGS)
}

func TestParseGit2(t *testing.T) {
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("github.com/lsnow99/mycsl:/main.csl@715ea4ca5c982ce83f60323329305554481fab34")
	if err != nil {
		t.Fatal(err)
	}
	gs, ok := importedScript.(GitScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	expectedGS := GitScript{
		host:     "github.com",
		repo:     "lsnow99/mycsl",
		filePath: "/main.csl",
		version:  "715ea4ca5c982ce83f60323329305554481fab34",
	}
	checkGS(t, gs, expectedGS)
}

func TestParseGit3(t *testing.T) {
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("github.com/lsnow99/mycsl:scripts/main.csl")
	if err != nil {
		t.Fatal(err)
	}
	gs, ok := importedScript.(GitScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	expectedGS := GitScript{
		host:     "github.com",
		repo:     "lsnow99/mycsl",
		filePath: "scripts/main.csl",
		version:  "HEAD",
	}
	checkGS(t, gs, expectedGS)
}

func TestParseProfile1(t *testing.T) {
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("profile:filter:myprof?username=l&password=l&host=localhost&port=8282")
	if err != nil {
		t.Fatal(err)
	}
	ps, ok := importedScript.(ProfileScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	expectedPS := ProfileScript{
		name:       "myprof",
		scriptType: "filter",
		host:       "localhost",
		username:   "l",
		password:   "l",
		port:       8282,
	}
	checkPS(t, ps, expectedPS)
}

func TestParseProfile2(t *testing.T) {
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("profile:filter:myprof")
	if err != nil {
		t.Fatal(err)
	}
	ps, ok := importedScript.(ProfileScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	expectedPS := ProfileScript{
		name:       "myprof",
		scriptType: "filter",
		port:       8282,
	}
	checkPS(t, ps, expectedPS)
}

func TestParseProfile3(t *testing.T) {
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("profile:sorter:myprof?host=example.com&port=4982&auth_token=jklfdas")
	if err != nil {
		t.Fatal(err)
	}
	ps, ok := importedScript.(ProfileScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	expectedPS := ProfileScript{
		name:       "myprof",
		scriptType: "sorter",
		host:       "example.com",
		port:       4982,
		authToken:  "jklfdas",
	}
	checkPS(t, ps, expectedPS)
}

func TestParseWeb(t *testing.T) {
	const scriptURL = "https://raw.githubusercontent.com/lsnow99/myscripts/main/main.csl"
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport(scriptURL)
	if err != nil {
		t.Fatal(err)
	}
	ws, ok := importedScript.(WebScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	if ws.u.String() != scriptURL {
		t.Fatalf("got url %s, expected %s", ws.u.String(), scriptURL)
	}
}

func TestParseFileBasic(t *testing.T) {
	const filePath = "myscript.csl"
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	importedScript, err := cslpm.parseImport("file:" + filePath)
	if err != nil {
		t.Fatal(err)
	}
	fs, ok := importedScript.(FileScript)
	if !ok {
		t.Fatal("type assertion failed")
	}
	if fs.filePath != filePath {
		t.Fatalf("got file path %s, expected %s", fs.filePath, filePath)
	}
}

func checkGS(t *testing.T, gs, expectedGS GitScript) {
	if gs.host != expectedGS.host {
		t.Fatalf("got host %s, expected %s", gs.host, expectedGS.host)
	}
	if gs.repo != expectedGS.repo {
		t.Fatalf("got repo %s, expected %s", gs.repo, expectedGS.repo)
	}
	if gs.filePath != expectedGS.filePath {
		t.Fatalf("got filepath %s, expected %s", gs.filePath, expectedGS.filePath)
	}
	if gs.version != expectedGS.version {
		t.Fatalf("got version %s, expected %s", gs.version, expectedGS.version)
	}
}

func checkPS(t *testing.T, ps, expectedPS ProfileScript) {
	if ps.name != expectedPS.name {
		t.Fatalf("got name %s, expected %s", ps.name, expectedPS.name)
	}
	if ps.scriptType != expectedPS.scriptType {
		t.Fatalf("got scriptType %s, expected %s", ps.scriptType, expectedPS.scriptType)
	}
	if ps.host != expectedPS.host {
		t.Fatalf("got host %s, expected %s", ps.host, expectedPS.host)
	}
	if ps.username != expectedPS.username {
		t.Fatalf("got username %s, expected %s", ps.username, expectedPS.username)
	}
	if ps.password != expectedPS.password {
		t.Fatalf("got password %s, expected %s", ps.password, expectedPS.password)
	}
	if ps.authToken != expectedPS.authToken {
		t.Fatalf("got authToken %s, expected %s", ps.authToken, expectedPS.authToken)
	}
}
