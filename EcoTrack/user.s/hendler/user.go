package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"dodi/db"
	pb "dodi/genproto/pro1tos" // Protokol paketini Go tiliga ko'chirish

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type server struct {
	client pb.UserServiceClient
}

func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	fmt.Println(userID)

	// Check if s.client is nil
	if s.client == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("s.client is nil")
		return
	}

	resp, err := s.client.GetUser(context.Background(), &pb.GetUserRequest{UserId: int32(userID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

func (s *server) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req pb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	req.UserId = int32(userID)

	resp, err := s.client.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *server) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resp, err := s.client.DeleteUser(context.Background(), &pb.DeleteUserRequest{UserId: int32(userID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *server) GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resp, err := s.client.GetUserProfile(context.Background(), &pb.GetUserProfileRequest{UserId: int32(userID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *server) UpdateUserProfile(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req pb.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	req.UserId = int32(userID)

	resp, err := s.client.UpdateUserProfile(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func main() {
	db1, err := db.ConnectDB()
	db.NewCourseRepo(db1)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	s := &server{client: client}

	r := mux.NewRouter()

	r.HandleFunc("/api/users/:id", s.GetUser).Methods("GET")

	http.ListenAndServe(":8080", r)

}
