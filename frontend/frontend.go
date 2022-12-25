package frontend

import (
	"bytes"
	"dinolog/parser"
	"os"
	"path/filepath"
	"text/template"

	"github.com/kataras/iris/v12"
)

func GenHtmlForBlog(obj *parser.BlogData) string {
	tmpSrc := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>{{ .Meta.Title }}</title>
		    <meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
		</head>
		<body>
		<h1> {{ .Meta.Title }} </h1>
		<h4> {{ .Meta.Author }} </h4>
		<br/>
		<ul>
		{{ range $key , $val := .Follows }}
			<li><a href="{{$val}}"> {{$key}} </a></li>
		{{end}}
		</ul>
		{{ range .Posts }}
			{{ $isDraft := false }} 
			{{ range $key , $val := .BlogMeta }}
				{{ if (and (eq $key "draft") (eq $val "true") ) }}
					{{ $isDraft = true }}
					{{ continue }}
				{{ end }}
			<h3>{{ $key }} = {{ $val }}</h3>
			{{end}}

			{{ if $isDraft }} {{ continue }} {{ end }}

			<hr/>
			<div>
						</div>
			<p>{{ .Date }}</p>
			<pre>
			{{ .Text }}
			</pre>
			<hr/>
		{{end}}
		</body>
	</html>
	`
	tmpl, err := template.New("t").Parse(tmpSrc)

	if err != nil {
		panic(err)
	}

	var data bytes.Buffer
	err = tmpl.Execute(&data, obj)
	if err != nil {
		panic(err)
	}
	return data.String()
}

func Server() {
	app := iris.New()
	app.Get("/{blog}", index)
	app.Get("/" , func(ctx iris.Context) {
		ctx.Redirect("/example" , 301)	
	} )
	app.Use(iris.Compression)
	app.Listen(":8090")
}

func index(ctx iris.Context) {
	cwd, _ := os.Getwd()
	blgName := ctx.Params().Get("blog")
	dlpath := filepath.Join(cwd, blgName+".dinolog.txt")
	dbg := parser.GetDemoObj(dlpath)
	if dbg != nil {
		ctx.HTML(GenHtmlForBlog(dbg))
		//		fmt.Println(GenHtmlForBlog(dbg))
	} else {
		ctx.JSON("404")
	}
}
