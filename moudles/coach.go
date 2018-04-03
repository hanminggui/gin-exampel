package moudles

type Coach struct {
	Id        int64  `json:"id"`
	Level     int    `json:"level"`
	Type      int    `json:"type"`
	PicsUrl   string `json:"pics_url"`
	State     int    `json:"state"`
	CheckDesc string `json:"check_desc"`
	Specialty string `json:"specialty"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}
