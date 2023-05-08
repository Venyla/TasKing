package data

import "github.com/google/uuid"

type Task struct {
	TaskId        uuid.UUID     `pg:"type:uuid,default:gen_random_uuid(),pk" json:"id"`
	Title         string        `pg:",notnull" json:"title"`
	IconUrl       string        `pg:",notnull" json:"image_url"`
	XCoordinates  float64       `pg:",notnull" json:"x"`
	YCoordinates  float64       `pg:",notnull" json:"y"`
	TaskHistories []TaskHistory `pg:"rel:has-many,join_fk:task_id" json:"-"`
}
