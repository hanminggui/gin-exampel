package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
)

type Share struct {
	Id          int64    `json:"id"`
	Title       string   `json:"title"`
	StartAt     int64    `json:"start_at"`
	EndAt       int64    `json:"end_at"`
	Amount      int      `json:"amount"`
	Type        int      `json:"type"`
	AuditState  int      `json:"audit_state"`
	LookesState int      `json:"lookes_state"`
	IsDelete    int      `json:"is_delete"`
	UserId      int64    `json:"user_id"`
	User        *User    `json:"user"`
	Applys      []*Apply `json:"applys"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

/**
 * 获取分享信息
 */
func (share *Share) GetDetail() {
	err := db.QueryOne(share, "SELECT * from share where id=?", share.Id)
	Check(err)
}

/**
 * 获取报名列表
 */
func (share *Share) GetApplys() {
	share.Applys = make([]*Apply, 0)
	maps,err := db.QueryMaps("SELECT * FROM apply WHERE share_id=?", share.Id)
	Check(err)
	for i:=0; i<len(maps); i++ {
		apply := new(Apply)
		err = maps[i].Load(apply)
		Check(err)
		share.Applys = append(share.Applys, apply)
	}
}

func (share *Share) Add() (id int64, err error) {
	id,err = db.Insert("share", share)
	return
}

func (share *Share) Delete()  {
	share.IsDelete=1
	err := db.Update("share", share)
	Check(err)
}

func GetShares(limit, offset int) (shares []*Share) {
	shares = make([]*Share, 0)
	maps,err := db.QueryMaps("SELECT * FROM share limit ?,?", offset, limit)
	Check(err)
	for i:=0; i<len(maps); i++ {
		share := new(Share)
		err = maps[i].Load(share)
		Check(err)
		shares = append(shares, share)
	}
	return
}