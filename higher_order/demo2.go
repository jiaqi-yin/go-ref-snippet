package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type RequestFilter func(*http.Request) bool
type Middleware func(http.HandlerFunc) http.HandlerFunc
type Filters []RequestFilter
type Stack []Middleware
type Endpoint struct {
	Handler    http.HandlerFunc
	Allow      Filters
	Middleware Stack
}

type Routes map[string]Endpoint
type RouteManager struct {
	RouteMap         Routes
	GlobalMiddleware Stack
}

func CIDR(cidrs ...string) RequestFilter {
	nets := make([]*net.IPNet, len(cidrs))
	for i, cidr := range cidrs {
		_, nets[i], _ = net.ParseCIDR(cidr)
	}
	return func(r *http.Request) bool {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		ip := net.ParseIP(host)
		for _, net := range nets {
			if net.Contains(ip) {
				return true
			}
		}
		return false
	}
}

func PasswordHeader(password string) RequestFilter {
	return func(r *http.Request) bool {
		return r.Header.Get("X-Password") == password
	}
}

func Method(methods ...string) RequestFilter {
	return func(r *http.Request) bool {
		for _, m := range methods {
			if r.Method == m {
				return true
			}
		}
		return false
	}
}

func Allow(filter RequestFilter) Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(responseWriter http.ResponseWriter, request *http.Request) {
			if filter(request) {
				handlerFunc(responseWriter, request)
			} else {
				responseWriter.WriteHeader(http.StatusForbidden)
			}
		}
	}
}

func SetHeader(key, value string) Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set(key, value)
			hf(rw, r)
		}
	}
}

func Logging(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Printf("[%v] - %s %s\n", time.Now(), request.Method, request.RequestURI)
		handlerFunc(responseWriter, request)
	}
}

func (f Filters) Combine() RequestFilter {
	return func(r *http.Request) bool {
		for _, filter := range f {
			if !filter(r) {
				return false
			}
		}
		return true
	}
}

func (s Stack) Apply(f http.HandlerFunc) http.HandlerFunc {
	g := f
	for _, middleware := range s {
		g = middleware(g)
	}
	return g
}

func (e Endpoint) Build() http.HandlerFunc {
	allowFilter := e.Allow.Combine()
	restricted := Allow(allowFilter)(e.Handler)
	return e.Middleware.Apply(restricted)
}

func (r RouteManager) Serve(addr string) error {
	mux := http.NewServeMux()
	for pattern, endpoint := range r.RouteMap {
		mux.Handle(pattern, r.GlobalMiddleware.Apply(endpoint.Build()))
	}
	return http.ListenAndServe(addr, mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	routes := RouteManager{
		GlobalMiddleware: Stack{
			Logging,
			SetHeader("X-Foo", "Bar"),
		},
		RouteMap: Routes{
			"/hello": {
				Handler: hello,
			},
			"/private": {
				Handler: hello,
				Allow: Filters{
					CIDR("127.0.0.1/32"),
					PasswordHeader("foo"),
					Method("GET"),
				},
			},
			"/test": {
				Handler: hello,
				Middleware: Stack{
					SetHeader("X-Test", "Y"),
				},
			},
		},
	}

	log.Fatal(routes.Serve(":8080"))
}
