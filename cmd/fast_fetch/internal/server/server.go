package server

func FetchUsers() []string {
	return nil
}

type Server struct {
}

type ServerResponse struct {
	Users           []string
	Payments        []string
	DeliveryTracker string
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) fetchUsers() []string {
	return nil
}

func (s *Server) fetchPayments() []string {
	return nil
}

func (s *Server) fetchDeliveryTracker() []string {
	return nil
}

func (s *Server) FetchAllSync() (*ServerResponse, error) {
	return nil, nil
}

func (s *Server) FetchAll() (*ServerResponse, error) {
	return nil, nil
}
