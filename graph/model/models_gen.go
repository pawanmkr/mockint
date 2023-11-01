// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BookInterview struct {
	InterviewID string `json:"interviewId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type DeleteResponse struct {
	DeleteInterviewID string `json:"deleteInterviewId"`
}

type Interview struct {
	ID          string      `json:"id" bson:"_id"`
	Duration    int         `json:"duration"`
	Time        string      `json:"time"`
	Name        string      `json:"name"`
	Skills      string      `json:"skills"`
	Difficulty  string      `json:"difficulty"`
	GuestType   string      `json:"guestType"`
	Guest       []*TempUser `json:"guest"`
	Note        string      `json:"note"`
	Booked      bool        `json:"booked"`
	JoinURL     string      `json:"joinUrl"`
	MeetingCode string      `json:"meetingCode"`
}

type InterviewInput struct {
	Duration    int              `json:"duration"`
	Time        string           `json:"time"`
	Name        string           `json:"name"`
	Skills      string           `json:"skills"`
	Difficulty  string           `json:"difficulty"`
	GuestType   string           `json:"guestType"`
	Guest       []*TempUserInput `json:"guest"`
	Note        string           `json:"note"`
	Booked      bool             `json:"booked"`
	JoinURL     string           `json:"joinUrl"`
	MeetingCode string           `json:"meetingCode"`
}

type TempUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TempUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
