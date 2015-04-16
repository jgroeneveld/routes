package routes

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type RoutesDef struct {
	Imports []Import
	Routes  []Route
}

type Import string

type Route struct {
	Method  string
	Path    string
	Handler string
}

func ParseRoutes(reader io.Reader) (def RoutesDef, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	out := RoutesDef{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || isCommentLine(line) {
			continue
		}
		i, isImport := parseImportLine(line)
		if isImport {
			out.Imports = append(out.Imports, i)
			continue
		}
		r, isRoute := parseRouteLine(line)
		if isRoute {
			out.Routes = append(out.Routes, r)
			continue
		}
	}

	return out, nil
}

func isCommentLine(line string) bool {
	return strings.HasPrefix(line, "//")
}

func parseImportLine(line string) (Import, bool) {
	if !strings.HasPrefix(line, "import") {
		return "", false
	}

	fields := strings.Fields(line)
	if len(fields) != 2 {
		panic("Line starting with 'import' but has more than 2 fields")
	}

	importStatement := strings.Replace(fields[1], "\"", "", -1)
	return Import(importStatement), true
}

func parseRouteLine(line string) (Route, bool) {
	if !beginsWithHTTPVerb(line) {
		return Route{}, false
	}

	fields := strings.Fields(line)
	if len(fields) != 3 {
		panic(fmt.Sprintf("Can not parse route definition, line contains more than 3 fields. (%s)", line))
	}

	return Route{
		Method:  strings.ToUpper(fields[0]),
		Path:    fields[1],
		Handler: fields[2],
	}, true
}

func beginsWithHTTPVerb(str string) bool {
	verbs := []string{"GET", "PUT", "PATCH", "POST", "HEAD", "DELETE"}
	str = strings.ToUpper(str)

	for _, v := range verbs {
		if strings.HasPrefix(str, v) {
			return true
		}
	}

	return false
}
