// package main

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"log"
// 	"time"

// 	orderpb "grpc-demo/proto/orderpb"
// 	userpb "grpc-demo/proto/userpb"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func main() {
// 	// ═══════════════════════════════════════════════════════════
// 	// Create gRPC client connections to BOTH services
// 	// ═══════════════════════════════════════════════════════════

// 	// Connect to User Service
// 	userConn, err := grpc.NewClient("localhost:50051",
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	if err != nil {
// 		log.Fatalf("❌ Cannot connect to User Service: %v", err)
// 	}
// 	defer userConn.Close()
// 	userClient := userpb.NewUserServiceClient(userConn)

// 	// Connect to Order Service
// 	orderConn, err := grpc.NewClient("localhost:50052",
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	if err != nil {
// 		log.Fatalf("❌ Cannot connect to Order Service: %v", err)
// 	}
// 	defer orderConn.Close()
// 	orderClient := orderpb.NewOrderServiceClient(orderConn)

// 	// Timeout context — all RPCs should complete within 10 seconds
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// ═══════════════════════════════════════════════════════════
// 	// TEST 1: Unary RPC — Get a single user
// 	// ═══════════════════════════════════════════════════════════
// 	fmt.Println("\n━━━ TEST 1: GetUser (Unary RPC) ━━━")
// 	user, err := userClient.GetUser(ctx, &userpb.GetUserRequest{Id: 1})
// 	if err != nil {
// 		log.Fatalf("GetUser failed: %v", err)
// 	}
// 	fmt.Printf("Got user: %s (%s)\n", user.Name, user.Email)

// 	// ═══════════════════════════════════════════════════════════
// 	// TEST 2: Server Streaming RPC — List all users
// 	// ═══════════════════════════════════════════════════════════
// 	fmt.Println("\n━━━ TEST 2: ListUsers (Server Streaming RPC) ━━━")
// 	stream, err := userClient.ListUsers(ctx, &userpb.Empty{})
// 	if err != nil {
// 		log.Fatalf("ListUsers failed: %v", err)
// 	}
// 	for {
// 		user, err := stream.Recv() // Receive one user at a time
// 		if err == io.EOF {
// 			break // Stream finished
// 		}
// 		if err != nil {
// 			log.Fatalf("Stream error: %v", err)
// 		}
// 		fmt.Printf("  Streamed user: %s (%s)\n", user.Name, user.Email)
// 	}

// 	// ═══════════════════════════════════════════════════════════
// 	// TEST 3: Create orders (Order Service calls User Service internally!)
// 	// ═══════════════════════════════════════════════════════════
// 	fmt.Println("\n━━━ TEST 3: CreateOrder (Service-to-Service call) ━━━")

// 	order1, err := orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{
// 		UserId: 1, Item: "Mechanical Keyboard", Quantity: 1, Price: 79.99,
// 	})
// 	if err != nil {
// 		log.Fatalf("CreateOrder failed: %v", err)
// 	}
// 	fmt.Printf("Created: Order #%d — %s ($%.2f)\n", order1.Id, order1.Item, order1.Price)

// 	order2, err := orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{
// 		UserId: 1, Item: "USB-C Cable", Quantity: 3, Price: 12.99,
// 	})
// 	if err != nil {
// 		log.Fatalf("CreateOrder failed: %v", err)
// 	}
// 	fmt.Printf("Created: Order #%d — %s ($%.2f)\n", order2.Id, order2.Item, order2.Price)

// 	// ═══════════════════════════════════════════════════════════
// 	// TEST 4: Try creating order for non-existent user (should fail)
// 	// ═══════════════════════════════════════════════════════════
// 	fmt.Println("\n━━━ TEST 4: CreateOrder for invalid user (error handling) ━━━")
// 	_, err = orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{
// 		UserId: 999, Item: "Ghost Item", Quantity: 1, Price: 0,
// 	})
// 	if err != nil {
// 		fmt.Printf("Expected error: %v\n", err)
// 	}

// 	// ═══════════════════════════════════════════════════════════
// 	// TEST 5: Stream orders for user 1
// 	// ═══════════════════════════════════════════════════════════
// 	fmt.Println("\n━━━ TEST 5: GetOrdersByUser (Server Streaming) ━━━")
// 	orderStream, err := orderClient.GetOrdersByUser(ctx, &orderpb.OrdersByUserRequest{UserId: 1})
// 	if err != nil {
// 		log.Fatalf("GetOrdersByUser failed: %v", err)
// 	}
// 	for {
// 		order, err := orderStream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Stream error: %v", err)
// 		}
// 		fmt.Printf("  Streamed order: #%d — %s (qty: %d)\n", order.Id, order.Item, order.Quantity)
// 	}

// 	fmt.Println("\n✅ All tests passed!")
// }
package main

func main() {
	
}