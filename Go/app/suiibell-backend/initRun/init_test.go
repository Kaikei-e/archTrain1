package initRun

import "testing"

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}
