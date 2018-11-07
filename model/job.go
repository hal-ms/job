package model

import "sort"

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

func (j Jobs) IsDone(done bool) Jobs {
	res := make(Jobs, 0)
	for _, value := range j {
		if value.Done == done {
			res = append(res, value)
		}
	}
	return res
}

func (p Jobs) SortByCreated() Jobs {
	var res = make(Jobs, 0)
	sort.Slice(p, func(i, j int) bool {
		return p[i].Created.Unix() > p[j].Created.Unix()
	})
	for _, value := range p {
		res = append(res, value)
	}
	return res
}
