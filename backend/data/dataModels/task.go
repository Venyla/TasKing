package data

import "github.com/google/uuid"

type Task struct {
	TaskId       uuid.UUID `json:"id", pg:"type:uuid,default:gen_random_uuid(),pk"`
	Title        string    `json: "title"`
	IconUrl      string    `json: "image_url"`
	XCoordinates float64   `json: "x"`
	YCoordinates float64   `json: "y"`
}
