package server

import (
	"fmt"
	"sync"
	"time"
)

func FetchUsers() []string {
	return nil
}

type Server struct {
}

type Entity struct {
	User            string
	Payment         string
	DeliveryTracker string
}
type ServerResponse struct {
	TotalResults int
	Resources    []*Entity
}

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func NewServer() *Server {
	return &Server{}
}

func (s *Server) fetchUsers() []string {
	defer wg.Done()

	result := []string{"Jordan", "John", "Alex"}
	// Simulate short delay
	time.Sleep(3 * time.Second)
	return result
}

func (s *Server) fetchPayments() []string {
	defer wg.Done()

	result := []string{"$10.99", "$19.99", "$25.89"}
	// Simulate short delay
	time.Sleep(3 * time.Second)
	return result
}

func (s *Server) fetchDeliveryTracker() []string {
	defer wg.Done()

	result := []string{"completed", "pending", "pending"}
	// Simulate short delay
	time.Sleep(3 * time.Second)
	return result
}

func toResource(user string, payment string, deliveryTracker string) *Entity {
	return &Entity{
		User:            user,
		Payment:         payment,
		DeliveryTracker: deliveryTracker,
	}
}
func (s *Server) FetchAllSync() (*ServerResponse, error) {
	startTime := time.Now()
	fmt.Printf("Start time: %v \n", startTime.Format("15:04:05"))
	resources := make([]*Entity, 0, 3)

	users := s.fetchUsers()
	payments := s.fetchPayments()
	deliveries := s.fetchDeliveryTracker()

	for i := 0; i < 3; i++ {
		resources = append(resources, toResource(users[i], payments[i], deliveries[i]))
	}

	response := &ServerResponse{
		TotalResults: 3,
		Resources:    resources,
	}
	endTime := time.Now()

	fmt.Printf("End time: %v \n", endTime.Format("15:04:05"))
	fmt.Printf("Time Elapsed: ~%v seconds \n", (endTime.UnixMilli()-startTime.UnixMilli())/1000)

	return response, nil
}

func (s *Server) FetchAll() (*ServerResponse, error) {
	startTime := time.Now()
	fmt.Printf("Start time: %v \n", startTime.Format("15:04:05"))
	resources := make([]*Entity, 0, 3)

	wg.Add(3)

	var users, payments, deliveries []string

	go func() {
		users = s.fetchUsers()
	}()
	go func() {
		payments = s.fetchPayments()
	}()

	go func() {
		deliveries = s.fetchDeliveryTracker()
	}()

	wg.Wait()

	for i := 0; i < 3; i++ {
		resources = append(resources, toResource(users[i], payments[i], deliveries[i]))
	}

	response := &ServerResponse{
		TotalResults: 3,
		Resources:    resources,
	}
	endTime := time.Now()

	fmt.Printf("End time: %v \n", endTime.Format("15:04:05"))
	fmt.Printf("Time Elapsed: ~%v seconds \n", (endTime.UnixMilli()-startTime.UnixMilli())/1000)

	return response, nil
}
