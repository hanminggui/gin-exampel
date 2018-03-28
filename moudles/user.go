package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
)

type User struct {
	Id         int64        `json:"id"`
	NickName   string       `json:"nick_name"`
	Sex        int          `json:"sex"`
	Stage      int          `json:"stage"`
	State      int          `json:"state"`
	Type       int          `json:"type"`
	UserType   int          `json:"user_type"`
	Birthday   int64        `json:"birthday"`
	SchoolName string       `json:"school_name"`
	BriefInfo  string       `json:"brief_info"`
	Company    string       `json:"company"`
	Position   string       `json:"position"`
	Specialty  string       `json:"specialty"`
	HeadImgUrl string       `json:"head_img_url"`
	CoachId    int64        `json:"coach_id"`
	Coach      *Coach       `json:"coach"`
	Applys     []*Share     `json:"applys"`
	Shares     []*Share     `json:"shares"`
	Follows    []*Attention `json:"follows"`
	Fanss      []*Attention `json:"fanss"`
	CuAt
}

/**
 * 新增用户
 */
func (user *User) AddUser() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO user(nick_name, sex, stage, state, user_type, birthday, school_name, brief_info, company, position, "+
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
 * 报名
 */
func (user *User) Apply(share *Share, applyType int) (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO apply(share_id, user_id, apply_type, state) VALUES (?, ?, ?, ?)",
		share.Id, user.Id, applyType, 0)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

/**
 * 发布分享
 */
func (user *User) AddShare(share *Share) (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO share(title, start_at, end_at, amount, type, audit_state, lookes_state, is_delete) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		share.Title, share.StartAt, share.EndAt, share.Amount, share.Type, share.AuditState, share.LookesState, share.IsDelete)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	share.Id = id
	user.Shares = append(user.Shares, share)
	return
}

/**
 * 新增粉丝
 */
func (user *User) AddFans(fans *Attention) (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO attention(user_id, to_user_id, relation) VALUES (?, ?, ?)", fans.User.Id, user.Id, 0)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	user.Fanss = append(user.Fanss, fans)
	return
}

/**
 * 新增关注
 */
func (user *User) AddFollow(follow *Attention) (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO attention(user_id, to_user_id, relation) VALUES (?, ?, ?)", user.Id, follow.User.Id, 0)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	user.Follows = append(user.Follows, follow)
	return
}

/**
 * 获取用户信息
 */
func (user *User) GetDetail() {
	row := db.SqlDB.QueryRow("SELECT * from user where id=?", user.Id)
	err := row.Scan(user.Id, user.NickName, user.Sex, user.Stage, user.State, user.UserType, user.Birthday, user.SchoolName, user.BriefInfo,
		user.Company, user.Position, user.Specialty, user.HeadImgUrl)
	if err != nil {
		panic(err)
	}
}

/**
 * 获取用户的分享列表
 */
func (user *User) GetShares() {
	rows, err := db.SqlDB.Query("SELECT * from share where user_id=?", user.Id)
	if err != nil {
		panic(err)
	}
	user.Shares = make([]*Share, 0)
	for rows.Next() {
		share := Share{}
		rows.Scan(share.Id, share.Title, share.StartAt, share.EndAt, share.Amount, share.Type, share.AuditState, share.LookesState, share.IsDelete)
		user.Shares = append(user.Shares, &share)
	}
}

/**
 * 获取用户的关注列表
 */
func (user *User) GetFollows() {
	rows, err := db.SqlDB.Query("SELECT * from attention a, user u where a.user_id=? and a.to_user_id=u.id", user.Id)
	if err != nil {
		panic(err)
	}
	user.Follows = make([]*Attention, 0)
	for rows.Next() {
		follow := Attention{}
		rows.Scan(follow.User.Id, follow.User.NickName, follow.User.Stage, follow.User.State, follow.User.UserType, follow.User.Birthday,
			follow.User.SchoolName, follow.User.BriefInfo, follow.User.Company, follow.User.Position, follow.User.Specialty, follow.User.HeadImgUrl,
			follow.State, follow.Relation)
		user.Follows = append(user.Follows, &follow)
	}
}

/**
 * 获取用户的粉丝列表
 */
func (user *User) GetFanss() {
	rows, err := db.SqlDB.Query("SELECT * from attention a, user u where a.to_user_id=? and a.user_id=u.id", user.Id)
	if err != nil {
		panic(err)
	}
	user.Fanss = make([]*Attention, 0)
	for rows.Next() {
		fans := Attention{}
		rows.Scan(fans.User.Id, fans.User.NickName, fans.User.Stage, fans.User.State, fans.User.UserType, fans.User.Birthday,
			fans.User.SchoolName, fans.User.BriefInfo, fans.User.Company, fans.User.Position, fans.User.Specialty, fans.User.HeadImgUrl,
			fans.State, fans.Relation)
		user.Fanss = append(user.Fanss, &fans)
	}
}
