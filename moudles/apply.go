package moudles

type Apply struct {
	Id			int64	`json:"id"`
	ShareId		int64	`json:"share_id"`
	UserId 		int64	`json:"user_id"`
	ApplyType 	int		`json:"apply_type"`
	State 		int		`json:"state"`
	CuAt
}
