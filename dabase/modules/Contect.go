package mod

import (
	"context"
	"time"
)

// タイムアウト付きコンテキスト（書き味向上用）
func Ctx(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
