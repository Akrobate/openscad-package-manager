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

func TestGetRefFromDependencySpecString(t *testing.T) {
	result, _ := GetRefFromDependencySpecString("https://gitlab.com/openscad-modules/breadboard.git#Test")
	if result != "Test" {
		t.Errorf(
			"GetRefFromDependencySpecString(\"https://gitlab.com/openscad-modules/breadboard.git#Test\") = %s, attendu Test",
			result,
		)
	}
}

func TestGetURLFromDependencySpecString(t *testing.T) {
	result, _ := GetURLFromDependencySpecString("https://gitlab.com/openscad-modules/breadboard.git#Test")
	if result != "https://gitlab.com/openscad-modules/breadboard.git" {
		t.Errorf(
			"GetRefFromDependencySpecString(\"https://gitlab.com/openscad-modules/breadboard.git#Test\") = %s, attendu https://gitlab.com/openscad-modules/breadboard",
			result,
		)
	}
}
