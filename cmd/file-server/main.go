package main

import (
	"fmt"
	"github.com/MadJlzz/gopypi/pkg/listing"
	"github.com/MadJlzz/gopypi/pkg/storage/disk"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	storage := disk.NewStorage(sugar, disk.WithPath("/tmp/local-PyPI"))
	fmt.Println(storage)

	service := listing.NewService(sugar, storage)
	fmt.Println(service.GetAllPackages())
}
