package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
	"os/user"
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
	CuAt
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
	rows, err := db.SqlDB.Query("SELECT * from share s, apply p where p.share_id=?", share.Id)
	if err != nil {
		panic(err)
	}
	share.Applys = make([]*Apply, 0)
	for rows.Next() {
		apply := Apply{}
		rows.Scan(apply.Id, apply.State, apply.ApplyType, apply.user.Id, apply.user.NickName, apply.user.Sex, apply.user.Stage, apply.user.State,
			apply.user.UserType, apply.user.Birthday, apply.user.SchoolName, apply.user.BriefInfo, apply.user.Company, apply.user.Position,
			apply.user.Specialty, apply.user.HeadImgUrl)
		share.Applys = append(share.Applys, &apply)
	}
}
