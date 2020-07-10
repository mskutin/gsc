package github

import (
	"fmt"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
)

var errGetHead = errors.New("Get head request has failed")
var errGetRepository = errors.New("Get repository request has failed")
var errGithubTokenIsAbsent = errors.New("Github Token is required")
var errGithubUsernameIsAbsent = errors.New("Github username is required")

var errUnauthorized = errors.New("Unauthorized")
var ErrNotFound = errors.New("Requested repository was not found")

type Client struct {
	RestyClient *resty.Client
}

func New() (*Client, error) {
	baseURL := "https://api.github.com"
	client := resty.New()
	client.SetHostURL(baseURL)
	return &Client{RestyClient: client}, nil
}

func NewWithAuth(username, token string) (*Client, error) {
	baseURL := "https://api.github.com"
	err := validation.Validate(username, validation.Required)
	if err != nil {
		return nil, errGithubUsernameIsAbsent
	}
	err = validation.Validate(token, validation.Required)
	if err != nil {
		return nil, errGithubTokenIsAbsent
	}
	client := resty.New()
	client.SetHostURL(baseURL)
	client.SetBasicAuth(username, token)
	return &Client{RestyClient: client}, nil
}

func (c *Client) IsTokenValid() bool {
	res, err := c.RestyClient.R().Get("/user")

	if err != nil {
		return false
	}
	if res.StatusCode() == 401 {
		return false
	}
	return true
}

func (c *Client) GetHead(owner, repoName string) (*Head, error) {
	res, err := c.RestyClient.R().
		SetResult(&Head{}).
		Get(fmt.Sprintf("/repos/%s/%s/commits/HEAD", owner, repoName))
	if err != nil {
		return nil, errGetHead
	}
	if res.StatusCode() == 401 {
		log.Println("Status code: ", res.StatusCode())
		return nil, errUnauthorized
	}

	if res.StatusCode() == 404 {
		return nil, ErrNotFound
	}
	return res.Result().(*Head), nil
}

func (c *Client) GetRepository(owner, repoName string) (*Repository, error) {
	res, err := c.RestyClient.R().
		SetResult(&Repository{}).
		Get(fmt.Sprintf("/repos/%s/%s", owner, repoName))
	if err != nil {
		return nil, errGetRepository
	}
	if res.StatusCode() == 401 {
		return nil, errUnauthorized
	}

	if res.StatusCode() == 404 {
		return nil, ErrNotFound
	}
	return res.Result().(*Repository), nil
}
