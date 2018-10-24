package model

func (j JobCategorys) FindByName(name string) *JobCategory {
	for _, value := range j {
		if value.Name == name {
			return &value
		}
	}
	return nil
}
