package main

import "testing"

func TestFoo(t *testing.T) {
	initMap()
	cases := []struct {
		number         int
		expectedResult string
	}{
		{1, "one"},
		{5, "five"},
		{12, "twelve"},
		{56, "fifty six"},
		{1012, "one thousand and twelve"},
		{444, "four hundred forty four"},
		{404, "four hundred and four"},
		{1404, "one thousand four hundred and four"},
		{1000, "one thousand"},
		{9999, "nine thousand nine hundred ninety nine"},
		{100001, "one hundred thousand and one"},
		{-5, outOfRangeMessage},
		{1000000, outOfRangeMessage},
	}

	for _, currentCase := range cases {
		result := doFullNumber(currentCase.number)
		if result != currentCase.expectedResult {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, currentCase.expectedResult)
		}
	}
}
