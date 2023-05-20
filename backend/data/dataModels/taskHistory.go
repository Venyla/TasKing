package data

import (
	"github.com/google/uuid"
	"time"
)

type TaskHistory struct {
	TaskHistoryId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk" json:"-"`
	Task          Task      `pg:"rel:has-one,fk:task_id" json:"-"`
	TaskId        uuid.UUID `pg:"type:uuid" json:"task_id"`
	CreatedBy     string    `pg:",notnull" json:"username_of_creator"`
	CreatedOn     time.Time `pg:"default:now(),notnull" json:"-"`
}
