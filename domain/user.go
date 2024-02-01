package domain

type User struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Age   int64  `db:"age"`
}
