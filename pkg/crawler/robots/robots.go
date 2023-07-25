package robots

import (
	"github.com/forestyc/playground/pkg/http"
	"net/url"
	"strings"
)

type Robots struct {
	RobotsUrl string
	UserAgent string
	Disallows []string
}

func NewRobots(host, userAgent string) *Robots {
	r := Robots{
		RobotsUrl: host + "/robots.txt",
		UserAgent: userAgent,
	}
	return &r
}

func (r *Robots) Run() {
	client := http.NewClient(false)
	resp, _ := client.Do("get", r.RobotsUrl, nil, nil)
	r.Disallows = getDisallowByAgent(string(resp), r.UserAgent)
}

func (r *Robots) Disallow(path string) bool {
	if Url, err := url.Parse(path); err == nil {
		if len(Url.Path) == 0 {
			path = "/"
		} else {
			path = Url.Path
		}
	}
	for _, disallow := range r.Disallows {
		if strings.HasPrefix(path, disallow) {
			return true
		}
	}
	return false
}

func getDisallowByAgent(robots, agent string) []string {
	robots = strings.Replace(robots, "\r\n", "\n", -1) // 统一替换\n
	robots = strings.Replace(robots, " ", "", -1)      // 去掉空格
	lines := strings.Split(robots, "\n")
	var result []string
	if !strings.Contains(robots, agent) {
		agent = "*"
	}
	var pass bool
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		values := strings.Split(line, ":")
		if len(values) > 0 {
			if values[0] == "User-agent" {
				if values[1] == agent {
					pass = false
				} else {
					pass = true
				}
			} else if values[0] == "Disallow" {
				if !pass { // 记录disable
					result = append(result, values[1])
				}
			}
		}
	}
	return result
}
