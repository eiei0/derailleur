package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "derailleur/proto"
)

type Derailleur struct{}

func (e *Derailleur) CreateDerailleur(ctx context.Context, req *pb.CreateDerailleurRequest, rsp *pb.CreateDerailleurResponse) error {
	log.Infof("Received Derailleur.Call request: %v", req)

	rsp.Derailleur = &pb.BikeDerailleur{
		Id:           "urn:bicycle:derailleur:1",
		FrameId:      req.FrameId,
		Manufacturer: req.Manufacturer,
		Speed:        req.Speed,
		Model:        req.Model,
	}

	return nil
}
