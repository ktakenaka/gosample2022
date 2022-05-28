package logger

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type testWS struct {
	*buffer.Buffer
}

func (w *testWS) Sync() error {
	return nil
}

const want = `{"level":"info","ts":"2006-01-02T15:04:05Z07:00","msg":"info","key":"value"}
{"level":"warn","ts":"2006-01-02T15:04:05Z07:00","msg":"warn","key":"value"}
{"level":"error","ts":"2006-01-02T15:04:05Z07:00","msg":"error","key":"value"}
`

func TestNew(t *testing.T) {
	ws := &testWS{Buffer: buffer.NewPool().Get()}

	cfg := &Config{
		WS: ws,
		TimeEncoder: func(_ time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(time.RFC3339)
		},
	}
	log := New(cfg)
	log.Infow("info", "key", "value")
	log.Warnw("warn", "key", "value")
	log.Errorw("error", "key", "value")
	log.Sync()

	if diff := cmp.Diff(ws.String(), want); diff != "" {
		t.Errorf("diff %s", diff)
	}
}
