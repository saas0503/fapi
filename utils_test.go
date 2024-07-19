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
