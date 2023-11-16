package cloud

import (
	"context"

	"github.com/khulnasoft-lab/tunnel-aws/internal/adapters/cloud/aws"
	"github.com/khulnasoft-lab/tunnel-aws/internal/adapters/cloud/options"
	"github.com/khulnasoft/defsec/pkg/state"
)

// Adapt ...
func Adapt(ctx context.Context, opt options.Options) (*state.State, error) {
	cloudState := &state.State{}
	err := aws.Adapt(ctx, cloudState, opt)
	return cloudState, err
}
