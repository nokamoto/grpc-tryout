package example

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"connectrpc.com/connect"
	"github.com/nokamoto/grpc-tryout/pkg/apis/example"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LibraryService struct {
	memory map[string]*example.Shelf
	random *rand.Rand
}

func (l *LibraryService) generateName() (string, error) {
	left := []string{"red", "green", "blue", "yellow", "black", "white"}
	middle := []string{"small", "medium", "large"}
	right := []string{"cat", "dog", "bird", "fish", "rabbit", "turtle"}
	r := func(v []string) string {
		return v[l.random.Intn(len(v))]
	}
	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("%s-%s-%s", r(left), r(middle), r(right))
		if _, ok := l.memory[name]; !ok {
			return name, nil
		}
	}
	return "", fmt.Errorf("failed to generate a random name")
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		memory: map[string]*example.Shelf{},
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (l *LibraryService) GetShelf(
	_ context.Context,
	req *connect.Request[example.GetShelfRequest],
) (*connect.Response[example.Shelf], error) {
	msg := req.Msg
	slog.Debug("", slog.Any("msg", msg))
	if shelf, ok := l.memory[msg.GetName()]; ok {
		return connect.NewResponse(shelf), nil
	}
	return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("%s not found", msg.GetName()))
}

func (l *LibraryService) CreateShelf(
	_ context.Context,
	req *connect.Request[example.CreateShelfRequest],
) (*connect.Response[example.Shelf], error) {
	msg := req.Msg.GetShelf()
	slog.Debug("", slog.Any("msg", msg))

	if msg.GetCategory() == example.Category_UNSPECIFIED {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("category must be specified"))
	}

	name, err := l.generateName()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var shelf example.Shelf
	proto.Merge(&shelf, msg)
	shelf.Name = name
	l.memory[name] = &shelf
	return connect.NewResponse(&shelf), nil
}

func (l *LibraryService) DeleteShelf(
	_ context.Context,
	req *connect.Request[example.DeleteShelfRequest],
) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	slog.Debug("", slog.Any("msg", msg))
	delete(l.memory, msg.GetName())
	return connect.NewResponse(&emptypb.Empty{}), nil
}
