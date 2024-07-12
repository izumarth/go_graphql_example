package graph

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/izumarth/go-graphql-example/internal"
	"github.com/izumarth/go-graphql-example/middlewares/auth"
)

var Directive internal.DirectiveRoot = internal.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

const spicedbEndpoint = "grpc.authzed.com:443"

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	authzedToken := os.Getenv("authzedToken")

	name, ok := auth.GetUserName(ctx)
	if !ok {
		return nil, errors.New("not authencated")
	}

	systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
	if err != nil {
		log.Fatalf("unable to load system CA certificates: %s", err)
	}

	client, err := authzed.NewClient(
		spicedbEndpoint,
		grpcutil.WithBearerToken(authzedToken),
		systemCerts,
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	ctxBackGround := context.Background()

	user := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: "izumarth/user",
		ObjectId:   name,
	}}

	firstPost := &pb.ObjectReference{
		ObjectType: "izumarth/repo",
		ObjectId:   "1",
	}

	resp, err := client.CheckPermission(ctxBackGround, &pb.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "write",
		Subject:    user,
	})
	if err != nil {
		return nil, errors.New("not authencated")
	}

	if resp.Permissionship != pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		return nil, errors.New("not authencated")
	}

	return next(ctx)
}
