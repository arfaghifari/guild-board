package quest

type Quest struct {
	ID           int64  `json:"quest_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MinimumRank  int32  `json:"minimum_rank"`
	RewardNumber int32  `json:"reward_number"`
	Status       int32  `json:"status"`
}

type GetQuestByStatus struct {
	ID           int64  `json:"quest_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MinimumRank  int32  `json:"minimum_rank"`
	RewardNumber int32  `json:"reward_number"`
}

type TakenBy struct {
	QuestID      int64 `json:"quest_id"`
	AdventurerID int64 `json:"adv_id"`
}
