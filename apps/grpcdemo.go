package apps

import (
	"github.com/labstack/echo"
	pb "github.com/sunisdown/gqs/services/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func GRPCDEMO(c echo.Context) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		c.Logger().Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	name := c.QueryParam("name")

	con := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := con.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		c.Logger().Fatalf("could not greet: %v", err)
	}
	return c.String(http.StatusOK, r.Message)
}
