package main

import (
	"derailleur/handler"
	pb "derailleur/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "derailleur"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterDerailleurHandler(srv.Server(), new(handler.Derailleur))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
