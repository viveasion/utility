package genealogy

import "time"

type Member struct {
	Id          string
	Name        string
	Sex         int8
	Birthday    string //用数字表示精确到秒，人类能读懂的方式
	WeddingDate string //用数字表示精确到秒，人类能读懂的方式
	SpouseId    string
	FatherId    string
	ChildrenIds []string
	Deeds       []Deed
	CreateTime 	time.Time
	UpdateTime 	time.Time
}

type Deed struct {
	AppendTime  string //用数字表示精确到秒，人类能读懂的方式
	LeadId      string
	Action 		string
	Object 		string
	Information string
	CreateTime 	time.Time
	UpdateTime 	time.Time
}
