package adventurer

type Adventurer struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Rank           int32  `json:"rank"`
	CompletedQuest int32  `json:"completed_quest"`
}
