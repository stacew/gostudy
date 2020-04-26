package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct { // 소문자인 경우 외부로 보내면 경고뜸.
	owner   string // 구조체 변수 외부 접근 대문자 public
	balance int
}

var errNoMoney = errors.New("Can't withDraw you are poor")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	retA := Account{owner: owner, balance: 0}
	//Go의 컴파일러에는 escape analysis가 있어 지역 변수의 주소가 함수에서 반환될 때 (분석을 통해 오류가 발생하지 않도록 하기 위해) heap으로 escape 됩니다.
	return &retA
}

//Deposit Reciver함수의 규칙 struct의 첫글자 소문자 사용하자.
//Call by Reference
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance ~~
//Call by Value
func (a Account) Balance() int {
	return a.balance
}

//WithDraw ~~
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

//ChangeOwner of the Account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner ~~
func (a Account) Owner() string {
	return a.owner
}
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s Account.\nHas: ", a.Balance())
}
