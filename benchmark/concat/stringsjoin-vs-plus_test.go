package concat

import (
	"strings"
	"testing"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func BenchmarkStringsJoin(b *testing.B) {
	var result string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		urlArray := []string{CONN_HOST, ":", CONN_PORT}
		result = strings.Join(urlArray, "")
	}
	result += "done"
}

func BenchmarkPlus(b *testing.B) {
	var result string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = CONN_HOST + ":" + CONN_PORT

	}
	result += "done"
}
