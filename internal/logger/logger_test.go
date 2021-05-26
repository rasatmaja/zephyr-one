package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", "unknown-level")
		loggerTest := New()
		if instance == nil {
			t.Errorf("expecting instance not nil, got %v", loggerTest)
			t.Fail()
		}
	})

}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
