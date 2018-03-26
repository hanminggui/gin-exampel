package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
)

type User struct {
	Id 			int64	`json:"id"`
	NickName 	string	`json:"nick_name"`
	Sex 		int		`json:"sex"`
	Stage 		int		`json:"stage"`
	State 		int		`json:"state"`
	UserType 	int		`json:"user_type"`
	Birthday 	int64	`json:"birthday"`
	SchoolName 	string	`json:"school_name"`
	BriefInfo 	string	`json:"brief_info"`
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

/**
 * 新增用户
 */
func (user *User) AddUser() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO user(nick_name, sex, stage, state, user_type, birthday, school_name, brief_info, company, position, " +
		"specialty, head_img_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.NickName, user.Sex, user.Stage, user.State, user.UserType, user.Birthday, user.SchoolName, user.BriefInfo, user.Company, user.Position,
			user.Specialty, user.HeadImgUrl)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

/**
 * 获取用户信息
 */
func (user *User) GetDetail() {
	row := db.SqlDB.QueryRow("SELECT * from user where id=?", user.Id)
	err := row.Scan(&user.Id, user.NickName, user.Sex, user.Stage, user.State, user.UserType, user.Birthday, user.SchoolName, user.BriefInfo,
		user.Company, user.Position, user.Specialty, user.HeadImgUrl)
	if err != nil {
		panic(err)
	}
}

/**
 * 报名
 */
func (user *User) Apply(share *Share, apply_type int) (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO apply(share_id, user_id, apply_type, state) VALUES (?, ?, ?, ?)",
		share.Id, user.Id, apply_type, 0)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}