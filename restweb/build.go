package main

import (
	"bytes"
	"go/ast"
	// "go/build"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ControllerInfo struct {
	PkgPath string
	PkgName string
	Name    string
}
type RouterInfo struct {
	ControllerName string
	URL            template.HTML
	Action         string
	Method         string
}

var ContrInfos []ControllerInfo
var RouterInfos []RouterInfo

func buildApp() {
	filepath.Walk(appName+"/controller", walkFn)
	generateMain()
	generateRouter()
	os.Chdir(appName)
	cmd := exec.Command("go", "build")
	cmd.Stdout = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Build Failed\n[Error] ", err)
	} else {
		log.Println("Build Succeed")
	}
}

func walkFn(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		return nil
	}

	fset := &token.FileSet{}
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		log.Println(err)
		return err
	}
	var pkg *ast.Package
	for _, v := range pkgs {
		pkg = v
	}

	walkAstFiles(fset, path, pkg)

	return nil
}

func walkAstFiles(fset *token.FileSet, path string, pkg *ast.Package) {
	ControllerName := ""
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {

			if funcdecl, ok := decl.(*ast.FuncDecl); ok && funcdecl.Doc != nil {
				for _, cmt := range funcdecl.Doc.List {
					if strings.HasPrefix(cmt.Text, "//@") {
						adec.Clear()
						err := phaseApp(cmt.Text)
						if err != nil {
							log.Fatal("[error] ", fset.Position(cmt.Pos()), err)
						}

						RouterInfos = append(RouterInfos,
							RouterInfo{ControllerName: ControllerName,
								URL:    template.HTML(adec.URL),
								Action: funcdecl.Name.Name,
								Method: adec.Method})
					}
				}
			}

			if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
				spec := gen.Specs[0]
				if ts, ok := spec.(*ast.TypeSpec); ok && strings.ToLower(ts.Comment.Text()) == "@controller\n" {
					ControllerName = ts.Name.Name
					ContrInfos = append(ContrInfos, ControllerInfo{PkgPath: path, PkgName: pkg.Name, Name: ControllerName})
				}
			}
		}
	}
}

func generateMain() {

	tpl := `package main

	import (
	"restweb"
	"log"
	{{with .ContrInfos}}
	{{range .}}"{{.PkgPath}}"
	{{end}}
	{{end}}
	)

	func main(){
	{{with .ContrInfos}}
	{{range .}}restweb.RegisterController(&{{.PkgName}}.{{.Name}}{})
	{{end}}
	{{end}}
	restweb.AddFile("/static/", ".")
	log.Fatal(restweb.Run())
	}
	`
	t, err := template.New("foo").Parse(tpl)
	if err != nil {
		log.Println(err)
		return
	}
	bf := bytes.NewBufferString("")
	data := make(map[string]interface{})
	data["ContrInfos"] = ContrInfos
	err = t.Execute(bf, data)
	if err != nil {
		log.Println(err)
	}
	b, err := format.Source([]byte(bf.String()))
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create(appName + "/main.go")
	if err != nil {
		log.Println(err)
		return
	}
	f.Write(b)
}

func generateRouter() {
	tpl := `
{{with .RouterInfos}}
{{range .}}{{.Method}} 	^{{.URL}}$	 {{.ControllerName}}.{{.Action}}
{{end}}
{{end}}
`
	t, err := template.New("foo").Parse(tpl)
	if err != nil {
		log.Println(err)
		return
	}
	bf, err := os.Create(appName + "/config/router.conf")
	if err != nil {
		log.Println(err)
	}
	data := make(map[string]interface{})
	data["RouterInfos"] = RouterInfos
	err = t.Execute(bf, data)
	if err != nil {
		log.Println(err)
	}
}
