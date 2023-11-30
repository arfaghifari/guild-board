package adventurer

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var adv = model.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
}

func TestCreateAdventurer(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("INSERT INTO adventurer(name, rank) VALUES($1, $2)")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		adv model.Adventurer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success created an adventurer",
			fields: fields{
				db: db,
			},
			args: args{
				adv: adv,
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(adv.Name, adv.Rank).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.CreateAdventurer(tt.args.adv)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestUpdateAdventureRank(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("UPDATE adventurer SET rank = $1 WHERE id = $2")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		adv model.Adventurer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success updated rank an adventurer",
			fields: fields{
				db: db,
			},
			args: args{
				adv: adv,
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(adv.Rank, adv.ID).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.UpdateAdventurerRank(tt.args.adv)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestAddCompletedQuest(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("UPDATE adventurer SET completed_quest = completed_quest + 1 WHERE id = $1")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success added completed quest an adventurer",
			fields: fields{
				db: db,
			},
			args: args{
				ID: adv.ID,
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(adv.ID).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.AddCompletedQuest(tt.args.ID)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetAdventurer(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("SELECT name, rank, completed_quest FROM adventurer WHERE id = $1")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		outAdv  model.Adventurer
		wantErr bool
	}{
		{
			name: "success get an adventurer",
			fields: fields{
				db: db,
			},
			args: args{
				ID: adv.ID,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "rank", "completed_quest"}).
					AddRow(adv.Name, adv.Rank, adv.CompletedQuest)

				mock.ExpectQuery(query).WithArgs(adv.ID).WillReturnRows(rows)
			},
			outAdv:  adv,
			wantErr: false,
		},
		{
			name: "failed get an adventurer",
			fields: fields{
				db: db,
			},
			args: args{
				ID: adv.ID,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "rank", "completed_quest"})

				mock.ExpectQuery(query).WithArgs(adv.ID).WillReturnRows(rows)
			},
			outAdv:  model.Adventurer{ID: adv.ID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			res, err := r.GetAdventurer(tt.args.ID)
			assert.NotNil(t, res)
			assert.Equal(t, tt.outAdv, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}
