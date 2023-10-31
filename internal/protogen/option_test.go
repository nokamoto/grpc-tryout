package protogen

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/grpc-tryout/pkg/apis/tryout"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestOption_Run(t *testing.T) {
	tests := []struct {
		name string
		in   *pluginpb.CodeGeneratorRequest
		want map[string]*tryout.Proto
	}{
		{
			name: "no service definition",
			in: &pluginpb.CodeGeneratorRequest{
				ProtoFile: []*descriptorpb.FileDescriptorProto{
					{
						Name: proto.String("apis/example/resource.proto"),
					},
				},
			},
			want: map[string]*tryout.Proto{},
		},
		{
			name: "service definition",
			in: &pluginpb.CodeGeneratorRequest{
				ProtoFile: []*descriptorpb.FileDescriptorProto{
					{
						Name:    proto.String("apis/example/resource.proto"),
						Package: proto.String("example"),
						MessageType: []*descriptorpb.DescriptorProto{
							{
								Name: proto.String("GetShelfRequest"),
								Field: []*descriptorpb.FieldDescriptorProto{
									{
										Name:     proto.String("name"),
										JsonName: proto.String("name"),
									},
								},
							},
						},
					},
					{
						Name:    proto.String("apis/example/service.proto"),
						Package: proto.String("example"),
						Service: []*descriptorpb.ServiceDescriptorProto{
							{
								Name: proto.String("Library"),
								Method: []*descriptorpb.MethodDescriptorProto{
									{
										Name:      proto.String("GetShelf"),
										InputType: proto.String(".example.GetShelfRequest"),
									},
								},
							},
						},
					},
				},
			},
			want: map[string]*tryout.Proto{
				"apis/example/service.pb.json": {
					Name: "apis/example/service.proto",
					Services: []*tryout.Service{
						{
							Name: "Library",
							Methods: []*tryout.Method{
								{
									Name:   "GetShelf",
									Path:   "/example.Library/GetShelf",
									Fields: []string{"name"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := proto.Marshal(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			var out bytes.Buffer
			opt := Option{
				in:    bytes.NewBuffer(b),
				out:   &out,
				debug: io.Discard,
			}

			if err := opt.Run(); err != nil {
				t.Errorf("Run() error = %v", err)
			}

			var res pluginpb.CodeGeneratorResponse
			if err := proto.Unmarshal(out.Bytes(), &res); err != nil {
				t.Fatal(err)
			}

			got := map[string]*tryout.Proto{}
			for _, f := range res.GetFile() {
				var p tryout.Proto
				if err := protojson.Unmarshal([]byte(f.GetContent()), &p); err != nil {
					t.Fatalf("unmarshal error: %s: %v", f.GetName(), err)
				}
				got[f.GetName()] = &p
			}

			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
