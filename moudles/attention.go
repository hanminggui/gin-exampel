package moudles

type Attention struct {
	State		int		`json:"state"`
	Relation	int64	`json:"relation"`
	user 		*User
	CuAt
}

