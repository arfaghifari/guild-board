package quest

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/arfaghifari/guild-board/src/model/quest"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var bulkQuest = []model.Quest{
	{
		ID:           1,
		Name:         "menyelamatkan kucing",
		Description:  "menyelamatkan kucing yang terjebak di atas pohon",
		MinimumRank:  11,
		RewardNumber: 200000,
		Status:       0,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  11,
		RewardNumber: 200000,
		Status:       0,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
		Status:       2,
	},
}

func TestGetAllCompletedQuest(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	completedStatus := 2
	query := regexp.QuoteMeta("SELECT quest_id, name, description, minimum_rank, reward_number FROM quest WHERE status = $1")
	rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
		AddRow(bulkQuest[2].ID, bulkQuest[2].Name, bulkQuest[2].Description, bulkQuest[2].MinimumRank, bulkQuest[2].RewardNumber)

	mock.ExpectQuery(query).WithArgs(completedStatus).WillReturnRows(rows)

	quest, err := repo.GetAllCompletedQuest()
	assert.NotEmpty(t, quest)
	assert.NoError(t, err)
	assert.Len(t, quest, 1)
}

func TestGetAllAvailabeTest(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	availableStastus := 0
	query := regexp.QuoteMeta("SELECT quest_id, name, description, minimum_rank, reward_number FROM quest WHERE status = $1")
	rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
		AddRow(bulkQuest[0].ID, bulkQuest[0].Name, bulkQuest[0].Description, bulkQuest[0].MinimumRank, bulkQuest[0].RewardNumber).
		AddRow(bulkQuest[1].ID, bulkQuest[1].Name, bulkQuest[1].Description, bulkQuest[1].MinimumRank, bulkQuest[1].RewardNumber)

	mock.ExpectQuery(query).WithArgs(availableStastus).WillReturnRows(rows)

	quest, err := repo.GetAllAvailableQuest()
	assert.NotEmpty(t, quest)
	assert.NoError(t, err)
	assert.Len(t, quest, 2)
}

func TestCreateQuest(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	quest := bulkQuest[0]

	query := regexp.QuoteMeta("INSERT INTO quest(name, description, minimum_rank, reward_number) VALUES($1, $2, $3, $4)")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(quest.Name, quest.Description, quest.MinimumRank, quest.RewardNumber).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.CreateQuest(quest)
	assert.NoError(t, err)
}

func TestUpdateQuestRank(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	quest := bulkQuest[0]

	query := regexp.QuoteMeta("UPDATE quest SET minimum_rank = $1 WHERE quest_id = $2")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(quest.MinimumRank, quest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateQuestRank(quest)
	assert.NoError(t, err)
}

func TestUpdateQuestReward(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	quest := bulkQuest[0]

	query := regexp.QuoteMeta("UPDATE quest SET reward_number = $1 WHERE quest_id = $2")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(quest.RewardNumber, quest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateQuestReward(quest)
	assert.NoError(t, err)
}

func TestUpdateQuestStatus(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	quest := bulkQuest[0]

	query := regexp.QuoteMeta("UPDATE quest SET status = $1 WHERE quest_id = $2")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(quest.Status, quest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateQuestStatus(quest)
	assert.NoError(t, err)
}

func TestDeleteQuest(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	quest := bulkQuest[0]
	query := regexp.QuoteMeta("DELETE FROM quest WHERE quest_id = $1")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(quest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteQuest(quest)
	assert.NoError(t, err)
}

func TestGetQuest(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	quest := bulkQuest[0]
	query := regexp.QuoteMeta("SELECT name, description, minimum_rank, reward_number, status FROM quest WHERE quest_id = $1")

	rows := sqlmock.NewRows([]string{"name", "description", "minimum_rank", "reward_number", "status"}).
		AddRow(quest.Name, quest.Description, quest.MinimumRank, quest.RewardNumber, quest.Status)

	mock.ExpectQuery(query).WithArgs(quest.ID).WillReturnRows(rows)

	q, err := repo.GetQuest(quest.ID)
	assert.NotNil(t, q)
	assert.NoError(t, err)
	assert.Equal(t, quest, q)
}

func TestCreateTakenBy(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	var adv_id, quest_id int64
	quest_id = 1
	adv_id = 1

	query := regexp.QuoteMeta("INSERT INTO taken_by(quest_id, adv_id) VALUES($1, $2)")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(quest_id, adv_id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.CreateTakenBy(quest_id, adv_id)
	assert.NoError(t, err)
}
