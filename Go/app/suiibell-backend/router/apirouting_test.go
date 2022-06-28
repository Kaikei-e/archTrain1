package router

import (
	"github.com/labstack/echo"
	"testing"
)

func TestAuthRouting(t *testing.T) {

	e := echo.New()

	type args struct {
		e *echo.Echo
	}

	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{name: "pass1", args: args{e: e}, hasError: false},
		{name: "fail1", args: args{e: nil}, hasError: true},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer func() { recover() }()

			if err := AuthRouting(tt.args.e); (err != nil) != tt.hasError {
				t.Errorf("AuthRouting() error = %v, wantErr %v", err, tt.hasError)
			}
		})
	}

}
