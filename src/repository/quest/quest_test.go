package quest

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	constant "github.com/arfaghifari/guild-board/src/constant"
	modelAdv "github.com/arfaghifari/guild-board/src/model/adventurer"
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
		Status:       constant.AvailableQuest,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  11,
		RewardNumber: 200000,
		Status:       constant.WorkingQuest,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
		Status:       constant.WorkingQuest,
	},
}
var bulkQuestByStatus = []model.GetQuestByStatus{
	{
		ID:           1,
		Name:         "menyelamatkan kucing",
		Description:  "menyelamatkan kucing yang terjebak di atas pohon",
		MinimumRank:  11,
		RewardNumber: 200000,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  11,
		RewardNumber: 200000,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
	},
}
var adv = modelAdv.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
}

func TestClose(t *testing.T) {
	db, _ := NewMock()
	r := repository{
		db: db,
	}
	r.Close()
}

func TestNewRepository(t *testing.T) {
	res, err := NewRepository()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetAllCompletedQuest(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("SELECT quest_id, name, description, minimum_rank, reward_number FROM quest WHERE status = $1")
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name     string
		fields   fields
		mock     func()
		outQuest []model.GetQuestByStatus
		outLen   int
		wantErr  bool
	}{
		{
			name: "success get quest",
			fields: fields{
				db: db,
			},

			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
					AddRow(bulkQuestByStatus[2].ID, bulkQuestByStatus[2].Name, bulkQuestByStatus[2].Description, bulkQuestByStatus[2].MinimumRank, bulkQuestByStatus[2].RewardNumber)

				mock.ExpectQuery(query).WithArgs(constant.CompletedQuest).WillReturnRows(rows)
			},
			outQuest: bulkQuestByStatus[2:],
			outLen:   1,
			wantErr:  false,
		},
		{
			name: "none quest",
			fields: fields{
				db: db,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"})

				mock.ExpectQuery(query).WithArgs(constant.CompletedQuest).WillReturnRows(rows)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  false,
		},
		{
			name: "failed query",
			fields: fields{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery(query).WithArgs(constant.CompletedQuest).WillReturnError(sql.ErrConnDone)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  true,
		},
		{
			name: "failed scan query",
			fields: fields{
				db: db,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
					AddRow(bulkQuestByStatus[2].ID, nil, bulkQuestByStatus[2].Description, bulkQuestByStatus[2].MinimumRank, bulkQuestByStatus[2].RewardNumber)
				mock.ExpectQuery(query).WithArgs(constant.CompletedQuest).WillReturnRows(rows)

			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			res, err := r.GetAllCompletedQuest()
			assert.NotNil(t, res)
			assert.Len(t, tt.outQuest, tt.outLen)
			assert.Equal(t, tt.outQuest, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetAllAvailabeTest(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("SELECT quest_id, name, description, minimum_rank, reward_number FROM quest WHERE status = $1")
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name     string
		fields   fields
		mock     func()
		outQuest []model.GetQuestByStatus
		outLen   int
		wantErr  bool
	}{
		{
			name: "success get 1 quest",
			fields: fields{
				db: db,
			},

			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
					AddRow(bulkQuestByStatus[0].ID, bulkQuestByStatus[0].Name, bulkQuestByStatus[0].Description, bulkQuestByStatus[0].MinimumRank, bulkQuestByStatus[0].RewardNumber)

				mock.ExpectQuery(query).WithArgs(constant.AvailableQuest).WillReturnRows(rows)
			},
			outQuest: bulkQuestByStatus[:1],
			outLen:   1,
			wantErr:  false,
		},
		{
			name: "success get 2 quest",
			fields: fields{
				db: db,
			},

			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
					AddRow(bulkQuestByStatus[0].ID, bulkQuestByStatus[0].Name, bulkQuestByStatus[0].Description, bulkQuestByStatus[0].MinimumRank, bulkQuestByStatus[0].RewardNumber).
					AddRow(bulkQuestByStatus[1].ID, bulkQuestByStatus[1].Name, bulkQuestByStatus[1].Description, bulkQuestByStatus[1].MinimumRank, bulkQuestByStatus[1].RewardNumber)

				mock.ExpectQuery(query).WithArgs(constant.AvailableQuest).WillReturnRows(rows)
			},
			outQuest: bulkQuestByStatus[:2],
			outLen:   2,
			wantErr:  false,
		},
		{
			name: "none quest",
			fields: fields{
				db: db,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"})

				mock.ExpectQuery(query).WithArgs(constant.AvailableQuest).WillReturnRows(rows)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  false,
		},
		{
			name: "failed query",
			fields: fields{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery(query).WithArgs(constant.AvailableQuest).WillReturnError(sql.ErrConnDone)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  true,
		},
		{
			name: "failed scan query",
			fields: fields{
				db: db,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number"}).
					AddRow(bulkQuest[0].ID, nil, bulkQuest[0].Description, bulkQuest[0].MinimumRank, bulkQuest[0].RewardNumber)
				mock.ExpectQuery(query).WithArgs(constant.AvailableQuest).WillReturnRows(rows)

			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			res, err := r.GetAllAvailableQuest()
			assert.NotNil(t, res)
			assert.Len(t, tt.outQuest, tt.outLen)
			assert.Equal(t, tt.outQuest, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestCreateQuest(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("INSERT INTO quest(name, description, minimum_rank, reward_number) VALUES($1, $2, $3, $4) RETURNING quest_id")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mock     func()
		outQuest model.Quest
		wantErr  bool
	}{
		{
			name: "success created a quest",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id"}).
					AddRow(bulkQuest[0].ID)
				prep := mock.ExpectPrepare(query)
				prep.ExpectQuery().WithArgs(bulkQuest[0].Name, bulkQuest[0].Description, bulkQuest[0].MinimumRank, bulkQuest[0].RewardNumber).WillReturnRows(rows)
			},
			outQuest: bulkQuest[0],
			wantErr:  false,
		},
		{
			name: "failed created a quest query",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectQuery().WillReturnError(sql.ErrConnDone)
			},
			outQuest: model.Quest{},
			wantErr:  true,
		},
		{
			name: "failed created a quest prepare",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.WillReturnError(sql.ErrConnDone)
			},
			outQuest: model.Quest{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			res, err := r.CreateQuest(tt.args.quest)
			assert.Equal(t, tt.outQuest, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestUpdateQuestRank(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("UPDATE quest SET minimum_rank = $1 WHERE quest_id = $2")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success updated quest rank",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(bulkQuest[0].MinimumRank, bulkQuest[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "failed updated quest rank",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.UpdateQuestRank(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestUpdateQuestReward(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("UPDATE quest SET reward_number = $1 WHERE quest_id = $2")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success updated quest reward",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(bulkQuest[0].RewardNumber, bulkQuest[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "failed updated quest reward",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.UpdateQuestReward(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestUpdateQuestStatus(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("UPDATE quest SET status = $1 WHERE quest_id = $2")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success updated quest status",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(bulkQuest[0].Status, bulkQuest[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "failed updated quest status",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.UpdateQuestStatus(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestDeleteQuest(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("DELETE FROM quest WHERE quest_id = $1")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success deleted quest",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(bulkQuest[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "failed deleted quest",
			fields: fields{
				db: db,
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.DeleteQuest(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetQuest(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("SELECT name, description, minimum_rank, reward_number, status FROM quest WHERE quest_id = $1")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mock     func()
		outQuest model.Quest
		wantErr  bool
	}{
		{
			name: "success get a quest",
			fields: fields{
				db: db,
			},
			args: args{
				ID: bulkQuest[0].ID,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "description", "minimum_rank", "reward_number", "status"}).
					AddRow(bulkQuest[0].Name, bulkQuest[0].Description, bulkQuest[0].MinimumRank, bulkQuest[0].RewardNumber, bulkQuest[0].Status)
				mock.ExpectQuery(query).WithArgs(bulkQuest[0].ID).WillReturnRows(rows)
			},
			outQuest: bulkQuest[0],
			wantErr:  false,
		},
		{
			name: "failed get a quest",
			fields: fields{
				db: db,
			},
			args: args{
				ID: bulkQuest[0].ID,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "description", "minimum_rank", "reward_number", "status"})
				mock.ExpectQuery(query).WithArgs(bulkQuest[0].ID).WillReturnRows(rows)
			},
			outQuest: model.Quest{ID: bulkQuest[0].ID},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			res, err := r.GetQuest(tt.args.ID)
			assert.NotNil(t, res)
			assert.Equal(t, tt.outQuest, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestCreateTakenBy(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("INSERT INTO taken_by(quest_id, adv_id) VALUES($1, $2)")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest_id int64
		adv_id   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success took a quest",
			fields: fields{
				db: db,
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.ExpectExec().WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "success took a quest",
			fields: fields{
				db: db,
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func() {
				prep := mock.ExpectPrepare(query)
				prep.WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.CreateTakenBy(tt.args.quest_id, tt.args.adv_id)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestIsExistTakenBy(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("SELECT 1 FROM taken_by WHERE quest_id = $1 AND adv_id = $2")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		quest_id int64
		adv_id   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "success get",
			fields: fields{
				db: db,
			},
			args: args{
				quest_id: bulkQuest[0].ID,
				adv_id:   adv.ID,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"1"}).
					AddRow(1)
				mock.ExpectQuery(query).WithArgs(bulkQuest[0].ID, adv.ID).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "failed get",
			fields: fields{
				db: db,
			},
			args: args{
				quest_id: bulkQuest[0].ID,
				adv_id:   adv.ID,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"1"})
				mock.ExpectQuery(query).WithArgs(bulkQuest[0].ID).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			err := r.IsExistTakenBy(tt.args.quest_id, tt.args.adv_id)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetGetQuestActiveAdventurer(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()
	query := regexp.QuoteMeta("SELECT quest_id, name, description, minimum_rank, reward_number, status FROM quest NATURAL JOIN taken_by WHERE status = $1 AND adv_id = $2")
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name     string
		fields   fields
		mock     func()
		outQuest []model.Quest
		outLen   int
		wantErr  bool
	}{
		{
			name: "success get quest",
			fields: fields{
				db: db,
			},

			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number", "status"}).
					AddRow(bulkQuest[2].ID, bulkQuest[2].Name, bulkQuest[2].Description, bulkQuest[2].MinimumRank, bulkQuest[2].RewardNumber, bulkQuest[2].Status)

				mock.ExpectQuery(query).WithArgs(constant.WorkingQuest, adv.ID).WillReturnRows(rows)
			},
			outQuest: bulkQuest[2:],
			outLen:   1,
			wantErr:  false,
		},
		{
			name: "none quest",
			fields: fields{
				db: db,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number", "status"})

				mock.ExpectQuery(query).WithArgs(constant.WorkingQuest, adv.ID).WillReturnRows(rows)
			},
			outQuest: []model.Quest{},
			outLen:   0,
			wantErr:  false,
		},
		{
			name: "failed query",
			fields: fields{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery(query).WithArgs(constant.WorkingQuest, adv.ID).WillReturnError(sql.ErrConnDone)
			},
			outQuest: []model.Quest{},
			outLen:   0,
			wantErr:  true,
		},
		{
			name: "failed scan query",
			fields: fields{
				db: db,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"quest_id", "name", "description", "minimum_rank", "reward_number", "status"}).
					AddRow(bulkQuest[2].ID, nil, bulkQuest[2].Description, bulkQuest[2].MinimumRank, bulkQuest[2].RewardNumber, bulkQuest[2].Status)
				mock.ExpectQuery(query).WithArgs(constant.WorkingQuest, adv.ID).WillReturnRows(rows)
			},
			outQuest: []model.Quest{},
			outLen:   0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			tt.mock()
			res, err := r.GetQuestActiveAdventurer(adv.ID)
			assert.NotNil(t, res)
			assert.Len(t, tt.outQuest, tt.outLen)
			assert.Equal(t, tt.outQuest, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}
