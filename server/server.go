package server

import (
	"context"

	api "github.com/bominion/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
type Config struct {
	CommitLog CommitLog
}
*/

var _ api.BominionServer = (*grpcServer)(nil)

type grpcServer struct {
	api.UnimplementedBominionServer
	//*Config
}

/*
type CommitLog interface {
	Append(*api.Record) (uint64, error)
	Read(uint64) (*api.Record, error)
}
*/
func NewGRPCServer() (*grpc.Server, error) {
	gsrv := grpc.NewServer()
	//TODO:テスト用 リフレクションの設定
	reflection.Register(gsrv)
	srv, err := newgrpcServer()
	if err != nil {
		return nil, err
	}
	//ここでサービス登録を行っている。
	api.RegisterBominionServer(gsrv, srv)
	return gsrv, nil
}

func newgrpcServer() (srv *grpcServer, err error) {
	srv = &grpcServer{}
	return srv, nil
}

func (s *grpcServer) Buy(ctx context.Context, req *api.BuyRequest) (
	*api.BuyResponse, error) {

	//sys.Players[sys.WhosTurn].BuyCard(sys.Supply[sys.Players[sys.WhosTurn].Pointer])
	return &api.BuyResponse{Done: true}, nil
}

/*

func (s *grpcServer) Consume(ctx context.Context, req *api.ConsumeRequest) (
	*api.ConsumeResponse, error) {
	record, err := s.CommitLog.Read(req.Offset)
	if err != nil {
		return nil, err
	}
	return &api.ConsumeResponse{Record: record}, nil
}

func (s *grpcServer) ProduceStream(
	stream api.Log_ProduceStreamServer,
) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		res, err := s.Produce(stream.Context(), req)
		if err != nil {
			return err
		}
		if err = stream.Send(res); err != nil {
			return err
		}
	}
}

func (s *grpcServer) ConsumeStream(
	req *api.ConsumeRequest,
	stream api.Log_ConsumeStreamServer,
) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			res, err := s.Consume(stream.Context(), req)
			switch err.(type) {
			case nil:
			case api.ErrOffsetOutOfRange:
				continue
			default:
				return err
			}
			if err = stream.Send(res); err != nil {
				return err
			}
			req.Offset++
		}
	}
}
*/
