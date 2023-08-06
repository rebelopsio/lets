package main

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Test_openDB(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openDB(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("openDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
