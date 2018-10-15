package repo

import "github.com/hal-ms/job/model"

var JobCategory jobCategoryRepo

type jobCategoryRepo struct {
	All model.JobCategorys
}

func init() {

}
