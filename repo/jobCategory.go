package repo

import "github.com/hal-ms/job/model"

var JobCategory jobCategoryRepo

type jobCategoryRepo struct {
	all model.JobCategorys
}

var publicBase = "https://hal-iot.net/public"

func init() {
	JobCategory = jobCategoryRepo{}
	JobCategory.all = model.JobCategorys{
		model.JobCategory{
			Name:        "programmer",
			DisplayName: "プログラマー",
			ImageColor:  "#8069ac",
			ImageIcon:   publicBase + "/icon/programmer.svg",
		},
		model.JobCategory{
			Name:        "cook",
			DisplayName: "料理人",
			ImageColor:  "#e74091",
			ImageIcon:   publicBase + "/icon/cook.svg",
		},
		model.JobCategory{
			Name:        "carpenter",
			DisplayName: "大工",
			ImageColor:  "#92c423",
			ImageIcon:   publicBase + "/icon/carpenter.svg",
		},
		model.JobCategory{
			Name:        "pianist",
			DisplayName: "ピアニスト",
			ImageColor:  "#f5b628",
			ImageIcon:   publicBase + "/icon/pianist.svg",
		},
		model.JobCategory{
			Name:        "priest",
			DisplayName: "お坊さん",
			ImageColor:  "#40bfec",
			ImageIcon:   publicBase + "/icon/priest.svg",
		},
	}
}

func (j *jobCategoryRepo) All() model.JobCategorys {
	return j.all
}
