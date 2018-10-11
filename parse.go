package adblockdomain

import (
	"io"
	"net/url"
	"strings"

	"github.com/pmezard/adblock/adblock"
)

type ruleFilterFunc func(*adblock.Rule) bool

var (
	noException   = func(r *adblock.Rule) bool { return !r.Exception }
	withException = func(r *adblock.Rule) bool { return r.Exception }
)

// ParseFromReader parses list of domain from adblock rules.
func ParseFromReader(r io.Reader) ([]string, error) {
	return parseFromReader(r, noException)
}

// ParseExceptionFromReader parses list of exceptional domain from adblock rules.
func ParseExceptionFromReader(r io.Reader) ([]string, error) {
	return parseFromReader(r, withException)
}

func parseFromReader(r io.Reader, filter ruleFilterFunc) ([]string, error) {
	rules, err := adblock.ParseRules(r)
	if err != nil {
		return nil, err
	}

	var (
		domains   []string
		domainSet = map[string]struct{}{}
	)
	for _, rule := range rules {
		if !filter(rule) {
			continue
		}
		// NOTE: substring match / wildcard will be treats as domain too
		for _, part := range rule.Parts {
			if part.Type != adblock.Exact {
				continue
			}

			rurl := part.Value
			if strings.HasPrefix(rurl, ".") {
				rurl = strings.TrimPrefix(rurl, ".")
			}

			if !strings.Contains(rurl, "://") {
				// add placehold scheme
				rurl = "http://" + rurl
			}

			u, err := url.Parse(rurl)
			if err != nil {
				continue
			}

			domain := u.Hostname()
			if _, exists := domainSet[domain]; !exists {
				domainSet[domain] = struct{}{}
				domains = append(domains, domain)
			}
		}
	}

	return domains, nil
}
