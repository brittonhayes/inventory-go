package schema

import (
	"context"
	"os"

	"entgo.io/ent"
	"github.com/go-kratos/kratos/v2/log"
)

var LogHook = func(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		logger := log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"service.name", "database",
		)

		defer func() {
			logger.Log(log.LevelInfo, m.Op().String())
		}()
		return next.Mutate(ctx, m)
	})
}
