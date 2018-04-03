package moudles

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
