package users

type FindAllCriteria struct {
	Name string `json:"name" uri:"name" form:"name"`
	Age  int    `json:"age" uri:"age" form:"age"`
}

type OverviewCriteria struct {
	Name string `json:"name" uri:"name" form:"name"`
	Age  int    `json:"age" uri:"age" form:"age"`
}
