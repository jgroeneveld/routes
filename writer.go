package routes

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

func Write(pkg string, rd RoutesDef, w io.Writer) error {
	tpl, err := template.New("routes").Parse(tpl)
	if err != nil {
		return err
	}
	return tpl.Execute(w, routesDefFile{
		Package:   pkg,
		RoutesDef: rd,
	})
	//	tpl.par
}

type routesDefFile struct {
	Package string
	RoutesDef
}

func (rdf routesDefFile) PathHelpers() []string {
	out := []string{}
	pathsHandled := map[string]bool{}

	for _, r := range rdf.Routes {
		_, ok := pathsHandled[r.Path]
		if ok {
			continue
		}

		params := []string{}
		for _, paramName := range r.ParamNames() {
			params = append(params, fmt.Sprintf("%s %s", paramName, "string"))
		}

		helper := fmt.Sprintf("func %sPath(%s) {}", RouteNameForPath(r.Path), strings.Join(params, ", "))
		out = append(out, helper)
		pathsHandled[r.Path] = true
	}

	return out
}

const tpl = `package {{ .Package }}

import (
	"github.com/julienschmidt/httprouter"
{{ range .Imports }}
	"{{ . }}"{{ end }}
)

func Router() *httprouter.Router {
	router := httprouter.New()

{{ range .Routes}}	router.{{ .Method }}("{{ .Path }}", {{ .Handler }})
{{ end }}
	return router
}
{{ range .PathHelpers }}
{{ . }}
{{end}}`
