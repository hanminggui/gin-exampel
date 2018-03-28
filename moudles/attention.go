package moudles

type Attention struct {
	State		int		`json:"state"`
	Relation	int64	`json:"relation"`
	UserId 		int64	`json:"user_id"`
	ToUserId 	int64	`json:"to_user_id"`
	User 		*User	`json:"user"`
	ToUser 		*User	`json:"to_user"`
	CuAt
}

