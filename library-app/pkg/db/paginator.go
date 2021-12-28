package db

type Pagination struct {
	Limit      int         `form:"limit,omitempty;query:limit"`
	Page       int         `form:"page,omitempty;query:page"`
	Sort       string      `form:"sort,omitempty;query:sort"`
	TotalRows  int64       `form:"total_rows"`
	TotalPages int         `form:"total_pages"`
	Rows       interface{} `form:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
