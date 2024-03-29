package utils

import (
	"context"
	"crypto/rand"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

const TimeDay = time.Hour * 24
const TimeWeek = TimeDay * 7
const TimeYear = TimeDay * 52

func StringPtr(v string) *string { return &v }

func IntPtr(v int) *int { return &v }

func TimePtr(v time.Time) *time.Time { return &v }

func IntArrayPtr(v []int) *[]int { return &v }

// RandomString generates a random string, and panics if there is an error.
// https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func RandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

const (
	layoutISO = "2006-01-02"
)

func MustParseTime(t *testing.T, toParse string) time.Time {
	t.Helper()

	parsed, err := time.Parse(layoutISO, toParse)
	if err != nil {
		t.Fatal(err)
	}

	return parsed
}

func AddParamToContext(t *testing.T, req *http.Request, key, value string) {
	t.Helper()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	*req = *req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
}
