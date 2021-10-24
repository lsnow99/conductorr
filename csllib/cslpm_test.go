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
