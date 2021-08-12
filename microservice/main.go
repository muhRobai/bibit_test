package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcServer "github.com/bibit/microservice/delivery/grpc"
	httpServer "github.com/bibit/microservice/delivery/http"
	"github.com/bibit/microservice/domain"
	"github.com/bibit/microservice/infrastructure/omdb"
	"github.com/bibit/microservice/service"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := os.Setenv("GRPC_GO_RETRY", "on"); err != nil {
		log.Printf("error while setting grpc retry envar: %v", err)
	}
}

func startHTTP(addr string, api domain.ServiceAPI) {
	router := mux.NewRouter()
	server := httpServer.NewHandler(router, api)
	// server := httpServer.NewHandler(router, api)

	log.Println("server running on localhost", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Println("[failed start server]", err)
		return
	}
}

// Serve listen for client request
func Serve(address string, server *grpc.Server) {

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Print(err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		<-c
		log.Println("Shutting down server gracefully...")
		server.GracefulStop()
	}()

	err = server.Serve(lis)
	if err != nil {
		log.Print(err)
		return
	}
}

func main() {
	movieAPI := omdb.NewOmdbService(
		"http://www.omdbapi.com",
		"faf7e5bb",
	)

	API := service.NewServiceAPI(movieAPI)
	server := grpc.NewServer()
	grpcServer.NewGrpcHandler(server, API)

	// start grpc server
	go func() {
		Serve(":8080", server)
	}()

	go startHTTP(":8081", API)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Println("All server stopped!")

}
