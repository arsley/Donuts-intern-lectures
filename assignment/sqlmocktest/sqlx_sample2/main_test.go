package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestDoubleCoin(t *testing.T) {
	type args struct {
		userID int64
	}
	type mockParams struct {
		coin int64
	}
	tests := []struct {
		name       string
		args       args
		mockParams mockParams
		setupMock  func(sqlmock.Sqlmock, args, mockParams)
		wantErr    bool
	}{
		{
			name: "所有コインを2倍にする",
			args: args{
				userID: 100,
			},
			mockParams: mockParams{
				coin: 100,
			},
			setupMock: func(mock sqlmock.Sqlmock, args args, mockParams mockParams) {
				mock.ExpectBegin()
				mock.ExpectQuery(
					"SELECT .+ FROM user_coin WHERE .+ FOR UPDATE",
				).WithArgs(
					args.userID,
				).WillReturnRows(
					sqlmock.NewRows([]string{"coin"}).AddRow(
						mockParams.coin,
					),
				)
				mock.ExpectExec(
					"UPDATE user_coin SET .+ WHERE .+",
				).WithArgs(
					mockParams.coin*2,
					args.userID,
				).WillReturnResult(
					sqlmock.NewResult(0, 1),
				)
				mock.ExpectCommit()
			},
		},
		{
			name: "コインがない時はエラー",
			args: args{
				userID: 100,
			},
			mockParams: mockParams{
				coin: 100,
			},
			setupMock: func(mock sqlmock.Sqlmock, args args, mockParam mockParams) {
				mock.ExpectBegin()
				mock.ExpectQuery(
					"SELECT .+ FROM user_coin WHERE .+ FOR UPDATE",
				).WithArgs(
					args.userID,
				).WillReturnError(
					sql.ErrNoRows,
				)
				mock.ExpectRollback()
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
			tt.setupMock(mock, tt.args, tt.mockParams)
			dbx := sqlx.NewDb(db, "mysql")
			if err := DoubleCoin(dbx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DoubleCoin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestCreateOrUpdateUser(t *testing.T) {
	type args struct {
		name         string
		serviceToken string
	}
	type mockParams struct {
		id            int64
		affectedRows1 int64
		affectedRows2 int64
	}
	tests := []struct {
		name       string
		args       args
		mockParams mockParams
		setupMock  func(sqlmock.Sqlmock, args, mockParams)
		wantErr    bool
	}{
		{
			name: "ユーザー生成できる",
			args: args{
				name:         "ユーザー1",
				serviceToken: "token1",
			},
			mockParams: mockParams{
				id:            100,
				affectedRows1: 1,
			},
			setupMock: func(mock sqlmock.Sqlmock, args args, mockParams mockParams) {
				mock.ExpectBegin()
				mock.ExpectExec(
					"INSERT INTO user .+ VALUES .+",
				).WithArgs(args.name).WillReturnResult(sqlmock.NewResult(mockParams.id, 1))
				mock.ExpectExec(
					"INSERT INTO user_service .+ VALUES .+",
				).WithArgs(
					mockParams.id, args.serviceToken, mockParams.id,
				).WillReturnResult(sqlmock.NewResult(0, mockParams.affectedRows1))
				mock.ExpectCommit()
			},
		},
		{
			name: "既存サービストークンと重複した場合, 重複データを書き変える",
			args: args{
				name:         "ユーザー1",
				serviceToken: "token1",
			},
			mockParams: mockParams{
				id:            100,
				affectedRows1: 2,
				affectedRows2: 1,
			},
			setupMock: func(mock sqlmock.Sqlmock, args args, mockParams mockParams) {
				mock.ExpectBegin()
				mock.ExpectExec(
					"INSERT INTO user .+ VALUES .+",
				).WithArgs(args.name).WillReturnResult(sqlmock.NewResult(mockParams.id, 1))
				mock.ExpectExec(
					"INSERT INTO user_service .+ VALUES .+",
				).WithArgs(
					mockParams.id, args.serviceToken, mockParams.id,
				).WillReturnResult(sqlmock.NewResult(0, mockParams.affectedRows1))
				mock.ExpectExec(
					"UPDATE service_link SET .+",
				).WithArgs(
					mockParams.id, args.serviceToken,
				).WillReturnResult(
					sqlmock.NewResult(0, mockParams.affectedRows1),
				)
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.setupMock(mock, tt.args, tt.mockParams)
			dbx := sqlx.NewDb(db, "mysql")
			if err := CreateOrUpdateUser(dbx, tt.args.name, tt.args.serviceToken); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}
