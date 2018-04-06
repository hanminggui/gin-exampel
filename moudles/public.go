package moudles

import (
	db "github.com/hanminggui/gin-exampel/database"
	"reflect"
	"log"
)

func init()  {
	structs := [...] interface{} {new(User), new(Share), new(Coach), new(Attention), new(Apply)}
	for _,s:= range structs{
		db.ModMap[reflect.TypeOf(s).Name()] = s
		db.ModMap["moudles." + reflect.TypeOf(s).Name()] = s
		db.ModMap["[]moudles." + reflect.TypeOf(s).Name()] = s
	}
}

func Check(err error)  {
	if err != nil {
		log.Panicln(err.Error())
	}
}