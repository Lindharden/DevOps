package simModels

type FilteredMessageRequest struct {
	Text     string `db:"text"`
	PubDate  int64  `db:"pub_date"`
	Username string `db:"username"`
}
