// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
// 	"sync"
// 	"time"

// 	orderpb "grpc-demo/proto/orderpb"
// 	userpb "grpc-demo/proto/userpb"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/status"
// )

// // ═══════════════════════════════════════════════════════════════
// // This service CALLS the User Service via gRPC (service-to-service)
// // This is the key pattern in microservices!
// // ═══════════════════════════════════════════════════════════════

// type orderServer struct {
// 	orderpb.UnimplementedOrderServiceServer

// 	mu         sync.Mutex
// 	orders     map[int32]*orderpb.Order
// 	nextID     int32
// 	userClient userpb.UserServiceClient // <-- gRPC client to call User Service
// }

// // --- Unary RPC: CreateOrder ---
// func (s *orderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.Order, error) {
// 	log.Printf("📨 CreateOrder called for user_id: %d, item: %s", req.UserId, req.Item)

// 	// ⭐ SERVICE-TO-SERVICE CALL: Verify user exists by calling User Service
// 	_, err := s.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
// 	if err != nil {
// 		return nil, status.Errorf(codes.NotFound, "cannot create order: user %d not found", req.UserId)
// 	}

// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	order := &orderpb.Order{
// 		Id:       s.nextID,
// 		UserId:   req.UserId,
// 		Item:     req.Item,
// 		Quantity: req.Quantity,
// 		Price:    req.Price,
// 		Status:   "CREATED",
// 	}
// 	s.orders[s.nextID] = order
// 	s.nextID++

// 	log.Printf("✅ Order #%d created successfully", order.Id)
// 	return order, nil
// }

// // --- Unary RPC: GetOrder ---
// func (s *orderServer) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.Order, error) {
// 	log.Printf("📨 GetOrder called with ID: %d", req.Id)

// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	order, exists := s.orders[req.Id]
// 	if !exists {
// 		return nil, status.Errorf(codes.NotFound, "order %d not found", req.Id)
// 	}
// 	return order, nil
// }

// // --- Server Streaming RPC: GetOrdersByUser ---
// func (s *orderServer) GetOrdersByUser(req *orderpb.OrdersByUserRequest, stream orderpb.OrderService_GetOrdersByUserServer) error {
// 	log.Printf("📨 GetOrdersByUser called for user_id: %d", req.UserId)

// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	for _, order := range s.orders {
// 		if order.UserId == req.UserId {
// 			if err := stream.Send(order); err != nil {
// 				return err
// 			}
// 			time.Sleep(300 * time.Millisecond)
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	// ═══════════════════════════════════════════════════════════
// 	// Connect to User Service as a gRPC CLIENT
// 	// ═══════════════════════════════════════════════════════════
// 	userConn, err := grpc.NewClient(
// 		"localhost:50051", // User Service address
// 		grpc.WithTransportCredentials(insecure.NewCredentials()), // No TLS for demo
// 	)
// 	if err != nil {
// 		log.Fatalf("❌ Failed to connect to User Service: %v", err)
// 	}
// 	defer userConn.Close()

// 	userClient := userpb.NewUserServiceClient(userConn)
// 	fmt.Println("🔗 Connected to User Service on :50051")

// 	// ═══════════════════════════════════════════════════════════
// 	// Start Order Service as a gRPC SERVER
// 	// ═══════════════════════════════════════════════════════════
// 	lis, err := net.Listen("tcp", ":50052")
// 	if err != nil {
// 		log.Fatalf("❌ Failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	orderpb.RegisterOrderServiceServer(grpcServer, &orderServer{
// 		orders:     make(map[int32]*orderpb.Order),
// 		nextID:     1,
// 		userClient: userClient,
// 	})

// 	fmt.Println("🚀 Order Service running on :50052")
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("❌ Failed to serve: %v", err)
// 	}
// }
package main

func main() {
	
}



// two servies order and user get the orders of the users i want that data tommarw crack the logic flow of rpc and build form sractg