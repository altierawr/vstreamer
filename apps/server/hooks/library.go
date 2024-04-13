package hooks

import (
	"context"
	"fmt"

	"github.com/altierawr/vstreamer/ent"
	gen "github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ent/hook"
	"github.com/altierawr/vstreamer/features"
)

func RegisterLibraryHooks(ctx context.Context, client *ent.Client) {
	client.Library.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.LibraryFunc(func(ctx context.Context, m *gen.LibraryMutation) (ent.Value, error) {
					defer func() {
						id, exists := m.ID()

						if !exists {
							fmt.Printf("id doesn't exist\n")
							return
						}

						features.ScanLibrary(ctx, client, id)
					}()

					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	)
}
