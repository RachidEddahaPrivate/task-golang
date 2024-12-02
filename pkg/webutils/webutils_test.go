package webutils

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func Test_customValidator_Validate(t *testing.T) {
	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Case",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Name  string `validate:"required"`
					Email string `validate:"required,email"`
				}{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Case",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Name  string `validate:"required"`
					Email string `validate:"required,email"`
				}{
					Name:  "",
					Email: "invalid-email",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &customValidator{
				validator: tt.fields.validator,
			}
			if err := cv.Validate(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
