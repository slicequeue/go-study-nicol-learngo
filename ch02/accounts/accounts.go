package accounts

import (
	"errors"
	"fmt"
)

// struct 만들기!

// Account struct
type Account struct {
	// 대문자로 하면 public 임!
	owner string
	balance int	
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0} // 여기에서는 접근이 가능함
	return &account // pointer 사용하며 복사 만들지 않고 주소 자체를 리턴
	// 생성자 역할
}

var errNoMoney = errors.New("Can't withdraw") // err~~ 형태로 에러 하는것을 추천

// Deposit x amount on your account
func (a *Account) /* 리시버라고 부름 꼭 포인터로 해야 안그러면 복사해서 새 객체임 */ Deposit(amount int) { 
	a.balance += amount
}

func (a Account) /* 이 부분은 포인터로 하지 않는다? */ Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// 예외처리 하기!
	if a.balance < amount {
		// return errors.New("Can't withdraw you are poor")
		return errNoMoney
	}
	a.balance -= amount
	return nil // 에러 없는 경우!
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}

