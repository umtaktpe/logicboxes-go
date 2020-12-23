package logicboxes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	BaseURL    string
	AuthUserID string
	APIKey     string
	HTTPClient *http.Client
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient(authUserID, apiKey string) *Client {
	return &Client{
		BaseURL:    "https://httpapi.com/api/",
		AuthUserID: authUserID,
		APIKey:     apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) Request(req *http.Request, response interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		errResponse := ErrorResponse{}
		if err = json.NewDecoder(res.Body).Decode(&errResponse); err == nil {
			return errors.New(errResponse.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

func (c *Client) PostRequest(requestUrl string, response interface{}, query map[string]interface{}) error {
	param := url.Values{}

	param.Add("auth-userid", c.AuthUserID)
	param.Add("api-key", c.APIKey)

	for k, v := range query {
		rt := reflect.TypeOf(v)

		switch rt.Kind() {
		case reflect.Int:
			param.Add(k, strconv.Itoa(v.(int)))
		case reflect.String:
			param.Add(k, v.(string))
		case reflect.Bool:
			param.Add(k, strconv.FormatBool(v.(bool)))
		case reflect.Float64:
			amount := v.(float64)
			param.Add(k, fmt.Sprintf("%f", amount))
		case reflect.Slice:
			if rt.Elem().String() != "string" {
				return errors.New("array elements should be string")
			}
			for _, value := range v.([]string) {
				param.Add(k, value)
			}
		default:
			return errors.New("invalid data type")
		}
	}

	req, err := http.NewRequest("POST", c.BaseURL+requestUrl, strings.NewReader(param.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err := c.Request(req, &response); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetRequest(requestUrl string, response interface{}, query map[string]interface{}) error {
	req, err := http.NewRequest("GET", c.BaseURL+requestUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	param := req.URL.Query()
	param.Add("auth-userid", c.AuthUserID)
	param.Add("api-key", c.APIKey)

	for k, v := range query {
		rt := reflect.TypeOf(v)

		switch rt.Kind() {
		case reflect.Int:
			param.Add(k, strconv.Itoa(v.(int)))
		case reflect.String:
			param.Add(k, v.(string))
		case reflect.Bool:
			param.Add(k, strconv.FormatBool(v.(bool)))
		case reflect.Float64:
			amount := v.(float64)
			param.Add(k, fmt.Sprintf("%f", amount))
		case reflect.Slice:
			if rt.Elem().String() != "string" {
				return errors.New("array elements should be string")
			}
			for _, value := range v.([]string) {
				param.Add(k, value)
			}
		default:
			return errors.New("invalid data type")
		}
	}
	req.URL.RawQuery = param.Encode()

	if err := c.Request(req, &response); err != nil {
		return err
	}

	return nil
}
