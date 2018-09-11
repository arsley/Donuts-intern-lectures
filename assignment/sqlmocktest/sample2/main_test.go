package main

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConsumeCoin(t *testing.T) {
	type args struct {
		userID int64
		amount int64
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
			name: "コインを消費できる",
			args: args{
				userID: 100,
				amount: 1000,
			},
			mockParams: mockParams{
				coin: 10000,
			},
			setupMock: func(mock sqlmock.Sqlmock, args args, params mockParams) {
				mock.ExpectBegin()
				mock.ExpectQuery(
					"SELECT .+ FROM user_coin .+ FOR UPDATE",
				).WithArgs(
					args.userID,
				).WillReturnRows(
					sqlmock.NewRows([]string{"coin"}).AddRow(
						params.coin,
					),
				)
				mock.ExpectExec(
					"UPDATE user_coin SET .+",
				).WithArgs(
					args.amount, args.userID,
				).WillReturnResult(
					sqlmock.NewResult(0, 1),
				)
				mock.ExpectCommit()
			},
		},
		{
			name: "Beginに失敗する",
			args: args{
				userID: 100,
				amount: 1000,
			},
			mockParams: mockParams{
				coin: 10000,
			},
			setupMock: func(mock sqlmock.Sqlmock, args args, params mockParams) {
				mock.ExpectBegin().WillReturnError(driver.ErrBadConn)
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
			if err := ConsumeCoin(db, tt.args.userID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("ConsumeCoin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}
