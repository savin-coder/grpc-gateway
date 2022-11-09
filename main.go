package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
	//gw "github.com/gtm/v1" // Update
	gw "ecards-grpc-gateway/generated/ecard" // Update
)

var (
	//grpcServerEndpoint = flag.String("grpc-server-endpoint", os.Getenv("HOSTPORT"), "gRPC server endpoint")
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	js, err := json.Marshal(grpcServerEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible

	mux := runtime.NewServeMux()

	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1000 * 1024 * 1024)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(20 * 1024 * 1024)),
	}

	//opts := []grpc.MaxRecvMsgSizeCallOption(20*1024*1024)

	err := gw.RegisterTemplateServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts) //Register<YOURSERVICENAME>HandlerFromEndpoint - doing this your work will be okay
	if err != nil {
		return err
	}
	err3 := gw.RegisterCategoryServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err3 != nil {
		return err3
	}
	err4 := gw.RegisterStickerServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err4 != nil {
		return err4
	}
	err5 := gw.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err5 != nil {
		return err5
	}
	err6 := gw.RegisterCheckServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err6 != nil {
		return err6
	}

	return http.ListenAndServe(":8090", allowCORS(mux))
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Grpc-Metadata-app-id", "Grpc-Metadata-api-key"}
	w.Header().Set("Access-Control-Expose-Headers", strings.Join(headers, ","))
	w.Header().Set("Content-Type", "application/json") //extra add
	w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			//w.Header().Set("Access-Control-Allow-Origin", origin)
			(w).Header().Set("Access-Control-Allow-Origin", "*")
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
