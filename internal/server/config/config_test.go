package config

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{name: "test get config", want: &Config{
			Host:       "localhost:8080",
			HostGRPC:   "localhost:3200",
			LogLevel:   "info",
			DSN:        "",
			SigningKey: "",
			CryptoKey:  "private.pem",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ADDRESS", tt.want.Host)
			t.Setenv("LOG_LEVEL", tt.want.LogLevel)
			t.Setenv("DATABASE_DSN", tt.want.SigningKey)
			t.Setenv("KEY", tt.want.SigningKey)
			t.Setenv("CRYPTO_KEY", tt.want.CryptoKey)
			got, err := GetConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}