package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIResponseSch - JSON response with camera-schedule data from UniFi-Video Controller
type APIResponseSch struct {
	Data []ScheduleEntity `json:"data"`

	Meta struct {
		TotalCount    int `json:"totalCount"`
		FilteredCount int `json:"filteredCount"`
	} `json:"meta"`
}

type ScheduleEntity struct {
	DaySchedules []struct {
		DayOfWeek     int `json:"dayOfWeek"`
		ScheduleItems []struct {
			StartHour   int    `json:"startHour"`
			StartMinute int    `json:"startMinute"`
			EndHour     int    `json:"endHour"`
			EndMinute   int    `json:"endMinute"`
			Action      string `json:"action"`
		} `json:"scheduleItems"`
	} `json:"daySchedules"`
	Name string `json:"name"`
	ID   string `json:"_id"`
}

// APIResource - return resource name in API path
func (APIResponseSch) APIResource() string {
	return "cameraschedule"
}

// Get - make API call and fill Response
func (a *APIResponseSch) Get() error {
	URL := getURL(a)

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	return nil
}

// GetEntity - gets schedule entity with given id
func (a *APIResponseSch) GetEntity(id string) *ScheduleEntity {
	for _, d := range a.Data {
		if d.ID == id {
			return &d
		}
	}
	return nil
}
