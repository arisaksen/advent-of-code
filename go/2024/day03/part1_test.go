package main

import (
	_ "embed"
	"github.com/lmittmann/tint"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestPart1(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)

	// 161 (2*4 + 5*5 + 11*8 + 8*5).
	assert.Equal(t, 161, part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"))
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	}
	b.ReportAllocs()
}
