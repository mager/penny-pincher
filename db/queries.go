package db

import (
	"fmt"

	"github.com/mager/penny-pincher/entity"
)

func GetBudgetQuery(bID string) string {
	return fmt.Sprintf("select * from budgets where id = '%s'", bID)
}

func GetBudgetsQuery(uID string) string {
	return fmt.Sprintf("select * from budgets where userID = '%s'", uID)
}

func GetBudgetTransactionsQuery(bID string) string {
	return fmt.Sprintf("select * from trxs where budgetID = '%s'", bID)
}

func GetUserQuery(userID string) string {
	return fmt.Sprintf("select * from users where userID = '%s'", userID)
}

func CreateUserQuery(u entity.CreateUserReq) string {
	return fmt.Sprintf(
		"insert into users (userID, email, phone, locale) values"+
			" ('%s', '%s', '%s', '%s');",
		u.UserID, u.Email, u.Phone, u.Locale,
	)
}
