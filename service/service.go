package service

import (
	"context"
	"fmt"
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

// CreateConfiguration will check the current configuration state.
func (s *Service) CreateConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateConfiguration"}, nil
}

// ReadConfiguration will check the current configuration state.
func (s *Service) ReadConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadConfiguration"}, nil
}

// UpdateConfiguration will check the current configuration state.
func (s *Service) UpdateConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateConfiguration"}, nil
}

// DestroyConfiguration will check the current configuration state.
func (s *Service) DestroyConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Destroyed configuration."}, nil
}

// CreateUser will check the current User state.
func (s *Service) CreateUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateUser"}, nil
}

// ReadUser will check the current User state.
func (s *Service) ReadUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadUser"}, nil
}

// ListUsers will check the current User state.
func (s *Service) ListUsers(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadUser"}, nil
}

// UpdateUser will check the current User state.
func (s *Service) UpdateUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateUser"}, nil
}

// ArchiveUser will check the current User state.
func (s *Service) ArchiveUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ArchiveUser"}, nil
}

// CreateMovement will check the current Movement state.
func (s *Service) CreateMovement(ctx context.Context, in *pb.CreateMovementRequest) (*pb.StubResponse, error) {
	newID, err := store.CreateMovement(in.Title)
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check Movement ->"+err.Error())
	}

	return &pb.StubResponse{Message: fmt.Sprintf("Succsefully created: %v", *newID)}, nil
}

// ReadMovement will check the current Movement state.
func (s *Service) ReadMovement(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check Movement ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadMovement"}, nil
}

// ListMovements will check the current Movement state.
func (s *Service) ListMovements(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check Movement ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadMovement"}, nil
}

// UpdateMovement will check the current Movement state.
func (s *Service) UpdateMovement(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check Movement ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateMovement"}, nil
}

// ArchiveMovement will check the current Movement state.
func (s *Service) ArchiveMovement(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check Movement ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ArchiveMovement"}, nil
}

// CreateLandingPage will check the current LandingPage state.
func (s *Service) CreateLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateLandingPage"}, nil
}

// ReadLandingPage will check the current LandingPage state.
func (s *Service) ReadLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadLandingPage"}, nil
}

// ListLandingPages will check the current LandingPage state.
func (s *Service) ListLandingPages(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadLandingPage"}, nil
}

// UpdateLandingPage will check the current LandingPage state.
func (s *Service) UpdateLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateLandingPage"}, nil
}

// ArchiveLandingPage will check the current LandingPage state.
func (s *Service) ArchiveLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to check LandingPage ->"+err.Error())
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
