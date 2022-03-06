package tracing

import (
	"context"

	misc "github.com/Pavel7004/Common/misc"
	opentracing "github.com/opentracing/opentracing-go"
)

func StartSpanFromContext(ctx context.Context) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, misc.GetParentFuncName())
}
