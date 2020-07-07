package github

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

//TODO: Add auth
//TODO: Handle not found repos
func GetCommits(owner, repoName string) ([]Commit, error) {
	client := resty.New()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repoName)
	res, err := client.R().
		EnableTrace().
		SetResult([]Commit{}).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Get commits has failed")
	}
	githubResponse, ok := res.Result().(*[]Commit)

	if !ok {
		err := fmt.Errorf("invalid GitHub response")
		return nil, errors.Wrap(err, "GitHub request failed")
	}
	var commits []Commit
	commits = append(commits, *githubResponse...)
	return commits, nil
}

func GetRepository(owner, repoName string) (*Repository, error) {
	client := resty.New()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)
	res, err := client.R().
		EnableTrace().
		SetResult(&Repository{}).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Get repository has failed")
	}
	githubResponse, ok := res.Result().(*Repository)

	if !ok {
		err := fmt.Errorf("invalid GitHub response")
		return nil, errors.Wrap(err, "GitHub request failed")
	}

	return githubResponse, nil
}

func GetHead(owner, repoName string) (*Head, error) {
	client := resty.New()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/HEAD", owner, repoName)
	res, err := client.R().
		EnableTrace().
		SetResult(&Head{}).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Get repository has failed")
	}
	githubResponse, ok := res.Result().(*Head)

	if !ok {
		err := fmt.Errorf("invalid GitHub response")
		return nil, errors.Wrap(err, "GitHub request failed")
	}

	return githubResponse, nil
}
