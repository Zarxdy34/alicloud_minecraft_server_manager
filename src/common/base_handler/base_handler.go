package base_handler

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

type BaseHandler struct {
	ctx  context.Context
	name string
}

func New(ctx context.Context, name string) *BaseHandler {
	return &BaseHandler{
		ctx:  ctx,
		name: name,
	}
}

func (h *BaseHandler) GetContext() context.Context {
	return h.ctx
}

func (h *BaseHandler) GetName() string {
	return h.name
}

func (h *BaseHandler) LogInfo(format string, args ...interface{}) {
	logger.CtxInfof(h.ctx, format, args)
}

func (h *BaseHandler) LogWarn(format string, args ...interface{}) {
	logger.CtxInfof(h.ctx, format, args)
}

func (h *BaseHandler) LogError(format string, args ...interface{}) {
	logger.CtxInfof(h.ctx, format, args)
}

func (h *BaseHandler) LogDebug(format string, args ...interface{}) {
	logger.CtxInfof(h.ctx, format, args)
}
