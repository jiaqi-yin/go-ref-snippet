package main

import (
	"fmt"
	"strings"
)

type Host string
type HostList []Host
type HostFilter func(Host) bool

var Or = func(clauses ...HostFilter) HostFilter {
	var f HostFilter = clauses[0]
	for _, c := range clauses[1:] {
		f = f.Or(c)
	}
	return f
}

func (l HostList) Select(f HostFilter) HostList {
	result := make(HostList, 0, len(l))
	for _, h := range l {
		if f(h) {
			result = append(result, h)
		}
	}
	return result
}

func (f HostFilter) Or(g HostFilter) HostFilter {
	return func(h Host) bool {
		return f(h) || g(h)
	}
}

func (f HostFilter) And(g HostFilter) HostFilter {
	return func(h Host) bool {
		return f(h) && g(h)
	}
}

var IsDotOrg HostFilter = func(h Host) bool {
	return strings.HasSuffix(string(h), ".org")
}

var HasGo HostFilter = func(h Host) bool {
	return strings.Contains(string(h), "go")
}

var IsAcademic HostFilter = func(h Host) bool {
	return strings.Contains(string(h), "academy")
}

func main() {
	hostnames := HostList{
		"golang.org",
		"google.com",
		"gopheracademy.org",
	}

	fmt.Printf("%v\n", hostnames)
	fmt.Printf("Go sites: %v\n", hostnames.Select(HasGo))
	fmt.Printf("Academies: %v\n", hostnames.Select(IsDotOrg.And(IsAcademic)))
	fmt.Printf("Or %v\n", hostnames.Select(Or(IsDotOrg, IsAcademic)))
	fmt.Printf("OrAnd %v\n", hostnames.Select(Or(IsDotOrg, HasGo).And(IsAcademic)))
}
