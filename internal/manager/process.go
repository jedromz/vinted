package manager

import "context"

type Process interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
