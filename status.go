package main

import (
	"encoding/json"
	"io/ioutil"
)

type StatusEntry struct {
	Message string `json:"message"`
	Emoji   string `json:"emoji"`
}

func ReadStatusEntriesFromFile() ([]StatusEntry, error) {
	filename := "statuses.json"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var entries []StatusEntry
	if err := json.Unmarshal(content, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}
