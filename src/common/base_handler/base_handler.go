package base_handler

import (
	"context"
	"fmt"

	"github.com/google/logger"
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
	logger.InfoDepth(1, fmt.Sprintf("["+h.GetName()+"]"+format, args...))
}

func (h *BaseHandler) LogWarn(format string, args ...interface{}) {
	logger.WarningDepth(1, fmt.Sprintf("["+h.GetName()+"]"+format, args...))
}

func (h *BaseHandler) LogError(format string, args ...interface{}) {
	logger.ErrorDepth(1, fmt.Sprintf("["+h.GetName()+"]"+format, args...))
}

func (h *BaseHandler) LogFatal(format string, args ...interface{}) {
	logger.FatalDepth(1, fmt.Sprintf("["+h.GetName()+"]"+format, args...))
}
