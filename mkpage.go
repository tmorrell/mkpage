//
// Package mkpage is an experiment in a light weight template and markdown processor.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package mkpage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	// 3rd Party Packages
	"github.com/russross/blackfriday"
)

const (
	// Version of the mkpage package.
	Version = "v0.0.16"

	// LicenseText provides a string template for rendering cli license info
	LicenseText = `
%s %s

Copyright (c) 2017, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`
	// Prefix for explicit string types

	// JSONPrefix designates a string as JSON formatted content
	JSONPrefix = "json:"
	// MarkdownPrefix designates a string as Markdown content
	MarkdownPrefix = "markdown:"
	// TextPrefix designates a string as text/plain not needed processing
	TextPrefix = "text:"

	// SOMEDAY: should add XML, BibTeX, YaML support...

	// DefaultTemplateSource holds the default HTML provided by mkpage package, you probably want to override this...
	DefaultTemplateSource = `<!DOCTYPE html>
<html>
<head>
  {{if .title -}}<title>{{- .title -}}</title>{{- end}}
  {{if .csspath -}}<link href="{{ .csspath }}" rel="stylesheet" />{{else -}}
  <style>
/**
 * site.css - stylesheet for the Library's Digital Library Development Group's sandbox.
 *
 * orange: #FF6E1E;
 *
 * Secondary pallet:
 *
 * lightgrey: #C8C8C8
 * grey: #76777B
 * darkgrey: #616265
 * slategrey: #AAA99F
 *
 * Impact Pallete see: http://identity.caltech.edu/web/colors
 */
body {
     margin: 0;
     border: 0;
     padding: 0;
     width: 100%;
     height: 100%;
     color: black;
     background-color: white;
     /*
     color: #FF6E1E;
     background-color: #AAA99F; /* #76777B;*/
     */
     font-family: Open Sans, Helvetica, Sans-Serif;
     font-size: 16px;
}

header {
     position: relative;
     display: block;
     color: white;
     background-color: white;
     z-index: 2;
     height: 105px;
     vertical-align: middle;
}

header img {
     position: relative;
     display: inline;
     padding-left: 20px;
     margin: 0;
     height: 42px;
     padding-top: 32px;
}

header h1 {
     position: relative;
     display: inline-block;
     margin: 0;
     border: 0;
     padding: 0;
     font-size: 3em;
     font-weight: normal;
     vertical-align: 0.78em;
     color: #FF6E1E;
}

header a, header a:link, header a:visited, header a:active, header a:hover, header a:focus {
     color: #FF6E1E;
     background-color: inherit;
}


a, a:link, a:visited {
     color: #76777B;
     background-color: inherit;
     text-decoration: none;
}

a:active, a:hover, a:focus {
    color: #FF6E1E;
    font-weight: bolder;
}

nav {
     position: relative;
     display: block;
     width: 100%;
     margin: 0;
     padding: 0;
     font-size: 0.78em;
     vertical-align: middle;
     color: black;
     background-color: #AAA99F; /* #76777B;*/
     text-align: left;
}

nav div {
     display: inline-block;
     /* padding-left: 10em; */
     margin-left: 10em;
     margin-right: 0;
}

nav a, nav a:link, nav a:visited, nav a:active {
     color: white;
     background-color: inherit;
     text-decoration: none;
}

nav a:hover, nav a:focus {
     color: #FF6E1E;
     background-color: inherit;
     text-decoration: none;
}


nav div h2 {
     position: relative;
     display: block;
     min-width: 20%;
     margin: 0;
     font-size: 1.24em;
     font-style: normal;
}

nav div > ul {
     display: none;
     padding-left: 0.24em;
     text-align: left;
}

nav ul {
     display: inline-block;
     padding-left: 0.24em;
     list-style-type: none;
     text-align: left;
     text-decoration: none; 
}

nav ul li {
     display: inline;
     padding: 1em;
}

section {
     position: relative;
     display: inline-block;
     width: 100%;
     min-height: 84%;
     color: black;
     background-color: white;
     margin: 0;
     padding-top 0;
     padding-bottom: 2em;
     padding-left: 1em;
     padding-right: 0;
}

section h1 {
    font-size: 1.32em;
}

section h2 {
    font-size: 1.12em;
    font-weight: italic;
}

section h3 {
    font-size: 1em;
    text-transform: uppercase;
}

section ul {
    display: block;
    list-style: inside;
    list-style-type: square;
    margin: 0;
    padding-left: 1.24em;
}

aside {
     margin: 0;
     border: 0;
     padding-left: 1em;
     position: relative;
     display: inline-block;
     text-align: right;
}

aside h2 {
     font-size: 1em;
     text-transform: uppercase;
}

aside h2 > a {
     font-style: normal;
}

aside ul {
     margin: 0;
     padding: 0;
     display: block;
     list-style-type: none;
}

aside ul li {
     font-size: 0.82em;
}

aside ul > ul {
     padding-left: 1em;
     font-size: 0.72em;
}

footer {
     position: fixed;
     bottom: 0;
     display: block;
     width: 100%;
     height: 2em;
     color: white;
     background-color: #616265;

     font-size: 80%;
     text-align: left;
     vertical-align: middle;
     z-index: 2;
}

footer h1, footer span, footer address {
     position: relative;
     display: inline-block;
     margin: 0;
     padding-left: 0.24em;
     font-family: Open Sans, Helvetica, Sans-Serif;
     font-size: 1em;
}

footer h1 {
     font-weight: normal;
}

footer a, footer a:link, footer a:visited, footer a:active, footer a:focus, footer a:hover {
     padding: 0;
     display: inline;
     margin: 0;
     color: #FF6E1E;
     text-decoration: none;
}

  </style>
  {{- end }}
</head>
<body>
  {{if .header -}}
  <header>{{- .header -}}</header>
  {{end}}
  {{if .nav -}}
  <nav>{{- .nav -}}</nav>
  {{end}}
  {{if .content -}}
  <section>{{ .content }}</section>
  {{end}}
  {{if .footer -}}
  <footer>{{ .footer }}</footer>
  {{end}}
</body>
</html>
`

	// DateExp is the default format used by mkpage utilities for date exp
	DateExp = `[0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9]`
	// BylineExp is the default format used by mkpage utilities
	BylineExp = (`^[B|b]y\s+(\w|\s|.)+` + DateExp + "$")
	// TitleExp is the default format used by mkpage utilities
	TitleExp = `^#\s+(\w|\s|.)+$`
)

// ResolveData takes a data map and reads in the files and URL sources
// as needed turning the data into strings to be applied to the template.
func ResolveData(data map[string]string) (map[string]interface{}, error) {
	var out map[string]interface{}

	isContentType := func(vals []string, target string) bool {
		for _, h := range vals {
			if strings.Contains(h, target) == true {
				return true
			}
		}
		return false
	}

	out = make(map[string]interface{})
	for key, val := range data {
		switch {
		case strings.HasPrefix(val, TextPrefix) == true:
			out[key] = strings.TrimPrefix(val, TextPrefix)
		case strings.HasPrefix(val, MarkdownPrefix) == true:
			out[key] = string(blackfriday.MarkdownCommon([]byte(strings.TrimPrefix(val, MarkdownPrefix))))
		case strings.HasPrefix(val, JSONPrefix) == true:
			var o interface{}
			err := json.Unmarshal(bytes.TrimPrefix([]byte(val), []byte(JSONPrefix)), &o)
			if err != nil {
				return out, fmt.Errorf("Can't JSON decode %s, %s", val, err)
			}
			out[key] = o

		case strings.HasPrefix(val, "http://") == true || strings.HasPrefix(val, "https://") == true:
			resp, err := http.Get(val)
			if err != nil {
				return out, fmt.Errorf("Error from %s, %s", val, err)
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				buf, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return out, err
				}
				if contentTypes, ok := resp.Header["Content-Type"]; ok == true {
					switch {
					case isContentType(contentTypes, "application/json") == true:
						var o interface{}
						err := json.Unmarshal(buf, &o)
						if err != nil {
							return out, fmt.Errorf("Can't JSON decode %s, %s", val, err)
						}
						out[key] = o
					case isContentType(contentTypes, "text/markdown") == true:
						out[key] = string(blackfriday.MarkdownCommon(buf))
					default:
						out[key] = string(buf)
					}
				} else {
					out[key] = string(buf)
				}
			}
		default:
			buf, err := ioutil.ReadFile(val)
			if err != nil {
				return out, fmt.Errorf("Can't read %s, %s", val, err)
			}
			ext := path.Ext(val)
			switch {
			case strings.Compare(ext, ".md") == 0:
				out[key] = string(blackfriday.MarkdownCommon(buf))
			case strings.Compare(ext, ".json") == 0:
				var o interface{}
				err := json.Unmarshal(buf, &o)
				if err != nil {
					return out, fmt.Errorf("Can't JSON decode %s, %s", val, err)
				}
				out[key] = o
			default:
				out[key] = string(buf)
			}
		}
	}
	return out, nil
}

// MakePage applies the key/value map to the template and renders to writer and returns an error if something goes wrong
func MakePage(wr io.Writer, tmpl *template.Template, keyValues map[string]string) error {
	data, err := ResolveData(keyValues)
	if err != nil {
		return fmt.Errorf("Can't resolve data source %s", err)
	}
	return tmpl.Execute(wr, data)
}

// MakePageString applies the key/value map to the template and renders the results to a string and error if someting goes wrong
func MakePageString(tmpl *template.Template, keyValues map[string]string) (string, error) {
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	err := MakePage(wr, tmpl, keyValues)
	return buf.String(), err
}

// RelativeDocPath calculate the relative path from source to target based on
// implied common base.
//
// Example:
//
//     docPath := "docs/chapter-01/lesson-02.html"
//     cssPath := "css/site.css"
//     fmt.Printf("<link href=%q>\n", MakeRelativePath(docPath, cssPath))
//
// Output:
//
//     <link href="../../css/site.css">
//
func RelativeDocPath(source, target string) string {
	var result []string

	sep := string(os.PathSeparator)
	dname, _ := path.Split(source)
	for i := 0; i < strings.Count(dname, sep); i++ {
		result = append(result, "..")
	}
	result = append(result, target)
	return strings.Join(result, sep)
}

//
// Below is addition code to support mkslides
//

// Slide is the metadata about a slide to be generated.
type Slide struct {
	CurNo   int
	PrevNo  int
	NextNo  int
	FirstNo int
	LastNo  int
	FName   string
	Title   string
	Content string
	CSSPath string
	JSPath  string
}

var (
	// DefaultSlideTemplateSource provides the default HTML template for mkslides package, you probably want to override this...
	DefaultSlideTemplateSource = `<!DOCTYPE html>
<html>
<head>
	{{if .Title -}}<title>{{- .Title -}}</title>{{- end}}
	{{if .CSSPath -}}
<link href="{{ .CSSPath }}" rel="stylesheet" />
   {{else -}}
<style>
    body {
    	width: 100%;
    	height: 100%;
    	margin: 10%;
    	padding: 0;
    	font-size: 24px;
    	font-family: sans-serif;
    }
    
    ul {
    	list-style: circle;
    	text-indent: 0.25em;
    }
    
    nav {
    	position: absolute;
    	top: 0em; 
    	margin:0;
    	padding:0.24em;
    	width: 100%;
    	height: 4em;
    	text-align: left;
    	font-size: 60%;
    }
    
	section, p {
		max-width: 85%;
		padding: 0.24em;
		margin: 0.24em;
	}

	code {
		color: teal;
	}
</style>
{{- end }}
</head>
<body>
	<nav>
{{ if ne .CurNo .FirstNo -}}
<a id="start-slide" href="{{printf "%02d-%s.html" .FirstNo .FName}}">Home</a>
{{- end}}
{{ if gt .CurNo .FirstNo -}} 
<a id="previ-slide" href="{{printf "%02d-%s.html" .PrevNo .FName}}">Prev</a>
{{- end}}
{{ if lt .CurNo .LastNo -}} 
<a id="next-slide" href="{{printf "%02d-%s.html" .NextNo .FName}}">Next</a>
{{- end}}
	</nav>
	<section>{{ .Content }}</section>
<script>
(function (document, window) {
    'use strict';
    var start = document.getElementById('start-slide'),
        prev = document.getElementById('prev-slide'),
        next = document.getElementById('next-slide');
    
    
    document.onkeydown = function(e) {
        switch (e.keyCode) {
            /* case 32: */
            case 37:
            // Previous: left arrow
                prev.click();
                break;
            case 39:
                // Next: right arrow
                next.click();
                break;
            case 72:
            case 83:
                // Home/Start: h, s
                start.click();
                break;
        }
    };
}(document, window));
</script>
</body>
</html>
`
)

// MarkdownToSlides turns a markdown file into one or more Slide using the fname, title and cssPath provided
func MarkdownToSlides(fname, title, cssPath, jsPath string, mdSource []byte) []*Slide {
	var slides []*Slide

	// Note: handle legacy CR/LF endings as well as normal LF line endings
	if bytes.Contains(mdSource, []byte("\r\n")) {
		mdSource = bytes.Replace(mdSource, []byte("\r\n"), []byte("\n"), -1)
	}
	// Note: We're only spliting on line that contain "--", not lines ending with where text might end with "--"
	mdSlides := bytes.Split(mdSource, []byte("\n--\n"))

	lastSlide := len(mdSlides) - 1
	for i, s := range mdSlides {
		data := blackfriday.MarkdownCommon(s)
		slides = append(slides, &Slide{
			FName:   strings.TrimSuffix(path.Base(fname), path.Ext(fname)),
			CurNo:   i,
			PrevNo:  (i - 1),
			NextNo:  (i + 1),
			FirstNo: 0,
			LastNo:  lastSlide,
			Title:   title,
			Content: string(data),
			CSSPath: cssPath,
			JSPath:  jsPath,
		})
	}
	return slides
}

// MakeSlide this takes a io.Writer, a template and slide and executes the template.
func MakeSlide(wr io.Writer, tmpl *template.Template, slide *Slide) error {
	return tmpl.Execute(wr, slide)
}

// MakeSlideFile this takes a template and slide and renders the results to a file.
func MakeSlideFile(tmpl *template.Template, slide *Slide) error {
	sname := fmt.Sprintf(`%02d-%s.html`, slide.CurNo, strings.TrimSuffix(path.Base(slide.FName), path.Ext(slide.FName)))
	fp, err := os.Create(sname)
	if err != nil {
		return fmt.Errorf("%s %s", sname, err)
	}
	defer fp.Close()
	err = MakeSlide(fp, tmpl, slide)
	if err != nil {
		return fmt.Errorf("%s %s", sname, err)
	}
	return nil
}

// MakeSlideString this takes a template and slide and renders the results to a string
func MakeSlideString(tmpl *template.Template, slide *Slide) (string, error) {
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	err := MakeSlide(wr, tmpl, slide)
	return buf.String(), err
}

// NormalizeDate takes a MySQL like date string and returns a time.Time or error
func NormalizeDate(s string) (time.Time, error) {
	switch len(s) {
	case len(`2006-01-02 15:04:05 -0700`):
		dt, err := time.Parse(`2006-01-02 15:04:05 -0700`, s)
		return dt, err
	case len(`2006-01-02 15:04:05`):
		dt, err := time.Parse(`2006-01-02 15:04:05`, s)
		return dt, err
	case len(`2006-01-02`):
		dt, err := time.Parse(`2006-01-02`, s)
		return dt, err
	default:
		return time.Time{}, fmt.Errorf("Can't format %s, expected format like 2006-01-02 15:04:05 -0700", s)
	}
}

// Walk takes a start path and walks the file system to process Markdown files for useful elements.
func Walk(startPath string, filterFn func(p string, info os.FileInfo) bool, outputFn func(s string, info os.FileInfo) error) error {
	err := filepath.Walk(startPath, func(p string, info os.FileInfo, err error) error {
		// Are we interested in this path?
		if filterFn(p, info) == true {
			// Yes, so send to output function.
			if err := outputFn(p, info); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// Grep looks for the first line matching the expression
// in src.
func Grep(exp string, src string) string {
	re, err := regexp.Compile(exp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q is not a valid, %s\n", exp, err)
		return ""
	}
	lines := strings.Split(src, "\n")
	for _, line := range lines {
		s := re.FindString(line)
		if len(s) > 0 {
			return s
		}
	}
	return ""
}
