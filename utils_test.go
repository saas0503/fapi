package api

import "testing"

func TestIfSlashPrefixString(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"api should be /api", "api", "/api"},
		{"Api/Users should be /api/users", "Api/Users", "/api/users"},
		{"api / users / aUth should be /api/users/auth", "api / users / aUth", "/api/users/auth"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := IfSlashPrefixString(tt.input)
			if ans != tt.want {
				t.Errorf("IfSlashPrefixString() = %v, want %v", ans, tt.want)
			}
		})
	}
}

func BenchmarkIfSlashPrefixString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IfSlashPrefixString("api")
	}
}

func TestToFormat(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"Api should be api", "Api", "api"},
		{"Api /U S ER shoud be api/user", "Api /U S ER", "api/user"},
		{"a P i / U s E r should be api/user", "a P i / U s E r", "api/user"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ToFormat(tt.input)
			if ans != tt.want {
				t.Errorf("ToFormat() = %v, want %v", ans, tt.want)
			}
		})
	}
}

func BenchmarkToFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToFormat("a P i / U s E r")
	}
}
