package pagination

import (
	"fmt"
	"html/template"
)

// SimplePagination simple paginate
type SimplePagination struct {
	Prev     int
	Next     int
	First    int
	Last     int
	Total    int
	TotalAll int
	Limit    int
	Params   string
}

//Render pagination
func (s *SimplePagination) Render() template.HTML {
	prev := prevHTML(s)
	next := nextHTML(s)

	if s.Total <= s.Limit && s.TotalAll <= s.Limit {
		return template.HTML("")
	}

	return template.HTML(
		fmt.Sprintf(`
			<nav>
				<ul class="pagination justify-content-end">
					%s
					%s
				</ul>
			</nav>
		`, prev, next),
	)
}

func prevHTML(s *SimplePagination) string {
	if s.Prev == 0 {
		return fmt.Sprintf(`
			 <li class="page-item  disabled">
                <a class="page-link" href="/" tabindex="-1">Previous</a>
            </li>
		`)
	}

	return `
			<li class="page-item">
			<a class="page-link" href="javascript: history.go(-1)" tabindex="-1">Previous</a>
		</li>
	`

}

func nextHTML(s *SimplePagination) string {
	if s.Next == s.Last {
		return `
			 <li class="page-item disabled">
                <a class="page-link" href="#">Next</a>
            </li>
		`
	}

	return fmt.Sprintf(`
			<li class="page-item">
			<a class="page-link" href="/?last_id=%v%v">Next</a>
		</li>
	`, s.Next, s.Params)

}
