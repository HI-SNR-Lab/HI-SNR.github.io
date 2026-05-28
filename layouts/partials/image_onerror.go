package partials

import "html/template"

func ImageOnError() template.HTMLAttr {
	return template.HTMLAttr(
		`onerror="this.src='images/fallback.svg';this.onerror=null;"`,
	)
}