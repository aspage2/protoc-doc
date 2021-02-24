package protodoc

import (
	"bytes"
	"embed"
	"google.golang.org/protobuf/compiler/protogen"
	"html/template"
	"io"
)

var (
	//go:embed template
	templates embed.FS

	//go:embed static/style.css
	style []byte
)

var fm = template.FuncMap{
	"allenums":    getEnums,
	"allmessages": getMessages,
	"comments":    commentString,
	"fileloc":     fileLoc,
	"messageid":   messageId,
}

func RealMain(p *protogen.Plugin) error {
	var (
		tmpl *template.Template
		err  error
	)
	_, err = io.Copy(
		p.NewGeneratedFile("static/style.css", ""),
		bytes.NewBuffer(style),
	)
	if err != nil {
		return err
	}

	tmpl = template.New("").Funcs(fm)
	tmpl, err = tmpl.ParseFS(templates, "template/*.gohtml")
	if err != nil {
		return err
	}

	gf := p.NewGeneratedFile("index.html", "")
	if err := tmpl.ExecuteTemplate(gf, "index.gohtml", sortedFiles(p.Files)); err != nil {
		return err
	}

	for _, f := range p.Files {
		gf := p.NewGeneratedFile(fileLoc(f.Desc), "")
		if err := tmpl.ExecuteTemplate(gf, "file.gohtml", f); err != nil {
			return err
		}
	}
	return nil
}
