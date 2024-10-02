package main
/*
	Data fetching and preprocessing done here
*/
import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "../protobuf/analysis.proto"
)

var (
	addr = flag.String("addr", "localhost:42069", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewAnalysisServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	
	// Prepare a sample MarketDataRequest with dummy data for testing.
    request := &pb.MarketDataRequest{
        Data: []*pb.MarketData{
            {Time: "2024-08-31T12:00:00Z", Symbol: "AAPL", Price: 150.25, Volume: 12000},
            {Time: "2024-08-31T12:01:00Z", Symbol: "AAPL", Price: 150.50, Volume: 13000},
            {Time: "2024-08-31T12:02:00Z", Symbol: "AAPL", Price: 150.75, Volume: 14000},
        },
    }
	r, err := c.AnalyzeData(ctx, request)
	
	if err != nil {
		log.Fatalf("could not analyze: %v", err)
	}


	log.Printf("Greeting: %s", r.GetPredictions())
}