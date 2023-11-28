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
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := regexp.QuoteMeta("INSERT INTO adventurer(name, rank) VALUES($1, $2)")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(adv.Name, adv.Rank).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.CreateAdventurer(adv)
	assert.NoError(t, err)

}

func TestUpdateAdventureRank(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := regexp.QuoteMeta("UPDATE adventurer SET rank = $1 WHERE id = $2")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(adv.Rank, adv.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateAdventurerRank(adv)
	assert.NoError(t, err)
}

func TestAddCompletedQuest(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := regexp.QuoteMeta("UPDATE adventurer SET completed_quest = completed_quest + 1 WHERE id = $1")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(adv.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.AddCompletedQuest(adv.ID)
	assert.NoError(t, err)
}

func TestGetAdventurer(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT name, rank, completed_quest FROM adventurer WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"name", "rank", "completed_quest"}).
		AddRow(adv.Name, adv.Rank, adv.CompletedQuest)

	mock.ExpectQuery(query).WithArgs(adv.ID).WillReturnRows(rows)

	user, err := repo.GetAdventurer(adv.ID)
	assert.NotNil(t, user)
	assert.Equal(t, adv, user)
	assert.NoError(t, err)
}
