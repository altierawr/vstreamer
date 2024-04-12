// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/altierawr/vstreamer/ent/library"
	"github.com/altierawr/vstreamer/ent/schema"
	"github.com/altierawr/vstreamer/ent/video"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	libraryHooks := schema.Library{}.Hooks()
	library.Hooks[0] = libraryHooks[0]
	libraryFields := schema.Library{}.Fields()
	_ = libraryFields
	// libraryDescCreatedAt is the schema descriptor for created_at field.
	libraryDescCreatedAt := libraryFields[1].Descriptor()
	// library.DefaultCreatedAt holds the default value on creation for the created_at field.
	library.DefaultCreatedAt = libraryDescCreatedAt.Default.(func() time.Time)
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescCreatedAt is the schema descriptor for created_at field.
	videoDescCreatedAt := videoFields[1].Descriptor()
	// video.DefaultCreatedAt holds the default value on creation for the created_at field.
	video.DefaultCreatedAt = videoDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
