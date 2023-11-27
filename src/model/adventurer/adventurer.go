package adventurer

type Adventurer struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Rank           int32  `json:"rank"`
	CompletedQuest string `json:"completed_quest"`
}
