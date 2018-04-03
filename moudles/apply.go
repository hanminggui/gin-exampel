package moudles

type Apply struct {
	Id        int64  `json:"id"`
	ApplyType int    `json:"apply_type"`
	State     int    `json:"state"`
	UserId    int64  `json:"user_id"`
	Share_id  int64  `json:"share_id"`
	User      *User  `json:"user"`
	Share     *Share `json:"share"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}
