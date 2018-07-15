package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/template"
	"github.com/factorapp/factor/files"

	"github.com/factorapp/factor/component"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	cdir := filepath.Join(cwd, "components")
	// err = component.ProcessAll(dir, "components")
	err = ProcessComponents(cdir)
	if err != nil {
		fmt.Println(err)
		return
	}

	pdir := filepath.Join(cwd, "pages")
	// err = component.ProcessAll(dir, "components")
	err = ProcessPages(pdir)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Context struct {
	Imports    string
	Funcs      string
	Template   string
	Name       string
	Parameters string
	Title      string
}

func ProcessComponents(dir string) error {

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && files.IsHTML(info) {
			// Parse the ghtml file

			f, err := os.Open(path)
			if err != nil {
				return err
			}

			contents, err := ioutil.ReadAll(f)

			c := &Context{
				Name: files.ComponentName(info.Name()),
			}
			scontents := string(contents)

			// get the template
			idx := strings.Index(scontents, "@")

			rawTemplate := scontents[:idx]
			c.Template = "`" + scontents[:idx] + "`"

			// get Parameters
			pidx := strings.Index(scontents, "@Parameters")
			c.Parameters = insideBrackets(scontents[pidx:])

			// get Imports
			iidx := strings.Index(scontents, "@Imports")
			c.Imports = insideBrackets(scontents[iidx:])

			// get Methods
			//midx := strings.Index(scontents, "@Methods")
			c.Funcs = methods(scontents)

			br := bytes.NewReader([]byte(rawTemplate))
			brc := ioutil.NopCloser(br)
			// write the Go file
			comp := files.ComponentName(path)
			transpiler, err := component.NewTranspiler(brc, false, "components", comp, "components")
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			//fmt.Println(transpiler.Code())

			rfn := filepath.Join(dir, strings.ToLower(comp)+"_render.go")
			_, err = os.Stat(rfn)
			renderfile, err := os.Create(GeneratedGoFileName(dir, comp+"_render"))
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			defer renderfile.Close()
			_, err = io.WriteString(renderfile, transpiler.Code())
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			gfn := filepath.Join(dir, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			gofile, err := os.Create(GeneratedGoFileName(dir, comp))
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			defer gofile.Close()
			tpl := template.Must(template.New("component").Parse(compTemplate))
			err = tpl.Execute(gofile, c)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("error walking the path %q: %v\n", dir, err)
	}
	return err
}

func methods(s string) string {
	start := strings.Index(s, "@Methods")
	rest := s[start:]
	rest = strings.Replace(rest, "@Methods", "", -1)

	return rest
}
func title(s string) string {
	start := strings.Index(s, "@Title")
	rest := s[start:]
	rest = strings.Replace(rest, "@Title", "", -1)
	end := strings.Index(rest, "\n")
	return rest[:end]
}
func insideBrackets(s string) string {
	start := strings.Index(s, "{")
	rest := s[start+1:]
	end := strings.Index(rest, "}")

	return rest[:end]

}

func GeneratedGoFileName(base, name string) string {
	return filepath.Join(base, strings.ToLower(name)+".go")
}

func ProcessPages(dir string) error {

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && files.IsHTML(info) {
			// Parse the ghtml file

			f, err := os.Open(path)
			if err != nil {
				return err
			}

			contents, err := ioutil.ReadAll(f)

			c := &Context{
				Name: files.ComponentName(info.Name()),
			}
			scontents := string(contents)

			// get the template
			idx := strings.Index(scontents, "@")

			rawTemplate := scontents[:idx]
			c.Template = "`" + scontents[:idx] + "`"

			// get Parameters
			pidx := strings.Index(scontents, "@Parameters")
			c.Parameters = insideBrackets(scontents[pidx:])

			// get Title
			tidx := strings.Index(scontents, "@Title")
			c.Title = title(scontents[tidx:])
			// get Imports
			iidx := strings.Index(scontents, "@Imports")
			c.Imports = insideBrackets(scontents[iidx:])

			// get Methods
			//midx := strings.Index(scontents, "@Methods")
			c.Funcs = methods(scontents)

			br := bytes.NewReader([]byte(rawTemplate))
			brc := ioutil.NopCloser(br)
			// write the Go file
			comp := files.ComponentName(path)

			//		appPkg := envy.CurrentPackage()
			//		fmt.Println(appPkg)
			appPkg := "github.com/factorapp/playground"
			transpiler, err := component.NewTranspiler(brc, false, appPkg, comp, "pages")
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			//fmt.Println(transpiler.Code())

			rfn := filepath.Join(dir, strings.ToLower(comp)+"_render.go")
			_, err = os.Stat(rfn)
			renderfile, err := os.Create(GeneratedGoFileName(dir, comp+"_render"))
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			defer renderfile.Close()
			_, err = io.WriteString(renderfile, transpiler.Code())
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			gfn := filepath.Join(dir, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			gofile, err := os.Create(GeneratedGoFileName(dir, comp))
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			defer gofile.Close()
			tpl := template.Must(template.New("pages").Parse(pageTemplate))
			err = tpl.Execute(gofile, c)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("error walking the path %q: %v\n", dir, err)
	}
	return err
}
