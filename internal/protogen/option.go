package protogen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nokamoto/grpc-tryout/pkg/apis/tryout"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Option struct {
	in        io.Reader
	out       io.Writer
	debug     io.Writer
	multiline bool
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

func debugf(indent int, w io.Writer, format string, args ...any) {
	fmt.Fprint(w, strings.Repeat(" ", indent))
	fmt.Fprintf(w, format, args...)
	fmt.Fprintln(w)
}

func (o *Option) debugf(format string, args ...any) {
	debugf(0, o.debug, format, args...)
}

func (o *Option) setOptions(opts string) {
	for _, opt := range strings.Split(opts, ",") {
		switch {
		case opt == "debug":
			o.debug = os.Stderr
		case opt == "multiline":
			o.multiline = true
		}
	}
}

func (o *Option) response(req *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	var res pluginpb.CodeGeneratorResponse
	for _, file := range req.GetProtoFile() {
		o.debugf("proto: %s", file.GetName())
		content, err := fileDesc{
			req:   req,
			desc:  file,
			debug: o.debug,
		}.proto()
		if err != nil {
			return nil, err
		}
		if content == nil {
			continue
		}

		f, err := o.responseFile(file, content)
		if err != nil {
			return nil, err
		}
		res.File = append(res.File, f)

		o.debugf("generated: %s", f.GetName())
	}
	return &res, nil
}

func (o *Option) responseFile(
	file *descriptorpb.FileDescriptorProto,
	content *tryout.Proto,
) (*pluginpb.CodeGeneratorResponse_File, error) {
	m := protojson.MarshalOptions{
		Multiline: o.multiline,
	}
	bytes, err := m.Marshal(content)
	if err != nil {
		return nil, err
	}

	filename := fmt.Sprintf(
		"%s/%s.pb.json",
		filepath.Dir(file.GetName()),
		strings.TrimSuffix(filepath.Base(file.GetName()), filepath.Ext(file.GetName())),
	)

	var buf strings.Builder
	buf.Write(bytes)
	fmt.Fprintln(&buf)

	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(filename),
		Content: proto.String(buf.String()),
	}, nil
}
