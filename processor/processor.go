package processor

import (
	"github.com/margostino/just/domain"
	"github.com/margostino/just/parser"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

func Process(tokens []html.Token) []*domain.JobPosition {
	var jobs []*domain.JobPosition
	var job *domain.JobPosition
	var isCardTitle, isCardSubtitle, isCardLocation bool

	for _, token := range tokens {
		if job == nil {
			job = &domain.JobPosition{}
		}
		attributes := parser.NewAttributes(token.Attr)
		href := attributes.Get("href")
		class := attributes.Get("class")

		if token.Type == html.StartTagToken && strings.Contains(class, TitleClass) {
			isCardTitle = true
		} else if token.Type == html.StartTagToken && strings.Contains(class, CompanyClass) {
			isCardSubtitle = true
		} else if token.Type == html.StartTagToken && strings.Contains(class, LocationClass) {
			isCardLocation = true
		}

		if token.Type == html.StartTagToken && token.DataAtom == atom.A && href != "" && strings.Contains(class, LinkClass) {
			job.Link = href
		} else if token.Type == html.TextToken && job != nil && isCardTitle {
			job.Title = token.Data
			isCardTitle = false
		} else if token.Type == html.TextToken && job != nil && isCardSubtitle {
			job.Company = token.Data
			isCardSubtitle = false
		} else if token.Type == html.TextToken && job != nil && isCardLocation {
			job.Location = token.Data
			isCardLocation = false
		} else if token.Type == html.StartTagToken && strings.Contains(class, DatetimeClass) {
			job.Datetime = attributes.Get("datetime")
			jobs = append(jobs, job)
			job = nil
		}
	}
	return jobs
}
