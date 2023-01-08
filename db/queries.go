package db

import "fmt"

func GetBudgetQuery(bID string) string {
	return fmt.Sprintf("SELECT * FROM budgets WHERE id = '%s'", bID)
}

func GetBudgetTransactionsQuery(bID string) string {
	return fmt.Sprintf("SELECT * FROM trxs WHERE budget_id = '%s'", bID)
}

func GetUserQuery(auth0ID string) string {
	return fmt.Sprintf("SELECT * FROM users WHERE auth0_id = '%s'", auth0ID)
}
