package example

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/nokamoto/grpc-tryout/internal/service/example"
	"github.com/nokamoto/grpc-tryout/pkg/apis/example/exampleconnect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Address() string {
	return "localhost:8080"
}

func Run() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})))
	mux := http.NewServeMux()
	path, handler := exampleconnect.NewLibraryHandler(example.NewLibraryService())
	mux.Handle(path, handler)
	handler = cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	slog.Info("listen and serve", slog.String("address", Address()), slog.String("path", path))
	if err := http.ListenAndServe(Address(), handler); err != nil {
		slog.Info("halt with error", slog.String("error", err.Error()))
	}
}
