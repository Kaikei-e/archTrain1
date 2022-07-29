package dbconn

import (
	"reflect"
	"suiibell/ent"
	"testing"
)

func TestDBConnection(t *testing.T) {
	tests := []struct {
		name    string
		want    *ent.Client
		wantErr bool
	}{
		// TODO: Add test cases.
		{"pass1", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DBConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("DBConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
