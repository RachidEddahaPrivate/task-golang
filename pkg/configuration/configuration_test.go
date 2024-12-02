package configuration

import (
	"reflect"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	tests := []struct {
		name              string
		wantConfiguration Config
		wantErr           bool
	}{
		{
			name:    "simple test, cannot find file config",
			wantErr: true},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gotConfiguration, err := LoadConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfiguration, tt.wantConfiguration) {
				t.Errorf("LoadConfiguration() gotConfiguration = %v, want %v", gotConfiguration, tt.wantConfiguration)
			}
		})
	}
}
