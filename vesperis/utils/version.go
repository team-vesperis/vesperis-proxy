package utils

import (
	"encoding/json"
	"net/http"
)

type gitHubRelease struct {
	TagName string `json:"tag_name"`
}

var vesperisProxyVersion string

func initializeVersions() {
	initializeVesperisProxyVersion()
}

func initializeVesperisProxyVersion() {
	resp, err := http.Get("https://api.github.com/repos/team-vesperis/vesperis-proxy/releases/latest")
	if err != nil {
		vesperisProxyVersion = "error"
		return
	}
	defer resp.Body.Close()

	var release gitHubRelease
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		vesperisProxyVersion = "error"
		return
	}

	vesperisProxyVersion = release.TagName
}

func GetVesperisProxyVersion() string {
	return vesperisProxyVersion
}
