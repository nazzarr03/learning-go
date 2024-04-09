package email_validator

import (
	"testing"
)

var validEmails = []string {
	"a@test.com",
	"postmaster@test.com",
	"president@kremlin.gov.ru",
	"test@test.co.uk",
}

var invalidEmails = []string {
	"",
	"test",
	"test@test",
	"test@test.",
	"test@test.c",
	"test@test.com",
	"space between@test.com",
	"test@space between.com",
	"test@.com",
	".com",
	"test@",
}

func TestIsValidEmail(t *testing.T) {
	for i, v := range validEmails {
		if !IsValidEmail(v) {
			t.Errorf("%d: didn't accept valid email: %s", i, v)
		}
	}

	for i, v := range invalidEmails {
		if IsValidEmail(v) {
			t.Errorf("%d: accepted invalid email: %s", i, v)
		}
	}
}	