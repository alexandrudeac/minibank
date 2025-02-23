package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/alexandrudeac/minibank/util"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)
	gotAcc, err := testStore.GetAccount(context.Background(), createdAcc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAcc)
	require.Equal(t, createdAcc.ID, gotAcc.ID)
	require.Equal(t, createdAcc.Balance, gotAcc.Balance)
	require.Equal(t, createdAcc.Owner, gotAcc.Owner)
	require.WithinDuration(t, createdAcc.CreatedAt, gotAcc.CreatedAt, 1, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testStore.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testStore.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    createRandomUser(t).Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testStore.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return acc
}
