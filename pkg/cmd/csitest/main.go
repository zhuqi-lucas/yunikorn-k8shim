package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

// FakeCSIPlugin represents our custom fake CSI plugin.
type FakeCSIPlugin struct {
	csi.UnimplementedControllerServer
	csi.UnimplementedIdentityServer // Add this to implement identity methods
}

func main() {

	log.Default().Printf("Starting csi driver server for yunikorn e2e testing")
	// Set up gRPC server
	listener, err := net.Listen("tcp", ":9890")
	if err != nil {
		log.Fatal("Failed to listen on port 9890: %v", err)
	}

	server := grpc.NewServer()
	csi.RegisterControllerServer(server, &FakeCSIPlugin{})
	csi.RegisterIdentityServer(server, &FakeCSIPlugin{}) // Register the Identity Server

	log.Default().Printf("Fake CSI Plugin is running on port 9890...")
	if err := server.Serve(listener); err != nil {
		log.Fatal("Failed to serve gRPC server: %v", err)
	}

	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		close(done)
	}()
	<-done
	server.Stop()
}

// Implement GetPluginInfo to return the driver name
func (p *FakeCSIPlugin) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          "csi.fake.plugin", // This name should match the provisioner name in the StorageClass
		VendorVersion: "1.0.0",
	}, nil
}

// Implement other CSI methods as needed
func (p *FakeCSIPlugin) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      "test-volume-id",
			CapacityBytes: req.CapacityRange.RequiredBytes,
			VolumeContext: req.Parameters,
		},
	}, nil
}

func (p *FakeCSIPlugin) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	// Simulate a binding error.
	return nil, fmt.Errorf("simulated bind failure: unable to bind volume to node")
}
