package main

import (
	"errors"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	// Тест с одним аргументом
	expected := "hello"
	result := echo("hello")
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}

	// Тест с несколькими аргументами
	expected = "hello world"
	result = echo("hello", "world")
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}

	// Тест с пустыми аргументами
	expected = ""
	result = echo()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestKillProcess(t *testing.T) {
	// Тест на ошибку при завершении несуществующего процесса
	t.Run("Kill_NonexistentProcess", func(t *testing.T) {
		noneExistentPID := 99999 // Несуществующий PID
		err := kill(noneExistentPID)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}

func TestCd(t *testing.T) {
	tests := []struct {
		name     string
		dir      string
		expected error
	}{
		{
			name:     "Valid directory",
			dir:      "./../dev07/",
			expected: nil,
		},
		{
			name:     "Invalid directory",
			dir:      "/path/does/not/exist",
			expected: os.ErrNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cd(tt.dir)

			if !errors.Is(err, tt.expected) {
				t.Errorf("Expected error: %v, got: %v", tt.expected, err)
			}
		})
	}
}
