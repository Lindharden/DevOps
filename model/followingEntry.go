package model

type FollowingEntry struct {
	WhoId  int `db:"who_id"`
	WhomId int `db:"whom_id"`
}
