package main

import (
	"net/http"
	"log"
	"time"
	"context"

	pb "github.com/josefmay/grah2.0/go/internal/genproto/analysis"
)

func handleOptimize(w http.ResponseWriter, r *http.Request){
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
		log.Printf("could not analyze: %v", err)
	}

	return
}