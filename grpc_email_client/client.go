package grpc_email_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	email "restApiProject/grpc_email_server"
	"restApiProject/models"
)

var GrpcAddr string

func Client(order models.Order) error {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %s", err)
	}
	defer conn.Close()

	c := email.NewEmailServiceClient(conn)

	_, err = c.Send(context.Background(),
		&email.Request{OrderId: int64(order.Id), ItemId: int64(order.ItemId), Email: order.Email})
	if err != nil {
		return fmt.Errorf("error when calling Send: %s", err)
	}
	return nil
}
