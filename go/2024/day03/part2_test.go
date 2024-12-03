package main

import (
	_ "embed"
	"github.com/lmittmann/tint"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMatchToDigits(t *testing.T) {

	if digit0, digit1, err := matchToDigits("mul(2,4)"); err != nil {
		t.Fatal(err)
	} else {
		assert.Equal(t, 2, digit0)
		assert.Equal(t, 4, digit1)
	}

	if digit0, digit1, err := matchToDigits("mul(11,8)"); err != nil {
		t.Fatal(err)
	} else {
		assert.Equal(t, 11, digit0)
		assert.Equal(t, 8, digit1)
	}

	for _, match := range []string{"do()", "don't()", "do_not_", "then"} {
		if _, _, err := matchToDigits(match); err != nil {
			assert.Error(t, strconv.ErrSyntax, err)
		} else {
			t.Error("this should result in error and not get here")
		}
	}

}

func TestPart2(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)

	//48 (2*4 + 8*5)
	assert.Equal(t, 48, part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part2("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	}
	b.ReportAllocs()
}
