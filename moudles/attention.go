package moudles

type Attention struct {
	Id 			int64	`json:"id"`
	State 		int		`json:"state"`
	//UserId 		int64	`json:"user_id"`
	//ToUserId 	int64	`json:"to_user_id"`
	Relation 	int		`json:"relation"`
	CuAt
}
