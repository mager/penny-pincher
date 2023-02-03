package db

import (
	"fmt"

	"github.com/mager/penny-pincher/entity"
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

func CreateUserQuery(u entity.CreateUserReq, userID string) string {
	return fmt.Sprintf(
		"insert into users (userid, email, phone, locale) values"+
			" ('%s', '%s', '%s', '%s');",
		userID, u.Email, u.Phone, u.Locale,
	)
}
