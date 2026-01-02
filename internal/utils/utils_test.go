package utils

import "testing"

func TestGetNameFromDependencySpecString(t *testing.T) {
	result, _ := GetNameFromDependencySpecString("https://gitlab.com/openscad-modules/breadboard.git")
	if result != "breadboard" {
		t.Errorf(
			"GetNameFromDependencySpecString(\"https://gitlab.com/openscad-modules/breadboard.git\") = %s, attendu breadboard",
			result,
		)
	}
}
