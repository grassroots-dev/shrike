package service

import (
	"context"
	"time"

	pb "github.com/grassroots-dev/shrike/api"
	"github.com/grassroots-dev/shrike/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service is a type represneting a configured server.
type Service struct {
	DB      string
	Cache   string
	Storage string
}

// NewService returns a new service configured with all the resources.
func NewService(db string, cache string, storage string) *Service {
	return &Service{DB: "hello db", Cache: "hello cache", Storage: "hello storage"}
}

// CheckConfiguration will check the current configuration state.
func (s *Service) CheckConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CheckConfiguration"}, nil
}

// Configure will configure a new running instance of the shrike service.
func (s *Service) Configure(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to set configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Configuration initialized."}, nil
}

// ResetConfiguration will reset the service to factory settings.
func (s *Service) ResetConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.ResetDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to reset configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Configuration destroyed."}, nil
}

// CheckUser will return the currently logged in user..
func (s *Service) CheckUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Account -> "+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CheckUser"}, nil
}

// CreateUser will create a new user and associated resources.
func (s *Service) CreateUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create user -> "+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateUser"}, nil
}

// CreateMovement will create a new movement and all dependencies.
func (s *Service) CreateMovement(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create movement ->  "+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateMovement"}, nil
}

// ArchiveMovement will archive an active movement.
func (s *Service) ArchiveMovement(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to archive movement ->  "+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ArchiveMovement"}, nil
}

// CreateLandingPage will create a new landing page belonging to a movement and all dependencies.
func (s *Service) CreateLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create landing page ->  "+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateLandingPage"}, nil
}

// ArchiveLandingPage will archive an active landing page.
func (s *Service) ArchiveLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to archive landing page -> "+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ArchiveLandingPage"}, nil
}

// GetVolunteerLocationStream will return a stream of Volunteer GPS coordinates.
func (s *Service) GetVolunteerLocationStream(in *pb.StubRequest, stream pb.Shrike_GetVolunteerLocationStreamServer) error {
	for {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&pb.StubResponse{Message: "hello!"}); err != nil {
			return err
		}

	}
}
