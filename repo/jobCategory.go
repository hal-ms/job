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
			ImageColor:  "#8069ac",
			ImageIcon:   "~assets/icon/programmer.svg",
		},
		model.JobCategory{
			Name:        "cook",
			DisplayName: "料理人",
			ImageColor:  "#e74091",
			ImageIcon:   "~assets/icon/cook.svg",
		},
		model.JobCategory{
			Name:        "carpenter",
			DisplayName: "大工",
			ImageColor:  "#92c423",
			ImageIcon:   "~assets/icon/carpenter.svg",
		},
		model.JobCategory{
			Name:        "pianist",
			DisplayName: "ピアニスト",
			ImageColor:  "#f5b628",
			ImageIcon:   "~assets/icon/pianist.svg",
		},
		model.JobCategory{
			Name:        "priest",
			DisplayName: "お坊さん",
			ImageColor:  "#40bfec",
			ImageIcon:   "~assets/icon/priest.svg",
		},
	}
}

func (j *jobCategoryRepo) All() model.JobCategorys {
	return j.all
}
