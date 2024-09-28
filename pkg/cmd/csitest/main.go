package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

// FakeCSIPlugin represents our custom fake CSI plugin.
type FakeCSIPlugin struct {
	csi.UnimplementedControllerServer
	csi.UnimplementedIdentityServer // Add this to implement identity methods

	mu          sync.Mutex // Protects retryCount
	retryCount  int        // Tracks the number of retries
	maxRetries  int        // Max retries before succeeding
}

// NewFakeCSIPlugin creates a new instance of the FakeCSIPlugin with a retry limit.
func NewFakeCSIPlugin(maxRetries int) *FakeCSIPlugin {
	return &FakeCSIPlugin{
		retryCount:  0,
		maxRetries:  maxRetries,
	}
}

func main() {

	log.Default().Printf("Starting csi driver server for yunikorn e2e testing")
	// Set up gRPC server
	listener, err := net.Listen("tcp", ":9890")
	if err != nil {
		log.Fatal("Failed to listen on port 9890: %v", err)
	}

	server := grpc.NewServer()
	fakeCSIPlugin := NewFakeCSIPlugin(3) // Fail 3 times, then succeed
	csi.RegisterControllerServer(server, fakeCSIPlugin)
	csi.RegisterIdentityServer(server, fakeCSIPlugin) // Register the Identity Server

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

// ControllerPublishVolume simulates binding with retries and eventual success
func (p *FakeCSIPlugin) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Simulate retry with failure for the first few attempts
	if p.retryCount < p.maxRetries {
		log.Default().Printf("simulated bind failure on attempt %d: unable to bind volume to node", p.retryCount+1)
		p.retryCount++
		return nil, fmt.Errorf("simulated bind failure: unable to bind volume to node")
	}

	// After maxRetries, return success
	log.Default().Printf("bind successful after %d retries", p.retryCount)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			"bindSuccess": "true",
		},
	}, nil
}

func (p *FakeCSIPlugin) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// Simulate successful volume publish (mounting)
	log.Default().Printf("Simulating NodePublishVolume success for volume %s on target path %s", req.VolumeId, req.TargetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}
