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
	Birthday   int64        `json:"birthday"`
	SchoolName string       `json:"school_name"`
	BriefInfo  string       `json:"brief_info"`
	Company    string       `json:"company"`
	Position   string       `json:"position"`
	Specialty  string       `json:"specialty"`
	HeadImgUrl string       `json:"head_img_url"`
	CoachId    int64        `json:"coach_id"`
	Coach      *Coach       `json:"coach"`
	Applys     []*Apply     `json:"applys"`
	Shares     []*Share     `json:"shares"`
	Follows    []*Attention `json:"follows"`
	Fanss      []*Attention `json:"fanss"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

/**
 * 新增用户
 */
func (user *User) Add() (id int64, err error) {
	id, err = db.Insert("user", user)
	return
}

/**
 * 更新用户
 */
func (user *User) Update() (err error) {
	err = db.Update("user", user)
	return
}

/**
 * 报名
 */
func (user *User) Apply(apply *Apply) (id int64, err error) {
	id, err = db.Insert("apply", apply)
	return
}

/**
 * 发布分享
 */
func (user *User) AddShare(share *Share) (id int64, err error) {
	id,err = db.Insert("share", share)
	if err == nil {
		user.Shares = append(user.Shares, share)
	}
	return
}

/**
 * 新增关注
 */
func (user *User) Follow(toUser *User) (id int64, err error) {
	follow := &Attention{}
	err = db.QueryOne(follow, "select * from attention where user_id=? and to_user_id=?", user.Id, toUser.Id)
	Check(err)
	reFollow := &Attention{}
	err = db.QueryOne(reFollow, "select * from attention where user_id=? and to_user_id=? and state=1", toUser.Id, user.Id)
	follow.State = 1
	if follow.Id == 0 {
		follow.UserId = user.Id
		follow.ToUserId = toUser.Id
	}
	if reFollow.Id > 0 {
		follow.Relation = 2
		reFollow.Relation = 2
		db.Insert("attention", reFollow)
	}
	err = db.Update("attention", follow)
	if err == nil {
		user.Follows = append(user.Follows, follow)
		id = follow.Id
	}
	return
}

/**
 * 取消关注
 */
func (user *User) UnFollow(toUser *User) (err error) {
	follow := &Attention{}
	err = db.QueryOne(follow, "select * from attention where user_id=? and to_user_id=?", user.Id, toUser.Id)
	Check(err)
	if follow.Id == 0 { // 如果没有关注过 返回
		return
	}
	// 取消关注逻辑
	follow.State = 0
	reFollow := &Attention{}
	err = db.QueryOne(reFollow, "select * from attention where user_id=? and to_user_id=? and state=1", toUser.Id, user.Id)
	if reFollow.Id > 0 { // 如果对方关注了我 互相关注状态更新为 单向关注
		follow.Relation = 1
		reFollow.Relation = 1
		db.Update("attention", reFollow)
	}
	err = db.Update("attention", follow)
	return
}

/**
 * 获取用户信息
 */
func (user *User) GetDetail() {
	err := db.QueryOne(user, "SELECT * from user where id=?", user.Id)
	Check(err)
}

func (user *User) Delete()  {
	user.State=1
	err := db.Update("user", user)
	Check(err)
}

/**
 * 获取用户的分享列表
 */
func (user *User) GetShares() {
	mps, err := db.QueryMaps("SELECT * from share where user_id=?", user.Id)
	Check(err)
	user.Shares = make([]*Share, 0)
	for i:=0; i<len(mps); i++ {
		s := new(Share)
		err := mps[i].Load(s)
		Check(err)
		user.Shares = append(user.Shares, s)
	}
}

/**
 * 获取用户的报名列表
 */
func (user *User) GetApplys() {
	mps, err := db.QueryMaps("SELECT * from apply where user_id=?", user.Id)
	Check(err)
	user.Applys = make([]*Apply, 0)
	for i:=0; i<len(mps); i++ {
		s := new(Apply)
		err := mps[i].Load(s)
		Check(err)
		user.Applys = append(user.Applys, s)
	}
}

/**
 * 获取用户的关注列表
 */
func (user *User) GetFollows() {
	mps,err := db.QueryMaps("SELECT * from attention a, user u where a.user_id=? and a.to_user_id=u.id", user.Id)
	Check(err)
	for i:=0; i<len(mps); i++ {
		attention := new(Attention)
		err := mps[i].Load(attention)
		Check(err)
		user.Follows = append(user.Follows, attention)
	}
}

/**
 * 获取用户的粉丝列表
 */
func (user *User) GetFanss() {
	mps, err := db.QueryMaps("SELECT * from attention a, user u where a.to_user_id=? and a.user_id=u.id", user.Id)
	Check(err)
	for i:=0; i<len(mps); i++ {
		attention := new(Attention)
		err := mps[i].Load(attention)
		Check(err)
		user.Fanss = append(user.Fanss, attention)
	}
}
