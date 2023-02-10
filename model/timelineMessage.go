package model

type TimelineMessage struct {
	Username  string `db:"username"`
	UserId    int    `db:"user_id"`
	Email     string `db:"email"`
	MessageId int    `db:"message_id"`
	AuthorId  int    `db:"author_id"`
	Text      string `db:"text"`
	PubDate   int64  `db:"pub_date"`
	Flagged   int    `db:"flagged"`
	PwHash    string `db:"pw_hash"`
}
