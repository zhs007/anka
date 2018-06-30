package grpcserv

import (
	"context"
	"net"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhs007/anka/base"
	pb "github.com/zhs007/anka/proto"
)

// server is used to implement netcore.NetCoreServ.
type server struct{}

// join implements netcore.NetCoreServ
func (s *server) Join(ctx context.Context, in *pb.CtrlJoin) (*pb.ReplyCtrl, error) {
	return &pb.ReplyCtrl{Result: 0}, nil
}

// subscribe implements netcore.NetCoreServ
func (s *server) Subscribe(in *pb.Subscribe, ss pb.NetCoreServ_SubscribeServer) error {
	return nil
}

// // updGameStatistics implements guestdb.GuestDBServ
// func (s *server) UpdGameStatistics(ctx context.Context, in *pb.UpdGameStatistics) (*pb.ReplyCtrl, error) {
// 	base.Info("UpdGameStatistics")
// 	err := model.UpdGameStatistics(in.Gamecode, in.Gamemodcode, int64(in.Playnums), int64(in.Totalpay), int64(in.Totalwin))

// 	return &pb.ReplyCtrl{Code: 0}, err
// }

// StartServ -
func StartServ(servaddr string, wg *sync.WaitGroup) {
	base.Info("grpcserv start...", zap.String("servaddr", servaddr))

	lis, err := net.Listen("tcp", servaddr)
	if err != nil {
		base.Fatal("failed to listen:", zap.Error(err))
		// log.Fatalf("failed to listen: %v", err)
	}

	// base.Info("listen end")

	s := grpc.NewServer()
	pb.RegisterNetCoreServServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// base.Info("Register end")

	if err := s.Serve(lis); err != nil {
		base.Fatal("failed to serve:", zap.Error(err))
	}

	base.Info("grpcserv end.")
	wg.Done()
}
