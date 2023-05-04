package data

import (
	"github.com/google/uuid"
	"time"
)

type TaskHistory struct {
	TaskHistoryId uuid.UUID `json:"id"`
	TaskId        uuid.UUID `json:"task_id"`
	CreatedBy     string    `json:"username_of_creator"`
	CreatedOn     time.Time
}
