package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "go-simple-bank/db/mock"
	db "go-simple-bank/db/sqlc"
	"go-simple-bank/db/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	account := randomAccount()

	ctlr := gomock.NewController(t)
	defer ctlr.Finish()

	store := mockdb.NewMockStore(ctlr)
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)

	// check response

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Account {
	return db.Account{
		ID:    util.RandomInt(1, 100),
		Owner: util.RandomOwner(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotAccount db.Account
	json.Unmarshal(data, &gotAccount)
	require.Equal(t, gotAccount, account)

}
