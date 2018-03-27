package moudles

type Coach struct {
	Id 			int64	`json:"id"`
	Level 		int		`json:"level"`
	Type 		int		`json:"type"`
	PicsUrl 	string	`json:"pics_url"`
	State 		int		`json:"state"`
	CheckDesc 	string	`json:"check_desc"`
	Specialty 	string	`json:"specialty"`
	User 		User
	CuAt
}
