package firebase

import (
	"firebase.google.com/go/db"
)

// ReferenceCreator generates a new db.Ref from a path, permitting operations
// on the reference.
type ReferenceCreator interface {
	NewRef(path string) *db.Ref
}
