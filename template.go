package utils

import (
	htmlTemplate "html/template"
	textTemplate "text/template"

	"go.olapie.com/utils/internal/templates"
)

var HTMLTemplateFuncMap = htmlTemplate.FuncMap{
	"plus":       templates.Plus,
	"minus":      templates.Minus,
	"multiple":   templates.Multiple,
	"divide":     templates.Divide,
	"join":       templates.Join,
	"lower":      templates.ToLower,
	"upper":      templates.ToUpper,
	"concat":     templates.Concat,
	"capitalize": templates.Capitalize,
}

var TextTemplateFuncMap = textTemplate.FuncMap{
	"plus":       templates.Plus,
	"minus":      templates.Minus,
	"multiple":   templates.Multiple,
	"divide":     templates.Divide,
	"join":       templates.Join,
	"lower":      templates.ToLower,
	"upper":      templates.ToUpper,
	"concat":     templates.Concat,
	"capitalize": templates.Capitalize,
}
