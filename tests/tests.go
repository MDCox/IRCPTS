package tests

type Test struct {
	name     string            // Name of the test
	criteria func(string) bool // The test itself that is run with eval
	Result   bool              // Result of test set during eval
	err      string            // Error message set during eval
	section  string            // Relevant protocol section
}

func (t *Test) Eval(msg string) bool {
	if crlf_missing(msg) {
		t.Result = false
		t.err = err_msg(
			t.name,
			msg,
			"Message does not end with a CRLF token (\\r\\n)",
			"2.3")
		return false
	}

	if too_long(msg) {
		t.Result = false
		t.err = err_msg(
			t.name,
			msg,
			"Message length exceeds 512 characters",
			"2.3")
		return false
	}

	if t.criteria(msg) {
		t.Result = true
		return true
	}

	t.Result = false
	t.err = err_msg(
		t.name,
		msg,
		t.err,
		t.section)
	return false
}

type testSet struct {
	section    string
	test_count int
	fail_count int
	pass_count int
	Tests      []Test
}

type Suite struct {
	Tests  []testSet
	Passed []Test
	Failed []Test
}

// Factory function that returns a brand new set of tests.
func NewTestSuite() Suite {
	tests := []testSet{
		connection_registration_tests(),
	}

	return Suite{
		Tests: tests, // All tests
	}
}
