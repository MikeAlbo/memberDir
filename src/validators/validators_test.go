// unit test for the validators package
package validators

import "testing"

// ** length validation
// test that the length validator returns the tested value if true
func TestStringLengthValidatorToReturnValue(t *testing.T)  {
	testCases := []string{"one", "two", "0ne", "1en"}
	testLengthInput := 3

	for _,c := range testCases {
		res, err:= StringLengthValidator(c, testLengthInput)

		if res != c {
			t.Fatalf("Return Value: %s, does not equel the input %s", res,c)
		}

		if err != nil {
			t.Fatalf(" %s, Should have returned nil", c)
		}
	}
}

// test that the length validator returns an error if the lengths do not match
func TestStringLengthValidatorToReturnErrors(t *testing.T) {
	testCases := []string{"one", "two", "0ne", "1en"}
	testLengthInput := 2

	for _,c := range testCases {
		val, err := StringLengthValidator(c, testLengthInput);

		if err == nil {
			t.Fatalf("Input Value %s, should have thrown an error", c)
		}

		if val != "" {
			t.Fatalf("%s, Should have returnd an epty string", c)
		}
	}
}

// ** int validator
// test that the int validator returns the input value if it matches
func TestIntValidatorToReturnValue(t *testing.T)  {
	testCases := []string{"01","2","100"}

	for _,c := range testCases {
		if val, _ := IntValidator(c); val != c {
			t.Fatalf("Int Validator Returned %s, for input %s", val, c)
		}
	}
}

// test that the int validators returns an error if the value is of incorrect type
func TestIntValidatorToReturnError(t *testing.T)  {
	testCases := []string{"0l", "one", "thr33"}

	for _,c := range testCases {
		if _, err := IntValidator(c); err == nil {
			t.Fatalf("Int Validator should have returned an error for case: %s", c)
		}
	}
}



// ** char validator
// test that the char validator returns thr inout value if it matches
func TestCharValidatorToReturnValue(t *testing.T)  {
	testCases := []string{"Michael", "Hello", "a", "MI"}

	for _,c := range testCases {
		if val, _ := CharValidator(c); val != c {
			t.Fatalf("Char Validator Returned %s, for input %s", val, c)
		}
	}
}

// test that the char validator returns an error if the value is of incorrect type
func TestCharValidatorToReturnError(t *testing.T)  {
	testCases := []string{"0l", "1", "thr33"}

	for _,c := range testCases {
		if _, err := CharValidator(c); err == nil {
			t.Fatalf("Char Validator should have returned an error for case: %s", c)
		}
	}
}