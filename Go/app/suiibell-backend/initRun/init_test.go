package initRun

import "testing"

func TestInit(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "pass1", wantErr: false},
		{name: "fail1", wantErr: true},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer func() { recover() }()

			if err := Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
