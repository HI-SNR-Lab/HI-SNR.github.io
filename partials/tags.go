package partials

import (
	"net/url"
	"regexp"
	"strings"
)

type Page struct {
	Dir string
}

type Include struct {
	Tags []string
	Link string
	Repo string
}

type TagsView struct {
	Tags  []string
	Link  string
	Repo  string
	Show  bool
}

var ws = regexp.MustCompile(`\s+`)

func BuildTagsView(in Include, page Page) TagsView {
	seen := map[string]struct{}{}
	out := make([]string, 0)

	for _, t := range in.Tags {
		t = strings.ToLower(t)
		parts := strings.Split(t, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			p = ws.ReplaceAllString(p, "-")
			if p == "" {
				continue
			}
			if _, ok := seen[p]; ok {
				continue
			}
			seen[p] = struct{}{}
			out = append(out, p)
		}
	}

	link := in.Link
	if link == "" {
		link = page.Dir
	}

	u, _ := url.Parse(link)

	return TagsView{
		Tags: out,
		Link: u.String(),
		Repo: in.Repo,
		Show: len(out) > 0 || in.Repo != "",
	}
}