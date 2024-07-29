package config

import "testing"

func TestLoad(t *testing.T) {
	cfg, err := Load(".")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cfg)
}
