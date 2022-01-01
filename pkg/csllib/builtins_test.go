package csllib

import "testing"

func TestRegisterFail(t *testing.T) {
	checkRegistration(t, "(", true)
	checkRegistration(t, ")", true)
	checkRegistration(t, `"`, true)
	checkRegistration(t, `13`, true)
	checkRegistration(t, `13G`, true)
	checkRegistration(t, `13Ki`, true)
	checkRegistration(t, `he lo`, true)
	checkRegistration(t, `he
	lo`, true)
}

func TestRegisterSucceed(t *testing.T) {
	checkRegistration(t, "hello", false)
	checkRegistration(t, "+", false)
	checkRegistration(t, "&", false)
}

func checkRegistration(t *testing.T, pattern string, shouldErr bool) {
	csl := NewCSL(true)
	err := csl.RegisterFunction(pattern, false, true, dummyEagerFn, nil)
	if shouldErr && err == nil {
		t.Fatalf("expected error from registering pattern %s but succeeded", pattern)
	} else if !shouldErr && err != nil {
		t.Fatalf("expected success from registering pattern %s but got error %v", pattern, err)
	}
}

func dummyEagerFn(env map[string]interface{}, args ...interface{}) (interface{}, error) {
	return nil, nil
}
