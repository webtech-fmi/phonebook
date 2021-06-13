package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SessionInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type session struct {
	ID          string      `json:"id"`
	CreatedTime time.Time   `json:"created_time"`
	ExpiryTime  time.Time   `json:"expiry_time"`
	Payload     SessionInfo `json:"payload"`
}

type sessionValidRequest struct {
	ID string `json:"id"`
}

type sessionValidResponse struct {
	session
}

type Service interface {
	ResolveSessionID(string) (*SessionInfo, error)
	CreateUser(string, string, string) (*CreateResponse, error)
	Lock(string, string) (*LockResponse, error)
	Login(string, string) (*LoginResponse, error)
	Logout(string) error
}

type httpService struct {
	URI string
}

func (s *httpService) ResolveSessionID(sessionID string) (*SessionInfo, error) {
	requestBody, err := json.Marshal(sessionValidRequest{
		ID: sessionID,
	})

	if err != nil {
		return nil, err
	}

	httpResponse, err := http.Post(
		s.URI+"/sessions/valid",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication-service returned status [%d]", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var response sessionValidResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Payload, nil
}

type createRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type Lock struct {
	Code        string     `json:"code"`
	CreatedTime *time.Time `json:"created_time"`
	Reason      string     `json:"reason"`
}

type CreateResponse struct {
	ID   string `json:"id"`
	Lock Lock   `json:"lock"`
}

func (s *httpService) CreateUser(email, password, fullName string) (*CreateResponse, error) {
	requestBody, err := json.Marshal(createRequest{
		Email:    email,
		Password: password,
		FullName: fullName,
	})

	if err != nil {
		return nil, err
	}

	httpResponse, err := http.Post(
		s.URI+"/users",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication-service returned status [%d]", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var createResponse CreateResponse

	err = json.Unmarshal(body, &createResponse)
	if err != nil {
		return nil, err
	}

	return &createResponse, nil
}

type lockRequest struct {
	Reason string `json:"reason"`
}

type LockResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Lock  Lock   `json:"lock"`
}

func (s *httpService) Lock(id, reason string) (*LockResponse, error) {
	requestBody, err := json.Marshal(lockRequest{
		Reason: reason,
	})

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodPut,
		s.URI+"/users/lock/"+id,
		bytes.NewBuffer(requestBody),
	)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpResponse, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication-service returned status [%d]", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var lockResponse LockResponse

	err = json.Unmarshal(body, &lockResponse)
	if err != nil {
		return nil, err
	}

	return &lockResponse, nil
}

type loginRequest struct {
	Email  string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

func (s *httpService) Login(email, password string) (*LoginResponse, error) {
	requestBody, err := json.Marshal(loginRequest{
		Email:  email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	httpResponse, err := http.Post(
		s.URI+"/login",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication-service returned status [%d]", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var loginResponse LoginResponse

	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		return nil, err
	}

	return &loginResponse, nil
}

type logoutRequest struct {
	SessionID string `json:"session_id"`
}

func (s *httpService) Logout(sessionID string) error {
	requestBody, err := json.Marshal(logoutRequest{
		SessionID: sessionID,
	})

	if err != nil {
		return err
	}

	httpResponse, err := http.Post(
		s.URI+"/logout",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication-service returned status [%d]", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	return nil
}

func NewService(URI string) (Service, error) {
	return &httpService{URI: URI}, nil
}
