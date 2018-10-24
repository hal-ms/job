package repo

import "github.com/hal-ms/job/model"

var JobCategory jobCategoryRepo

type jobCategoryRepo struct {
	all model.JobCategorys
}

func init() {
	JobCategory = jobCategoryRepo{}
	JobCategory.all = model.JobCategorys{
		model.JobCategory{
			Name:        "programmer",
			DisplayName: "プログラマー",
		},
		model.JobCategory{
			Name:        "cook",
			DisplayName: "料理人",
		},
		model.JobCategory{
			Name:        "carpenter",
			DisplayName: "大工",
		},
		model.JobCategory{
			Name:        "judge",
			DisplayName: "裁判官",
		},
		model.JobCategory{
			Name:        "pianist",
			DisplayName: "ピアニスト",
		},
		model.JobCategory{
			Name:        "priest",
			DisplayName: "僧侶",
		},
	}
}

func (j *jobCategoryRepo) All() model.JobCategorys {
	return j.all
}
