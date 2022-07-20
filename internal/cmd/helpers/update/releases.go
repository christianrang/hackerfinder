package update

import (
	"errors"
	"fmt"
	"runtime"

	resty "github.com/go-resty/resty/v2"
)

type Response []Release

type Release struct {
	Url             string   `json:"url"`
	AssetsUrl       string   `json:"assets_url"`
	UploadUrl       string   `json:"upload_url"`
	HtmlUrl         string   `json:"html_url"`
	Id              int      `json:"id"`
	Author          User     `json:"author"`
	NodeId          string   `json:"node_id"`
	TagName         string   `json:"tag_name"`
	TargetCommitish string   `json:"target_commitish"`
	Name            string   `json:"name"`
	Draft           bool     `json:"draft"`
	Prerelease      bool     `json:"prerelease"`
	CreatedAt       string   `json:"created_at"`
	PublishedAt     string   `json:"published_at"`
	Assets          []Assets `json:"assets"`
	TarballUrl      string   `json:"tarball_url"`
	ZipballUrl      string   `json:"zipball_url"`
	MentionsCount   int      `json:"mentions_count"`
}

type User struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubsciptionsUrl   string `json:"subsciptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Assets struct {
	Url                string `json:"url"`
	Id                 int    `json:"id"`
	NodeId             string `json:"node_id"`
	Name               string `json:"name"`
	Label              string `json:"label"`
	Uploader           User   `json:"uploader"`
	ContentType        string `json:"content_type"`
	State              string `json:"state"`
	Size               int    `json:"size"`
	DownloadCount      int    `json:"download_count"`
	CreatedAt          string `json:"created_at"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

func getReleases() (Response, error) {
	var (
		response Response
		err      error
	)
	_, err = resty.New().R().
		SetResult(&response).
		Get("https://api.github.com/repos/christianrang/hackerfinder/releases")
	if err != nil {
		return response, errors.New(fmt.Sprintf("error in sending request %s: ", err))
	}

	return response, nil
}

func FindCurrentReleaseUrl() (string, error) {
	var releaseUrl string

	releases, err := getReleases()
	if err != nil {
		return releaseUrl, err
	}

	for _, asset := range releases[0].Assets {
		if asset.Name == fmt.Sprintf("hackerfinder_%s_%s", runtime.GOOS, runtime.GOARCH) {
			releaseUrl = asset.BrowserDownloadUrl
			break
		}
	}

	return releaseUrl, nil
}
