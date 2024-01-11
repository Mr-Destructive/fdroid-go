package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type packageInfo struct {
    PackageName string `json:"packageName"`
    SuggestedVersionCode int `json:"suggestedVersionCode"`
    Packages []struct {
        VersionName string `json:"versionName"`
        VersionCode int `json:"versionCode"`
    } `json:"packages"`
}

func GetPackageInfo(name string) (packageInfo, error) {
    resp, err := http.Get(fmt.Sprintf("https://f-droid.org/api/v1/packages/%s", name))
    if err != nil || resp.StatusCode != 200 {
        return packageInfo{}, fmt.Errorf("Error fetching package info: %d", resp.StatusCode)
    }
    var pkgInfo packageInfo
    err = json.NewDecoder(resp.Body).Decode(&pkgInfo)
    if err != nil {
        return packageInfo{}, fmt.Errorf("Error decoding package info: %v", err)
    }
    return pkgInfo, nil
}

func SearchApps(query string) (interface{}, error) {
    resp, err := http.Get(fmt.Sprintf("https://search.f-droid.org/api/search_apps?q=%s", query))
    if err != nil || resp.StatusCode != 200 {
        return packageInfo{}, fmt.Errorf("Error fetching package info: %d", resp.StatusCode)
    }
    var pkgInfo interface{}
    err = json.NewDecoder(resp.Body).Decode(&pkgInfo)
    fmt.Println(pkgInfo)
    if err != nil {
        return packageInfo{}, fmt.Errorf("Error decoding package info: %v", err)
    }
    return pkgInfo, nil
}
