package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
)

type Apply struct {
	Id        int64  `json:"id"`
	ApplyType int    `json:"apply_type"`
	State     int    `json:"state"`
	UserId    int64  `json:"user_id"`
	ShareId  int64  `json:"share_id"`
	User      *User  `json:"user"`
	Share     *Share `json:"share"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

func (apply *Apply) GetDetail() {
	err := db.QueryOne(apply, "SELECT * from apply where id=?", apply.Id)
	Check(err)
	apply.User = new(User)
	err = db.QueryOne(apply.User, "select * from user where id=?", apply.UserId)
	Check(err)
	apply.Share = new(Share)
	err = db.QueryOne(apply.Share, "select * from user where id=?", apply.ShareId)
	Check(err)
}

func (apply *Apply) Add() (id int64, err error) {
	id, err = db.Insert("apply", apply)
	return
}

func (apply *Apply) Pass() (err error) {
	apply.State = 2
	err = db.Update("apply", apply)
	return
}

func (apply *Apply) Down() (err error) {
	apply.State = 3
	err = db.Update("apply", apply)
	return
}