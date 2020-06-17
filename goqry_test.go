package goqry

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	tbl := NewTable("users")
	tbl.Select("first_name", "last_name")
	fmt.Println(tbl.Indent("\n", ""))
}

func TestJoin(t *testing.T) {
	users := NewTable("users").Select("first_name", "last_name")
	assets := NewTable("user_assets").Select("name")
	users.InnerJoin(assets).Field("user_id").EqField(assets.Field("user_id"))
	fmt.Println(users.Indent("\n", ""))
}

func TestWhere(t *testing.T) {
	ut := NewTable("users").Select("first_name", "last_name")
	ut.Where(ut.Field("user_id").MakeNum().Eq("1"))
	fmt.Println(ut.Indent("\n", ""))
}
