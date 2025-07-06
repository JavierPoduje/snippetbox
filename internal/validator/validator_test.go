package validator

import (
	"testing"
)

func TestValidator_Valid(t *testing.T) {
	t.Run("no field errors", func(t *testing.T) {
		v := Validator{}
		if !v.Valid() {
			t.Error("expected valid to be true")
		}
	})

	t.Run("with field errors", func(t *testing.T) {
		v := Validator{
			FieldErrors: map[string]string{"foo": "bar"},
		}
		if v.Valid() {
			t.Error("expected valid to be false")
		}
	})
}

func TestValidator_AddFieldError(t *testing.T) {
	t.Run("add error when no errors", func(t *testing.T) {
		v := Validator{}
		v.AddFieldError("foo", "bar")
		if len(v.FieldErrors) != 1 {
			t.Errorf("expected 1 field error, got %d", len(v.FieldErrors))
		}
	})

	t.Run("add error when errors exist", func(t *testing.T) {
		v := Validator{
			FieldErrors: map[string]string{"baz": "qux"},
		}
		v.AddFieldError("foo", "bar")
		if len(v.FieldErrors) != 2 {
			t.Errorf("expected 2 field errors, got %d", len(v.FieldErrors))
		}
	})

	t.Run("do not add error if key exists", func(t *testing.T) {
		v := Validator{
			FieldErrors: map[string]string{"foo": "bar"},
		}
		v.AddFieldError("foo", "baz")
		if v.FieldErrors["foo"] != "bar" {
			t.Errorf("expected error message to be 'bar', got '%s'", v.FieldErrors["foo"])
		}
	})
}

func TestValidator_CheckField(t *testing.T) {
	t.Run("add error when check is false", func(t *testing.T) {
		v := Validator{}
		v.CheckField(false, "foo", "bar")
		if len(v.FieldErrors) != 1 {
			t.Errorf("expected 1 field error, got %d", len(v.FieldErrors))
		}
	})

	t.Run("do not add error when check is true", func(t *testing.T) {
		v := Validator{}
		v.CheckField(true, "foo", "bar")
		if len(v.FieldErrors) != 0 {
			t.Errorf("expected 0 field errors, got %d", len(v.FieldErrors))
		}
	})
}

func TestNotBlank(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"not blank", "hello", true},
		{"blank", "", false},
		{"whitespace", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotBlank(tt.value); got != tt.want {
				t.Errorf("NotBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChars(t *testing.T) {
	tests := []struct {
		name  string
		value string
		n     int
		want  bool
	}{
		{"less than max", "hello", 10, true},
		{"equal to max", "hello", 5, true},
		{"more than max", "hello", 4, false},
		{"multibyte", "こんにちは", 5, true},
		{"multibyte more than max", "こんにちは", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxChars(tt.value, tt.n); got != tt.want {
				t.Errorf("MaxChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermittedValue(t *testing.T) {
	t.Run("string permitted", func(t *testing.T) {
		if !PermittedValue("a", "a", "b", "c") {
			t.Error("expected true")
		}
	})

	t.Run("string not permitted", func(t *testing.T) {
		if PermittedValue("d", "a", "b", "c") {
			t.Error("expected false")
		}
	})

	t.Run("int permitted", func(t *testing.T) {
		if !PermittedValue(1, 1, 2, 3) {
			t.Error("expected true")
		}
	})

	t.Run("int not permitted", func(t *testing.T) {
		if PermittedValue(4, 1, 2, 3) {
			t.Error("expected false")
		}
	})
}
