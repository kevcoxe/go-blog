package server

import (
	"bytes"
	"errors"
	"text/template"
)

type Element interface {
	Render() (string, error)
}

type RenderedPageData struct {
	Elements []string
}

type PageData struct {
	Elements []Element
}

func (pd *PageData) Render() (RenderedPageData, error) {
	var data RenderedPageData
	for _, element := range pd.Elements {
		rendered, err := element.Render()
		if err != nil {
			return RenderedPageData{}, err
		}
		data.Elements = append(data.Elements, rendered)
	}

	return data, nil
}

type FormElement struct {
	Label string
	Type  string
	Name  string
	Value string
}

type InputForm struct {
	Title       string
	Description string
	Elements    []FormElement
	Action      string
	Method      string
}

func (inputForm *InputForm) Render() (string, error) {

	var tpl bytes.Buffer
	var formTemplate *template.Template = template.Must(template.ParseFiles("cmd/public/form.html"))

	fromErr := formTemplate.Execute(&tpl, inputForm)
	if fromErr != nil {
		return "", errors.New("cannot render form")
	}

	return tpl.String(), nil
}

type Post struct {
	ID      int
	Title   string
	Content string
}

func (p *Post) Render() (string, error) {
	var tpl bytes.Buffer
	var postTemplate *template.Template = template.Must(template.ParseFiles("cmd/public/post.html"))

	fromErr := postTemplate.Execute(&tpl, p)
	if fromErr != nil {
		return "", errors.New("cannot render post")
	}

	return tpl.String(), nil
}
