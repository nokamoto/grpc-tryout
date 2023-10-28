package protogen

import (
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

type Option struct {
	in    io.Reader
	out   io.Writer
	debug io.Writer
}

func NewOption() *Option {
	return &Option{
		in:    os.Stdin,
		out:   os.Stdout,
		debug: io.Discard,
	}
}

// Run is a protoc-gen-tryout plugin that generates a tryout json file.
func (o *Option) Run() error {
	// ref: https://github.com/protocolbuffers/protobuf-go/blob/master/compiler/protogen/protogen.go
	in, err := io.ReadAll(o.in)
	if err != nil {
		return err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return err
	}

	o.setOptions(req.GetParameter())

	o.debugf("GetParameter: %s", req.GetParameter())
	o.debugf("GetFileToGenerate: %s", req.GetFileToGenerate())

	res, err := o.response(req)
	if err != nil {
		return err
	}

	out, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	if _, err := o.out.Write(out); err != nil {
		return err
	}
	return nil
}

func (o *Option) debugf(format string, args ...any) {
	fmt.Fprintf(o.debug, format, args...)
	fmt.Fprintln(o.debug)
}

func (o *Option) setOptions(opts string) {
	for _, opt := range strings.Split(opts, ",") {
		switch {
		case opt == "debug":
			o.debug = os.Stderr
		}
	}
}

func (o *Option) response(req *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	var res pluginpb.CodeGeneratorResponse
	return &res, nil
}
