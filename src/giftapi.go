package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

var (
	api = "https://gift-api.stepcdn.space"
	h   = map[string]string{
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36",
		"Host":         "gift-api.stepcdn.space",
		"Content-Type": "application/json",
	}
)

func req(m, u string, b []byte) (any, error) {
	r, _ := http.NewRequest(m, u, bytes.NewBuffer(b))
	for k, v := range h {
		r.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var res any
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func Auth(t string) (any, error) {
	h["Authorization"] = t
	res, err := req("POST", api+"/auth/new", []byte("{}"))
	if m, ok := res.(map[string]any); ok {
		if token, ok := m["access_token"].(string); ok {
			h["Authorization"] = "Bearer " + token
		}
	}
	return res, err
}

func GetInventory(include string) (any, error) {
	return req("GET", api+"/inventory?include="+include, nil)
}

func GetWrapQuests(tag string) (any, error) {
	return req("GET", api+"/wrapquests?tag="+tag, nil)
}

func GetInventoryResources() (any, error) {
	return req("GET", api+"/inventory/resources", nil)
}

func CellsMerge(ids []int) (any, error) {
	data, _ := json.Marshal(map[string]any{"cell_ids": ids})
	return req("POST", api+"/game2048/cells/merge", data)
}

func BurnMerge(id int) (any, error) {
	return req("POST", api+"/game2048/cells/"+strconv.Itoa(id)+"/burn", []byte("{}"))
}

func PostSpawn() (any, error) {
	return req("POST", api+"/game2048/1/spawn", []byte("{}"))
}
