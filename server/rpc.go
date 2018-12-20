package server

import (
	"context"
	"net"

	phdstore "github.com/pizzahutdigital/datastore"
	"github.com/pkg/errors"
	"github.com/scottshotgg/storage-server/handlers"
	"github.com/scottshotgg/storage/impl/datastore"
	pb "github.com/scottshotgg/storage/protobufs"
	"github.com/spf13/viper"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

// RunRPC starts rpc server for geosearch
func RunRPC() error {
	var servicename = viper.GetString("servicename")
	if servicename == "" {
		return errors.New("You must supply a valid servicename for logging using the `servicename` flag")
	}

	var rpcPort = viper.GetString("rpc-port")
	if rpcPort == "" {
		return errors.New("You must supply a valid port using the 'rpc-port' argument")
	}

	var lis, err = net.Listen("tcp", ":"+rpcPort)
	if err != nil {
		return errors.Wrap(err, "failed to initialize TCP listen")
	}

	defer func() {
		var err = lis.Close()
		if err != nil {
			// log
		}
	}()

	// Switch on a config file
	// switch {}

	ds, err := datastore.New(phdstore.DSConfig{
		Context:            context.Background(),
		ServiceAccountFile: "/Users/sgg7269/Documents/serviceAccountFiles/ds-serviceaccount.json",
		ProjectID:          "phdigidev",
		Namespace:          "storage_test",
	})
	if err != nil {
		return errors.Wrap(err, "handlers.NewGeosearch")
	}

	// Try to make a new Geosearch before even starting the server
	s, err := handlers.New(ds)
	if err != nil {
		return errors.Wrap(err, "handlers.NewGeosearch")
	}

	var rpcServer = grpc.NewServer(
		grpc.StatsHandler(&ocgrpc.ServerHandler{
			StartOptions: trace.StartOptions{
				Sampler: trace.AlwaysSample(),
			},
		}))

	pb.RegisterStorageServer(rpcServer, s)

	// log
	return rpcServer.Serve(lis)
}
