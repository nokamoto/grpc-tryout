package example

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/nokamoto/grpc-tryout/internal/server/example"
	examplepb "github.com/nokamoto/grpc-tryout/pkg/apis/example"
	"github.com/nokamoto/grpc-tryout/pkg/apis/example/exampleconnect"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/proto"
)

var _ = BeforeSuite(func() {
	go example.Run()
})

var _ = Describe("Library", Ordered, func() {
	var client exampleconnect.LibraryClient
	var created *examplepb.Shelf

	It("should connect with HTTP", func(ctx SpecContext) {
		client = exampleconnect.NewLibraryClient(
			http.DefaultClient,
			fmt.Sprintf("http://%s", example.Address()),
		)
		Eventually(func() connect.Code {
			_, err := client.GetShelf(ctx, connect.NewRequest(&examplepb.GetShelfRequest{}))
			return connect.CodeOf(err)
		}).WithContext(ctx).Should(Equal(connect.CodeNotFound))
	}, SpecTimeout(5*time.Second))

	It("should connect with gRPC", func(ctx SpecContext) {
		client = exampleconnect.NewLibraryClient(
			http.DefaultClient,
			fmt.Sprintf("http://%s", example.Address()),
			connect.WithGRPC(),
		)
		Eventually(func() connect.Code {
			_, err := client.GetShelf(ctx, connect.NewRequest(&examplepb.GetShelfRequest{}))
			return connect.CodeOf(err)
		}).WithContext(ctx).Should(Equal(connect.CodeNotFound))
	}, SpecTimeout(5*time.Second))

	It("should create a shelf", func(ctx SpecContext) {
		shelf := &examplepb.Shelf{
			Category: examplepb.Category_SCIENCE,
		}
		res, err := client.CreateShelf(ctx, connect.NewRequest(&examplepb.CreateShelfRequest{
			Shelf: shelf,
		}))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(res.Msg.GetCategory()).To(Equal(shelf.GetCategory()))
		Expect(res.Msg.GetName()).ToNot(BeEmpty())
		created = res.Msg
	})

	It("should get a shelf", func(ctx SpecContext) {
		res, err := client.GetShelf(ctx, connect.NewRequest(&examplepb.GetShelfRequest{
			Name: created.GetName(),
		}))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(proto.Equal(created, res.Msg)).To(BeTrue())
	})

	It("should delete a shelf", func(ctx SpecContext) {
		_, err := client.DeleteShelf(ctx, connect.NewRequest(&examplepb.DeleteShelfRequest{
			Name: created.GetName(),
		}))
		Expect(err).ShouldNot(HaveOccurred())
	})
})

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Suite")
}
