package models

type TemplateData struct {
	StringMap map[string]string
	CSRFToken string
	Data      any
}
