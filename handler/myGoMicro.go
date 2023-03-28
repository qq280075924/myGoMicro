package handler

import (
	"context"
	first "first/proto"
	"go-micro.dev/v4"
	"io"
	"time"

	"go-micro.dev/v4/logger"

	pb "myGoMicro/proto"
)

type MyGoMicro struct{}

func (e *MyGoMicro) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received MyGoMicro.Call request: %v", req)
	//rsp.Msg = "Hello " + req.Name
	service := micro.NewService()
	service.Init()
	firstService := first.NewFirstService("first", service.Client())
	call, _ := firstService.Call(context.Background(), &first.CallRequest{
		Name: req.Name,
	})
	rsp.Msg = call.Msg
	return nil
}

func (e *MyGoMicro) ClientStream(ctx context.Context, stream pb.MyGoMicro_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *MyGoMicro) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.MyGoMicro_ServerStreamStream) error {
	logger.Infof("Received MyGoMicro.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *MyGoMicro) BidiStream(ctx context.Context, stream pb.MyGoMicro_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
