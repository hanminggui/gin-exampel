package moudles

type User struct {
	Id 			int64	`json:"id"`
	NickName 	string	`json:"nick_name"`
	Sex 		int		`json:"sex"`
	SchoolName 	string	`json:"school_name"`
	Stage 		int		`json:"stage"`
	State 		int		`json:"state"`
	BriefInfo 	string	`json:"brief_info"`
	UserType 	int		`json:"user_type"`
	Birthday 	int64	`json:"birthday"`
	Company 	string	`json:"company"`
	Position 	string	`json:"position"`
	Specialty 	string	`json:"specialty"`
	HeadImgUrl 	string	`json:"head_img_url"`
	Applys 		[]Apply
	Shares		[]Share
	Follows		[]User
	Fanss		[]User
	CuAt
}