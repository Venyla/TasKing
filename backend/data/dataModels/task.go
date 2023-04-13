package data

import "github.com/google/uuid"

type Task struct {
	TaskId       uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	IconUrl      string    `json:"image_url"`
	XCoordinates float64   `json:"x"`
	YCoordinates float64   `json:"y"`
}
