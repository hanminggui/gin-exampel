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
	row := db.SqlDB.QueryRow("SELECT * from share where id=?", share.Id)
	err := row.Scan(share.Title, share.StartAt, share.EndAt, share.Amount, share.Type, share.AuditState, share.LookesState, share.IsDelete)
	if err != nil {
		panic(err)
	}
}

/**
 * 获取报名列表
 */
func (share *Share) GetApplys() {
}
