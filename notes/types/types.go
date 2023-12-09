package types

type TemplateContextKey string

var NotesContextKey TemplateContextKey

type Note struct {
	Id      string
	Title   string
	Content string
}
