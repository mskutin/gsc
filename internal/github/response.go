package github

import "time"

type Parent struct {
	Sha     string `json:"sha"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type Verification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

type Tree struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

type Author struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Commiter struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Commit struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message      string       `json:"message"`
		Tree         Tree         `json:"tree"`
		URL          string       `json:"url"`
		CommentCount int          `json:"comment_count"`
		Verification Verification `json:"verification"`
	} `json:"commit"`
	URL         string   `json:"url"`
	HTMLURL     string   `json:"html_url"`
	CommentsURL string   `json:"comments_url"`
	Author      Author   `json:"author"`
	Committer   Commiter `json:"committer"`
	Parents     []Parent `json:"parents"`
}
