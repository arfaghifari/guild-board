package quest

import (
	"errors"
	"testing"

	constant "github.com/arfaghifari/guild-board/src/constant"
	modelAdv "github.com/arfaghifari/guild-board/src/model/adventurer"
	model "github.com/arfaghifari/guild-board/src/model/quest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var adv = modelAdv.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
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
		MinimumRank:  12,
		RewardNumber: 200000,
		Status:       constant.AvailableQuest,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
		Status:       constant.CompletedQuest,
	},
	{
		ID:           4,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  12,
		RewardNumber: 200000,
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

func TestNewUsecase(t *testing.T) {
	res, err := NewUsecase()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestCreateQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mock     func(*MockRepository)
		outQuest model.Quest
		wantErr  bool
	}{
		{
			name: "success created a quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().CreateQuest(bulkQuest[0]).Return(bulkQuest[0], nil).Times(1)
			},
			outQuest: bulkQuest[0],
			wantErr:  false,
		},
		{
			name: "failed created a quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().CreateQuest(bulkQuest[0]).Return(model.Quest{}, errors.New("any error")).Times(1)
			},
			outQuest: model.Quest{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			res, err := u.CreateQuest(tt.args.quest)
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
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository)
		wantErr bool
	}{
		{
			name: "success updated a quest rank",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().UpdateQuestRank(bulkQuest[0]).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "failed updated a quest rank",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().UpdateQuestRank(bulkQuest[0]).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			err := u.UpdateQuestRank(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestUpdateQuestReward(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository)
		wantErr bool
	}{
		{
			name: "success updated a quest reward",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().UpdateQuestReward(bulkQuest[0]).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "failed updated a quest reward",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().UpdateQuestReward(bulkQuest[0]).Return(errors.New("any errors")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			err := u.UpdateQuestReward(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestDeleteQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		quest model.Quest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository)
		wantErr bool
	}{
		{
			name: "success deleted a quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().DeleteQuest(bulkQuest[0]).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "failed deleted a quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				quest: bulkQuest[0],
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().DeleteQuest(bulkQuest[0]).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			err := u.DeleteQuest(tt.args.quest)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetQuestByStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		status int32
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mock     func(*MockRepository)
		outQuest []model.GetQuestByStatus
		outLen   int
		wantErr  bool
	}{
		{
			name: "success get empty available quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				status: constant.AvailableQuest,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAllAvailableQuest().Return([]model.GetQuestByStatus{}, nil).Times(1)
			},
			outLen:   0,
			outQuest: []model.GetQuestByStatus{},
			wantErr:  false,
		},
		{
			name: "success get 2 available quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				status: constant.AvailableQuest,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAllAvailableQuest().Return(bulkQuestByStatus[:2], nil).Times(1)
			},
			outQuest: bulkQuestByStatus[:2],
			outLen:   2,
			wantErr:  false,
		},
		{
			name: "failed available quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				status: constant.AvailableQuest,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAllAvailableQuest().Return([]model.GetQuestByStatus{}, errors.New("any error")).Times(1)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  true,
		},
		{
			name: "success get none completed quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				status: constant.CompletedQuest,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAllCompletedQuest().Return([]model.GetQuestByStatus{}, nil).Times(1)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  false,
		},
		{
			name: "success get 1 completed quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				status: constant.CompletedQuest,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAllCompletedQuest().Return(bulkQuestByStatus[2:], nil).Times(1)
			},
			outQuest: bulkQuestByStatus[2:],
			outLen:   1,
			wantErr:  false,
		},
		{
			name: "failed completed quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				status: constant.CompletedQuest,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAllCompletedQuest().Return([]model.GetQuestByStatus{}, errors.New("any error")).Times(1)
			},
			outQuest: []model.GetQuestByStatus{},
			outLen:   0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			res, err := u.GetQuestByStatus(tt.args.status)
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

func TestTakeQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
		a *AdvMockRepository
	}
	type args struct {
		quest_id int64
		adv_id   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository, *AdvMockRepository)
		wantErr bool
	}{
		{
			name: "success took a quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(1)).Return(bulkQuest[0], nil).Times(1)
				advRepo.EXPECT().GetAdventurer(int64(1)).Return(adv, nil).Times(1)
				repo.EXPECT().CreateTakenBy(int64(1), int64(1)).Return(nil).Times(1)
				copyQuest := bulkQuest[0]
				copyQuest.Status = constant.WorkingQuest
				repo.EXPECT().UpdateQuestStatus(copyQuest).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "failed took a quest because taken",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 3,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(3)).Return(bulkQuest[2], nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "failed took a quest because not enough rank",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 2,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(2)).Return(bulkQuest[1], nil).Times(1)
				advRepo.EXPECT().GetAdventurer(int64(1)).Return(adv, nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "failed took a quest because no quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(1)).Return(model.Quest{}, errors.New("err")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "failed took a quest because no adv",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(1)).Return(bulkQuest[0], nil).Times(1)
				advRepo.EXPECT().GetAdventurer(int64(1)).Return(modelAdv.Adventurer{}, errors.New("err")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "error add taken by to db",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(1)).Return(bulkQuest[0], nil).Times(1)
				advRepo.EXPECT().GetAdventurer(int64(1)).Return(adv, nil).Times(1)
				repo.EXPECT().CreateTakenBy(int64(1), int64(1)).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "error update status quest db",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id: 1,
				adv_id:   1,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().GetQuest(int64(1)).Return(bulkQuest[0], nil).Times(1)
				advRepo.EXPECT().GetAdventurer(int64(1)).Return(adv, nil).Times(1)
				repo.EXPECT().CreateTakenBy(int64(1), int64(1)).Return(nil).Times(1)
				copyQuest := bulkQuest[0]
				copyQuest.Status = constant.WorkingQuest
				repo.EXPECT().UpdateQuestStatus(copyQuest).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo:    tt.fields.r,
				repoAdv: tt.fields.a,
			}
			tt.mock(tt.fields.r, tt.fields.a)
			err := u.TakeQuest(tt.args.quest_id, tt.args.adv_id)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestReportQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
		a *AdvMockRepository
	}
	type args struct {
		quest_id     int64
		adv_id       int64
		is_completed bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository, *AdvMockRepository)
		wantErr bool
	}{
		{
			name: "report completed quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: true,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[3].ID).Return(bulkQuest[3], nil).Times(1)
				repo.EXPECT().UpdateQuestStatus(model.Quest{ID: bulkQuest[3].ID, Status: constant.CompletedQuest}).Return(nil).Times(1)
				advRepo.EXPECT().AddCompletedQuest(int64(1)).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "quest not taken",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: true,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "quest not exist",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: true,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[3].ID).Return(bulkQuest[3], errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "quest status not working quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[0].ID,
				adv_id:       adv.ID,
				is_completed: true,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[0].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[0].ID).Return(bulkQuest[0], nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "report completed quest failed update status",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: true,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[3].ID).Return(bulkQuest[3], nil).Times(1)
				advRepo.EXPECT().AddCompletedQuest(adv.ID).Return(nil).Times(1)
				repo.EXPECT().UpdateQuestStatus(model.Quest{ID: bulkQuest[3].ID, Status: constant.CompletedQuest}).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "report completed quest failed add completed test",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: true,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[3].ID).Return(bulkQuest[3], nil).Times(1)
				advRepo.EXPECT().AddCompletedQuest(adv.ID).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "report uncompleted quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: false,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[3].ID).Return(bulkQuest[3], nil).Times(1)
				repo.EXPECT().UpdateQuestStatus(model.Quest{ID: bulkQuest[3].ID, Status: constant.AvailableQuest}).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "report uncompleted quest failed",
			fields: fields{
				r: NewMockRepository(mockCtrl),
				a: NewAdvMockRepository(mockCtrl),
			},
			args: args{
				quest_id:     bulkQuest[3].ID,
				adv_id:       adv.ID,
				is_completed: false,
			},
			mock: func(repo *MockRepository, advRepo *AdvMockRepository) {
				repo.EXPECT().IsExistTakenBy(bulkQuest[3].ID, adv.ID).Return(nil).Times(1)
				repo.EXPECT().GetQuest(bulkQuest[3].ID).Return(bulkQuest[3], nil).Times(1)
				repo.EXPECT().UpdateQuestStatus(model.Quest{ID: bulkQuest[3].ID, Status: 0}).Return(errors.New("any error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo:    tt.fields.r,
				repoAdv: tt.fields.a,
			}
			tt.mock(tt.fields.r, tt.fields.a)
			err := u.ReportQuest(tt.args.quest_id, tt.args.adv_id, tt.args.is_completed)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetQuestActiveAdventurer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mock     func(*MockRepository)
		outQuest []model.Quest
		outLen   int
		wantErr  bool
	}{
		{
			name: "success get empty quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				ID: adv.ID,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetQuestActiveAdventurer(adv.ID).Return([]model.Quest{}, nil).Times(1)
			},
			outLen:   0,
			outQuest: []model.Quest{},
			wantErr:  false,
		},
		{
			name: "success get  quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				ID: adv.ID,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetQuestActiveAdventurer(adv.ID).Return(bulkQuest[3:], nil).Times(1)
			},
			outQuest: bulkQuest[3:],
			outLen:   1,
			wantErr:  false,
		},
		{
			name: "failed  quest",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				ID: adv.ID,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetQuestActiveAdventurer(adv.ID).Return([]model.Quest{}, errors.New("any error")).Times(1)
			},
			outQuest: []model.Quest{},
			outLen:   0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			res, err := u.GetQuestActiveAdventurer(tt.args.ID)
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
