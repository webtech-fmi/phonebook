package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type Profile struct {
	ID uuid.UUID `json:"id"`
}

type Service interface {
	ResolveUserID(string) (*Profile, error)
}

type httpService struct {
	URI string
}

func (s *httpService) ResolveUserID(userID string) (*Profile, error) {
	httpResponse, err := http.Get(
		fmt.Sprintf("%s/profiles/by-owner?id=%s", s.URI, userID),
	)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("profile-service returned status [%d]", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var response Profile
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func NewService(URI string) (Service, error) {
	return &httpService{URI: URI}, nil
}
