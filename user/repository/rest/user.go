package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spektr-account-api/domain"
)

type myUserRepository struct {
	base string
}

func (m myUserRepository) SignIn(ctx context.Context, user domain.User) (string, error) {

	url := "http://10.0.0.10:8082/system_api/?format=json&context=web&model=users&method1=web_cabinet.login&arg1={\"login\":\"" + user.Login + "\",\"passwd\":\"" + user.Password + "\"}"

	client := &http.Client{}

	// Create a new GET request with the URL
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	var res domain.User
	err = json.Unmarshal(body, &res)
	fmt.Println(res.SessionId)
	if res.SessionId == "0" {
		return "", domain.ErrUnauthorized
	}
	return res.SessionId, nil
}

func (m myUserRepository) GetBalance(ctx context.Context, SessionId string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (m myUserRepository) GetUserInfo(ctx context.Context, SessionId string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewDeliverRepository(url string) domain.UserRepository {
	return &myUserRepository{url}
}
