package ir

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func splitLine(s string, limit int) (r []string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	// TODO(tdakkota): handle links in docs
	for {
		if len(s) < limit {
			r = append(r, s)
			return
		}

		idx := strings.LastIndexFunc(s[:limit-1], func(r rune) bool {
			return unicode.IsSpace(r) || r == '.' || r == ','
		})
		if idx < 0 || len(s)-1 == idx {
			r = append(r, s)
			return
		}

		if ch, _ := utf8.DecodeRuneInString(s[idx:]); unicode.IsSpace(ch) {
			r = append(r, s[:idx])
			s = s[idx+1:]
		} else {
			// Do not cut dots.
			r = append(r, s[:idx+1])
			s = s[idx+1:]
		}
	}
}

func prettyDoc(s string) (r []string) {
	const (
		lineLimit = 100
	)

	// TODO(tdakkota): basic common mark rendering?
	for _, line := range strings.Split(s, "\n") {
		r = append(r, splitLine(line, lineLimit)...)
	}
	if len(r) > 0 {
		r[0] = capitalize(r[0])

		if last := r[len(r)-1]; len(last) > 0 && last[len(last)-1] != '.' {
			r[len(r)-1] = last + "."
		}
	}

	return r
}