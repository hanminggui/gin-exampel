package moudles

import (
	db "../database"
)

type Person struct {
	Id        int    `json:"id"`
	Nickname  string `json:"nickname" form:"nickname"`
	Password  string `json:"password" form:"password"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO user(nickname, password, first_name, last_name) VALUES (?, ?)",
		p.Nickname, p.Password, p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetDetail() {
	row := db.SqlDB.QueryRow("select id, nickname, password, first_name, last_name from user where id=?", p.Id)
	err := row.Scan(&p.Id, &p.Nickname, &p.Password, &p.FirstName, &p.LastName)
	if err != nil {
		panic(err)
	}
}
