package example

import (
	"os"
	"os/exec"
	"testing"

	"github.com/nokamoto/grpc-tryout/pkg/apis/tryout"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/encoding/protojson"
)

var _ = Describe("protoc-gen-tryout", func() {
	It("should run protoc-gen-tryout", func() {
		cmd := exec.Command("buf", "generate", "--template", "build/buf.gen.yaml")
		cmd.Dir = "../../.."
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
		err := cmd.Run()
		Expect(err).ShouldNot(HaveOccurred())

		bytes, err := os.ReadFile("../../../build/tryout/apis/example/service.pb.json")
		Expect(err).ShouldNot(HaveOccurred())

		var proto tryout.Proto
		err = protojson.Unmarshal(bytes, &proto)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(proto.GetName()).To(Equal("apis/example/service.proto"))
	})
})

func TestProtoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "protoc-gen-tryout Suite")
}
