package config

import (
	"reflect"
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{name: "test get config", want: &Config{
			HostGRPC:      "localhost:3200",
			LogLevel:      "info",
			DSN:           "postgresql://postgres:password@10.66.66.3:5432/postgres?sslmode=disable",
			MigrationDir:  "./migrations",
			Files:         "./files",
			TokenLifetime: 30000 * time.Second,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("LOG_LEVEL", tt.want.LogLevel)
			t.Setenv("DATABASE_DSN", tt.want.DSN)
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
