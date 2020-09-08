package main

import "testing"

func TestWeb(t *testing.T) {
	resultado := greeting("Code.education Rocks!")
	if resultado != "<b>Code.education Rocks!</b>f" {
		t.Errorf("Esperava: %s, obteve: %s", "<b>Code.education Rocks!</b>", resultado)
	}
}
