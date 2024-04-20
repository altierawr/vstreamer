package hooks

import (
	"context"

	"github.com/altierawr/vstreamer/ent"
	gen "github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ent/hook"
)

func RegisterPlaySessionHooks(ctx context.Context, client *ent.Client) {
	client.PlaySession.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PlaySessionFunc(func(ctx context.Context, m *gen.PlaySessionMutation) (ent.Value, error) {
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	)
}
