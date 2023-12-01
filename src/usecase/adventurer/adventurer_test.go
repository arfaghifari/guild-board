package adventurer

import (
	"errors"
	"testing"

	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var adv = model.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
}

func TestNewUsecase(t *testing.T) {
	res, err := NewUsecase()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestCreateAdventurer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		adv model.Adventurer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository)
		outAdv  model.Adventurer
		wantErr bool
	}{
		{
			name: "success created an adventurer",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				adv: adv,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().CreateAdventurer(adv).Return(adv, nil).Times(1)
			},
			outAdv:  adv,
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				adv: adv,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().CreateAdventurer(adv).Return(model.Adventurer{}, errors.New("any error")).Times(1)
			},
			outAdv:  model.Adventurer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			res, err := u.CreateAdventurer(tt.args.adv)
			assert.Equal(t, tt.outAdv, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestUpdateAdventureRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		adv model.Adventurer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository)
		wantErr bool
	}{
		{
			name: "success updated rank an adventurer",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				adv: adv,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().UpdateAdventurerRank(adv).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				adv: adv,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().UpdateAdventurerRank(adv).Return(errors.New("any error")).Times(1)
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
			err := u.UpdateAdventurerRank(tt.args.adv)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}

func TestGetAdventurer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		r *MockRepository
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(*MockRepository)
		outAdv  model.Adventurer
		wantErr bool
	}{
		{
			name: "success get an adventurer",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				ID: adv.ID,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAdventurer(adv.ID).Return(adv, nil).Times(1)
			},
			outAdv:  adv,
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				r: NewMockRepository(mockCtrl),
			},
			args: args{
				ID: adv.ID,
			},
			mock: func(repo *MockRepository) {
				repo.EXPECT().GetAdventurer(adv.ID).Return(model.Adventurer{}, errors.New("any error")).Times(1)
			},
			outAdv:  model.Adventurer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo: tt.fields.r,
			}
			tt.mock(tt.fields.r)
			res, err := u.GetAdventurer(tt.args.ID)
			assert.Equal(t, tt.outAdv, res)
			if tt.wantErr {
				assert.Error(t, err, tt.name)
			} else {
				assert.NoError(t, err, tt.name)
			}
		})
	}
}
