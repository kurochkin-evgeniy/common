package tmp_session_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"tmp_session_contract"
)

const BaseUrl = "http://localhost:5008/api/sessions/"

type Client struct {
}

func (client *Client) CreateSession(id, auxKey, data string) (string, error) {

	request := tmp_session_contract.CreationRequest{
		AuxKey:      auxKey,
		SessionData: data,
	}

	json_data, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	response, err := http.Post(BaseUrl+id, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(response.StatusCode))
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseObject tmp_session_contract.CreationResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return "", err
	}

	return responseObject.UserCode, nil
}

func (client *Client) GetSessionData(id, auxKey, userCode string) (string, error) {
	url := BaseUrl + id + fmt.Sprintf("?aux_key=%s&user_code=%s", auxKey, userCode)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(response.StatusCode))
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseObject tmp_session_contract.GetResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return "", err
	}

	return responseObject.SessionData, nil
}

func (client *Client) RetriveSession(id, auxKey, userCode string) (string, error) {
	url := BaseUrl + id + "/retrive" + fmt.Sprintf("?aux_key=%s&user_code=%s", auxKey, userCode)
	response, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(response.StatusCode))
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseObject tmp_session_contract.GetResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return "", err
	}

	return responseObject.SessionData, nil
}
