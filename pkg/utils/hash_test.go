package utils

import (
	"testing"
)

func TestHash(t *testing.T) {
	pass, _ := GenPassword("12345")
	pass_str := string(pass)
	print(pass_str)
	print(ComparePassword("$12$GLl9WEfPrgJvDUasNuJc/eKeN013tmojMzQcgOB1NuaxtG0ACU1rC", "12345"))
}
