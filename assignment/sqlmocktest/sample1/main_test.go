package main

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFetchUser(t *testing.T) {
	type args struct {
		id int64
	}
	type mockParams struct {
		id   int64
		name string
	}
	tests := []struct {
		name       string
		args       args
		mockParams mockParams
		want       *User
		wantErr    bool
		domock     func(sqlmock.Sqlmock, args, mockParams)
	}{
		{
			name: "ユーザーデータを取得できる",
			args: args{
				id: 100,
			},
			mockParams: mockParams{
				id:   100,
				name: "name",
			},
			domock: func(mock sqlmock.Sqlmock, args args, mockParams mockParams) {
				mock.ExpectQuery(
					"SELECT .+ FROM user WHERE .+",
				).WithArgs(
					args.id,
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"id", "name",
					}).AddRow(
						mockParams.id, mockParams.name,
					),
				)
			},
			want: &User{
				ID:   100,
				Name: "name",
			},
		},
		{
			name: "When use incorrect id, gets nothing",
			args: args{
				id: 0,
			},
			domock: func(mock sqlmock.Sqlmock, args args, mockParams mockParams) {
				mock.ExpectQuery(
					"SELECT .+ FROM user WHERE .+",
				).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.domock(mock, tt.args, tt.mockParams)
			got, err := FetchUser(db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
