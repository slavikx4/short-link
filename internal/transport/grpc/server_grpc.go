package grpc

import (
	"context"
	"github.com/slavikx4/short-link/internal/services"
	"github.com/slavikx4/short-link/pkg/api/proto"
	"github.com/slavikx4/short-link/pkg/logger"
)

type ServerShortLinkGRPC struct {
	proto.UnimplementedShortLinkServer
	Service *services.Service
}

func (s *ServerShortLinkGRPC) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	logger.Logger.Process.Println("пришёл запрос GET: ", req.GetShortLink())
	link, err := s.Service.GetLink(req.GetShortLink())
	if err != nil {
		return nil, err
	}
	response := &proto.GetResponse{OriginalLink: link.OriginalLink}
	return response, nil
}

func (s *ServerShortLinkGRPC) Post(ctx context.Context, req *proto.PostRequest) (*proto.PostResponse, error) {
	logger.Logger.Process.Println("пришёл запрос POST: ", req.GetOriginLink())
	link, err := s.Service.SetLink(req.GetOriginLink())
	if err != nil {
		return nil, err
	}
	response := &proto.PostResponse{ShortLink: link.ShortLink}
	return response, nil
}
