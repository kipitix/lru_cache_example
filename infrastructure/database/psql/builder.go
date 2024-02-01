package psql

import (
	"lrucache/domain"

	sq "github.com/Masterminds/squirrel"
)

const userTable = "users"

func prepareInsertOrUpdateUser(user domain.User) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Insert(userTable).
		Columns("email, name, age").
		Values(user.Email, user.Name, user.Age).
		Suffix(`
			ON CONFLICT (email) DO UPDATE SET
				name = EXCLUDED.name,
				age  = EXCLUDED.age
		`)

	return rawQuery.ToSql()
}

func prepareFindUserByEmailQuery(email string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return psqlSq.Select("email, name, age").
		From(userTable).
		Where(sq.Eq{"email": email}).
		ToSql()
}
