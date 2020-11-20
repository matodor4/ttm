package main

import (
	"testing"
)

func TestGetPort(t *testing.T) {
	t.Run("test #1", func(t *testing.T) {
		expected := 4643
		_ = GetPort(JsonPath, &port)

		if port.Port != expected {
			t.Errorf("expected %d but got %d", expected, port.Port)
		}
	})
}
