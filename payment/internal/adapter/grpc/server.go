package grpcadapter

import (
	"context"
	"fmt"

	paymentpb "github.com/nillocoelho/microservices-proto/golang/payment"
	"github.com/nillocoelho/microservices/payment/internal/application/core/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	paymentpb.UnimplementedPaymentServer
	app *api.Application
}

func NewServer(app *api.Application) *Server {
	return &Server{app: app}
}

func (s *Server) Create(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	p, b, err := s.app.CreatePayment(req.UserId, req.OrderId, req.TotalPrice)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return nil, err
		}

		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v", err)).Err()
	}

	return &paymentpb.CreatePaymentResponse{
		PaymentId: int64(p.ID),
		BillId:    int64(b.ID),
	}, nil
}
