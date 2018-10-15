package cache

import (
	"github.com/hal-ms/job/model"
	"github.com/hal-ms/job/repo"
)

var Job = NewJobCache()

type jobCache struct {
	job model.Jobs
}

func NewJobCache() jobCache {
	res := jobCache{}
	res.Load()
	return res
}

func (j *jobCache) Load() {
	j.job = repo.Job.All()
}

func (j *jobCache) All() model.Jobs {
	return j.job
}
