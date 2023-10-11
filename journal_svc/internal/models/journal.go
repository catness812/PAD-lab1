package models

type JournalEntry struct {
	Username string `bson:"username"`
	Title    string `bson:"title"`
	Content  string `bson:"content"`
}
