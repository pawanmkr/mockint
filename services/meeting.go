package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Meeting struct {
	JoinUrl       string `json:"joinUrl"`
	MeetingCode   string `json:"meetingCode"`
	Subject       string `json:"subject"`
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
}

func CreateMeeting(postBody []byte) (*Meeting, error) {
	url := "https://graph.microsoft.com/v1.0/users/367533eb-6ce5-49cb-9f35-6a0518426b0f/onlineMeetings"
	bearerToken := os.Getenv("BEARER_TOKEN")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
		return nil, err
	}
	// fmt.Println(resp.StatusCode) // do check this
	defer resp.Body.Close()

	var res Meeting

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(responseData, &res); err != nil {
		log.Fatal(err)
	}
	return &res, nil
}
