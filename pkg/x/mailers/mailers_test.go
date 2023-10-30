package mailers

import (
	"testing"
)

func TestEmail(t *testing.T) {
	SendHTML("pkg/template/email_test.html", "iduchungho@gmail.com")
}
