package marqeta

import (
	"context"
	"net/http"
)

func doSomething(usersApi UsersApiServiceInferface) {
	usersApi.GetUsers(context.Background(), nil)
}

type MockedUsersApiService struct{}

func (a *MockedUsersApiService) GetUsers(ctx context.Context, localVarOptionals *UsersApiGetUsersOpts) (UserCardHolderListResponse, *http.Response, error) {
	// mocked
	return UserCardHolderListResponse{}, nil, nil
}

func main() {
	actualApiService := &UsersApiService{}
	doSomething(actualApiService)
	mockedApiService := &MockedUsersApiService{}
	doSomething(mockedApiService)
}
