package page

import "math"

type Info struct {
	CurrentPage int
	Total int
	PageSize int
	totalPage int
	FirstPage int
	LastPage  int
}

func New(currentPage int,total int,pageSize int) *Info {
	if pageSize==0 {
		pageSize = 10
	}
	if currentPage ==0 {
		currentPage =1
	}
	return  &Info{CurrentPage:currentPage,PageSize:pageSize,Total:total}
}
func (info *Info)TotalPage() int {
	return	int(math.Ceil(float64(info.Total)/float64(info.PageSize)))
}