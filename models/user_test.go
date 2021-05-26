package models

import "testing"

func TestGetUserDao(t *testing.T) {
	dao := GetUserDao()
	t.Log(dao)
}
