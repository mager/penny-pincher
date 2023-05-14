package db

import (
	"fmt"
)

func GetBudgetsQuery(userID string) string {
	return fmt.Sprintf(
		"select id, userid, name, amount, year, month from budgets "+
			"where userid = '%s' order by createdat desc", userID)
}

func GetBudgetQuery(uID string) string {
	return fmt.Sprintf("select * from budgets where id = '%s'", uID)
}

func GetBudgetTransactionsQuery(bID string) string {
	return fmt.Sprintf("select * from trxs where budgetid = '%s'", bID)
}

func GetUserQuery(email string) string {
	return fmt.Sprintf("select * from users where email = '%s'", email)
}

func CreateUserQuery(userID, email, phone, country string) string {
	return fmt.Sprintf(
		"insert into users (userid, email, phone, country) values"+
			" ('%s', '%s', '%s', '%s');",
		userID, email, phone, country,
	)
}
