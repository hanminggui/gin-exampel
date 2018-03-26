package moudles

type Share struct {
	Id 			int64	`json:"id"`
	//UserId 		int64	`json:"user_id"`
	Title 		string	`json:"title"`
	StartAt 	int64	`json:"start_at"`
	EndAt 		int64	`json:"end_at"`
	Amount 		int		`json:"amount"`
	Type 		int		`json:"type"`
	AuditState 	int		`json:"audit_state"`
	LookesState int		`json:"lookes_state"`
	IsDelete 	int		`json:"is_delete"`
	Applys		[]Apply
	CuAt
}
