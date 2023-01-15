package db

import (
	"fmt"

	"github.com/mager/penny-pincher/entity"
)

func GetBudgetQuery(bID string) string {
	return fmt.Sprintf("select * from budgets where id = '%s'", bID)
}

func GetBudgetTransactionsQuery(bID string) string {
	return fmt.Sprintf("select * from trxs where budget_id = '%s'", bID)
}

func GetUserQuery(auth0 string) string {
	return fmt.Sprintf("select * from users where auth0 = '%s'", auth0)
}

func CreateUserQuery(u entity.CreateUserReq) string {
	return fmt.Sprintf(
		"insert into users (auth0, email, phone, locale) values"+
			" ('%s', '%s', '%s', '%s');",
		u.Auth0, u.Email, u.Phone, u.Locale,
	)
}
