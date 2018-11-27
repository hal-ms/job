package model

var pds = []ProvidingDestination{
	ProvidingDestination{
		PID:   1,
		PName: "サーバー",
		PRate: 1000,
	},
	ProvidingDestination{
		PID:   2,
		PName: "ストレージ",
		PRate: 1200,
	},
	ProvidingDestination{
		PID:   3,
		PName: "AI",
		PRate: 1600,
	},
	ProvidingDestination{
		PID:   4,
		PName: "演算",
		PRate: 700,
	},
}

func (u *User) PDSet(id int) {
	for _, val := range pds {
		if val.PID == id {
			u.ProvidingDestination = val
		}
	}
}
