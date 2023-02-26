package simModels

type FilteredMessageRequest struct {
	Text     string `db:"content" json:"content"`
	PubDate  int64  `db:"pub_date" json:"pub_date"`
	Username string `db:"username" json:"user"`
}
