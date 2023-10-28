package example

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/nokamoto/grpc-tryout/pkg/apis/example"
)

type LibraryService struct {
	memory map[string]*example.Shelf
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		memory: map[string]*example.Shelf{},
	}
}

func (l *LibraryService) GetShelf(
	_ context.Context,
	req *connect.Request[example.GetShelfRequest],
) (*connect.Response[example.Shelf], error) {
	msg := req.Msg
	slog.Debug("", slog.Any("msg", msg))
	if shelf, ok := l.memory[msg.GetName()]; ok {
		return &connect.Response[example.Shelf]{
			Msg: shelf,
		}, nil
	}
	return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("%s not found", msg.GetName()))
}
