package router

import (
	"errors"
	"regexp"
	"strings"
)

type Path string

func (p *Path) String() string {
	return string(*p)
}

func (p *Path) IsEqual(path Path) bool {
	return string(*p) == string(path)
}

func (p *Path) IsRoot() bool {
	return string(*p) == "/"
}

func (p *Path) IsValid() (bool, error) {
	path := string(*p)

	pathSize := len(path)
	if pathSize == 0 {
		return false, errors.New("path should not be empty")
	}
	if path[0] != '/' {
		return false, errors.New("path should start with '/'")
	}
	if match, _ := regexp.MatchString("^(/:?[a-z0-9]+(-[a-z0-9]+)*?)*/?$", path); !match {
		return false, errors.New("is not a valid rest path")
	}
	return true, nil
}

func (p *Path) IsRouteValid() (bool, error) {
	path := string(*p)

	pathSize := len(path)
	if pathSize == 0 {
		return false, errors.New("path should not be empty")
	}
	if path[0] != '/' {
		return false, errors.New("path should start with '/'")
	}
	if match, _ := regexp.MatchString("^(/[a-z0-9]+(-[a-z0-9]+)*?)*/?$", path); !match {
		return false, errors.New("is not a valid rest path")
	}
	return true, nil
}

func (p *Path) IsParam() bool {
	path := string(*p)

	if match, _ := regexp.MatchString("^/:[a-z]+(/:?[a-z0-9]+(-[a-z0-9]+)*?)*/?$", path); match {
		return true
	}

	return false
}

func (p *Path) Split() []Path {
	path := string(*p)
	splittedPath := strings.Split(path, "/")

	paths := make([]Path, len(splittedPath)-1)
	for i, str := range splittedPath {
		if i != 0 {
			paths[i-1] = Path("/" + str)
		}
	}

	return paths
}

func (p *Path) Merge(path Path) Path {
	leftPath := string(*p)
	rightPath := string(path)
	return Path(strings.TrimSuffix(leftPath, "/") + rightPath)
}

func (p *Path) GetRoutePattern() string {
	path := string(*p)
	re := regexp.MustCompile(":[a-z-]+")
	return re.ReplaceAllString(path, "<>")
}

func (p *Path) IsSubtreePath() bool {
	return strings.HasSuffix(string(*p), "/")
}