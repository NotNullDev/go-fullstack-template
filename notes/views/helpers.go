package views

import (
	"context"
	"fmt"

	"github.com/notnulldev/go-templ-playground/notes/types"
)

func GetNotes(ctx context.Context) []types.Note {
	notes, ok := ctx.Value(types.NotesContextKey).([]types.Note)
	if !ok {
		println("failed to get notes from context")
		return []types.Note{}
	}
	return notes
}

func GetNotesLengthString(ctx context.Context) string {
	return fmt.Sprintf("Found %d notes!", len(GetNotes(ctx)))
}
