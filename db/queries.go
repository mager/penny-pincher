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

func GetUserQuery(auth0ID string) string {
	return fmt.Sprintf("select * from users where auth0_id = '%s'", auth0ID)
}

func CreateUserQuery(u entity.CreateUserReq) string {
	return fmt.Sprintf(
		"insert into users (auth0_id, email, phone, locale) values"+
			" ('%s', '%s', '%s', '%s');",
		u.Auth0ID, u.Email, u.Phone, u.Locale,
	)
}
