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
)

var _ = BeforeSuite(func() {
	go example.Run()
})

var _ = Describe("Library", Ordered, func() {
	var client exampleconnect.LibraryClient

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
})

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Suite")
}
