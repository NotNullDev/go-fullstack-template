package notes

import (
	"context"
	"fmt"
)

type TemplateContextKey string

var NotesContextKey TemplateContextKey

type Note struct {
	Id      string
	Title   string
	Content string
}

func GetNotes(ctx context.Context) []Note {
	notes, ok := ctx.Value(NotesContextKey).([]Note)
	if !ok {
		println("failed to get notes from context")
		return []Note{}
	}
	return notes
}

func GetNotesLengthString(ctx context.Context) string {
	return fmt.Sprintf("Found %d notes!", len(GetNotes(ctx)))
}
