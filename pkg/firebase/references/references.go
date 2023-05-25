package references

import (
	"context"

	"firebase.google.com/go/db"
)

// Creator generates a new db.Ref from a path, permitting operations
// on the reference.
type Creator interface {
	NewRef(path string) Operator
}

// Operator specifies the contract for db.Ref so that it can be mocked.
type Operator interface {
	OrderByChild(child string) *db.Query
	OrderByKey() *db.Query
	OrderByValue() *db.Query
	Parent() *db.Ref
	Child(path string) *db.Ref
	Get(ctx context.Context, v interface{}) error
	GetWithETag(ctx context.Context, v interface{}) (string, error)
	GetShallow(ctx context.Context, v interface{}) error
	GetIfChanged(ctx context.Context, etag string, v interface{}) (
		bool,
		string,
		error,
	)
	Set(ctx context.Context, v interface{}) error
	SetIfUnchanged(ctx context.Context, etag string, v interface{}) (bool, error)
	Push(ctx context.Context, v interface{}) (*db.Ref, error)
	Update(ctx context.Context, v map[string]interface{}) error
	Transaction(ctx context.Context, fn db.UpdateFn) error
	Delete(ctx context.Context) error
}
