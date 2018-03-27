package moudles

type Apply struct {
	Id			int64	`json:"id"`
	ApplyType 	int		`json:"apply_type"`
	State 		int		`json:"state"`
	user 		*User
	CuAt
}
