package accounts

// Account struck
type Account struct {
	owner   string
	balance int
}

//NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
