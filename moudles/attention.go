package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
)

type Attention struct {
	Id 		 int64 `json:"id"`
	State    int   `json:"state"`
	Relation int64 `json:"relation"`
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
	User     *User `json:"user"`
	ToUser   *User `json:"to_user"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

func (attention *Attention) GetDetail() {
	err := db.QueryOne(attention, "SELECT * from attention where id=?", attention.Id)
	Check(err)
}

/**
 * 关注别人
 */
func (attention *Attention) Follow() {
	follow := new(Attention)
	reFollow := new(Attention)
	db.QueryOne(follow, "SELECT * FROM attention WHERE user_id=? AND to_user_id=?", attention.UserId, attention.ToUserId)
	db.QueryOne(reFollow, "SELECT * FROM attention WHERE state=1 AND user_id=? AND to_user_id=?", attention.ToUserId,
		attention.UserId)
	if reFollow.Id > 0 { //  如果对方也关注了我，把关系设置为互相关注
		attention.Relation = 2
		reFollow.Relation = 2
		err := db.Update("attention", reFollow)
		Check(err)
	}
	if follow.Id > 0 { // 如果 我曾经关注过他，或者现在就是关注状态，修改关注状态为正常状态 重新保存一下
		follow.State = 1
		follow.Relation = 2
		err := db.Update("attention", follow)
		Check(err)
		attention.Id = follow.Id
	} else { // 如果是新数据 直接插入
		id,err := db.Insert("attention", attention)
		Check(err)
		attention.Id = id
	}
	attention.GetDetail()
}

/**
 * 取消关注
 */
func (attention *Attention) UnFollow() {
	// 直接更新 关注状态
	attention.State = 0
	attention.Relation = 1
	err := db.Update("attention", attention)
	Check(err)

	reFollow := new(Attention)
	db.QueryOne(reFollow, "SELECT * FROM attention WHERE state=1 AND user_id=? AND to_user_id=?", attention.ToUserId,
		attention.UserId)
	if reFollow.Id > 0 { //  如果对方也关注了我，把关系变成单向关注
		reFollow.Relation = 1
		err := db.Update("attention", reFollow)
		Check(err)
	}
	attention.GetDetail()
}
