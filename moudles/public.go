package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
	"reflect"
)
type CuAt struct {
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

func init()  {
	structs := [...] interface{} {new(User), new(Share), new(Coach), new(Attention), new(Apply)}
	for _,s:= range structs{
		db.ModMap[reflect.TypeOf(s).Name()] = s
		db.ModMap["moudles." + reflect.TypeOf(s).Name()] = s
		db.ModMap["[]moudles." + reflect.TypeOf(s).Name()] = s
	}
}