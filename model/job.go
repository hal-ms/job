package model

func (j Jobs) FindById(id string) *Job {
	for _, value := range j {
		if value.ID.Hex() == id {
			return &value
		}
	}
	return nil
}

func (j Jobs) FindByIds(ids []string) Jobs {
	res := make(Jobs, 0)
	for _, value := range ids {
		job := j.FindById(value)
		if job != nil {
			res = append(res, *job)
		}
	}
	return res
}
