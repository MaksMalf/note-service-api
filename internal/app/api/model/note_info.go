package model

type NoteInfo struct {
	Id     int64  `db:"id"`
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
}
