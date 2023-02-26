package utils

import (
	"go.olapie.com/ctxutil/internal/templates"
	htmlTemplate "html/template"
	textTemplate "text/template"
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
