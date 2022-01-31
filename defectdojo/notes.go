package defectdojo

import "time"

type NotesService struct {
	client *Client
}

type Note struct {
	Id      *int  `json:"id,omitempty"`
	Author  *User `json:"author,omitempty"`
	Editor  *User `json:"editor,omitempty"`
	History []*struct {
		Id            *int `json:"id,omitempty"`
		CurrentEditor *struct {
			Id        *int    `json:"id,omitempty"`
			Username  *string `json:"username,omitempty"`
			FirstName *string `json:"first_name,omitempty"`
			LastName  *string `json:"last_name,omitempty"`
		} `json:"current_editor,omitempty"`
		Data     *string    `json:"data,omitempty"`
		Time     *time.Time `json:"time,omitempty"`
		NoteType *int       `json:"note_type,omitempty"`
	} `json:"history,omitempty"`
	Entry    *string    `json:"entry,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	Private  bool       `json:"private,omitempty"`
	Edited   bool       `json:"edited,omitempty"`
	EditTime *time.Time `json:"edit_time,omitempty"`
	NoteType *int       `json:"note_type,omitempty"`
}

type Notes *struct {
	Count    *int    `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []*Note `json:"results,omitempty"`
}
