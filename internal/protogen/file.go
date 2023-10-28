package protogen

import (
	"fmt"
	"io"

	"github.com/nokamoto/grpc-tryout/pkg/apis/tryout"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type fileDesc struct {
	req   *pluginpb.CodeGeneratorRequest
	desc  *descriptorpb.FileDescriptorProto
	debug io.Writer
}

func (d fileDesc) debugf(format string, args ...any) {
	debugf(2, d.debug, format, args...)
}

func (d fileDesc) message(typ string) (*descriptorpb.DescriptorProto, error) {
	for _, f := range d.req.GetProtoFile() {
		for _, m := range f.GetMessageType() {
			name := fmt.Sprintf(".%s.%s", f.GetPackage(), m.GetName())
			if name == typ {
				return m, nil
			}
		}
	}
	return nil, fmt.Errorf("not found: %s", typ)
}

func (d fileDesc) proto() (*tryout.Proto, error) {
	var services []*tryout.Service
	for _, s := range d.desc.GetService() {
		d.debugf("service: %s", s.GetName())
		service, err := serviceDesc{
			fileDesc: d,
			desc:     s,
		}.service()
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	if len(services) == 0 {
		return nil, nil
	}
	return &tryout.Proto{
		Name:     d.desc.GetName(),
		Services: services,
	}, nil
}

type serviceDesc struct {
	fileDesc
	desc *descriptorpb.ServiceDescriptorProto
}

func (d serviceDesc) debugf(format string, args ...any) {
	debugf(4, d.debug, format, args...)
}

func (d serviceDesc) service() (*tryout.Service, error) {
	var methods []*tryout.Method
	for _, m := range d.desc.GetMethod() {
		d.debugf("method: %s(%s)", m.GetName(), m.GetInputType())
		method, err := methodDesc{
			serviceDesc: d,
			desc:        m,
		}.method()
		if err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}
	return &tryout.Service{
		Name:    d.desc.GetName(),
		Methods: methods,
	}, nil
}

type methodDesc struct {
	serviceDesc
	desc *descriptorpb.MethodDescriptorProto
}

func (d methodDesc) debugf(format string, args ...any) {
	debugf(6, d.debug, format, args...)
}

func (d methodDesc) method() (*tryout.Method, error) {
	inputType, err := d.message(d.desc.GetInputType())
	if err != nil {
		return nil, err
	}
	var fields []string
	for _, field := range inputType.GetField() {
		d.debugf("field: %s %s %s", field.GetName(), field.GetType(), field.GetTypeName())
		fields = append(fields, field.GetJsonName())
	}
	return &tryout.Method{
		Name:   d.desc.GetName(),
		Fields: fields,
	}, nil
}
