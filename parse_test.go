package adblockdomain

import (
	"io"
	"strings"
	"testing"
)

func newInput(lines ...string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}

func domainListShouldContain(t *testing.T, ds []string, d string) {
	for _, domain := range ds {
		if domain == d {
			return
		}
	}

	t.Errorf("%v should contain %s", ds, d)
}

func domainListShouldNotContain(t *testing.T, ds []string, d string) {
	for _, domain := range ds {
		if domain == d {
			t.Errorf("%v should not contain %s", ds, d)
		}
	}
}

func TestParseFromReader(t *testing.T) {
	r := newInput(`
        |foo.com
        ||foo.org
        foo.xyz
        @@foo.io
        !foo.xxx
        `)

	domains, err := ParseFromReader(r)
	if err != nil {
		t.Error(err)
	}
	domainListShouldContain(t, domains, "foo.com")
	domainListShouldContain(t, domains, "foo.xyz")
	domainListShouldContain(t, domains, "foo.org")
	domainListShouldNotContain(t, domains, "foo.io")
	domainListShouldNotContain(t, domains, "foo.xxx")
}

func TestParseExceptionFromReader(t *testing.T) {
	r := newInput(`
        |foo.com
        ||foo.org
        foo.xyz
        @@foo.io
        !foo.xxx
        `)

	domains, err := ParseExceptionFromReader(r)
	if err != nil {
		t.Error(err)
	}
	domainListShouldNotContain(t, domains, "foo.com")
	domainListShouldNotContain(t, domains, "foo.xyz")
	domainListShouldNotContain(t, domains, "foo.org")
	domainListShouldContain(t, domains, "foo.io")
	domainListShouldNotContain(t, domains, "foo.xxx")
}
