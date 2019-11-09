package service

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
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

// CreateConfiguration will create the current configuration state.
func (s *Service) CreateConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create configuration ->"+err.Error())
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
		return nil, status.Error(codes.Unknown, "persistent store failed to update configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateConfiguration"}, nil
}

// DestroyConfiguration will check the current configuration state.
func (s *Service) DestroyConfiguration(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to destroy configuration ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Destroyed configuration."}, nil
}

// CreateUser will create the current User state.
func (s *Service) CreateUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateUser"}, nil
}

// ReadUser will read the current User state.
func (s *Service) ReadUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to read User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadUser"}, nil
}

// ListUsers will check the current User state.
func (s *Service) ListUsers(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to list User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadUser"}, nil
}

// UpdateUser will update the current User state.
func (s *Service) UpdateUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to update User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateUser"}, nil
}

// ArchiveUser will archive the current User state.
func (s *Service) ArchiveUser(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to archive User ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ArchiveUser"}, nil
}

func convertPGtoProto(pg store.Movement) *pb.Movement {
	return &pb.Movement{
		ID:            pg.ID.String(),
		Title:         pg.Title,
		Description:   pg.Description,
		URI:           pg.URI,
		FeaturedImage: pg.FeaturedImage,
	}
}

// CreateMovement will create the current Movement state.
func (s *Service) CreateMovement(ctx context.Context, in *pb.CreateMovementRequest) (*pb.CreateMovementResponse, error) {
	newMovement, err := store.CreateMovement(in.Title, in.User, in.Description, in.URI, in.FeaturedImage)
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create Movement ->"+err.Error())
	}

	return &pb.CreateMovementResponse{Item: convertPGtoProto(*newMovement)}, nil
}

// ReadMovement will read the current Movement state.
func (s *Service) ReadMovement(ctx context.Context, in *pb.ReadMovementRequest) (*pb.ReadMovementResponse, error) {
	movement, err := store.GetMovement(in.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to read Movement ->"+err.Error())
	}

	return &pb.ReadMovementResponse{Item: &pb.Movement{ID: movement.ID.String(), Title: movement.Title, Description: movement.Description, URI: movement.URI, FeaturedImage: movement.FeaturedImage}}, nil
}

// ListMovements will check the current Movement state.
func (s *Service) ListMovements(ctx context.Context, in *pb.ListMovementsRequest) (*pb.ListMovementsResponse, error) {
	movements, err := store.ListMovements()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to list Movements ->"+err.Error())
	}
	list := []*pb.Movement{}
	for _, movement := range *movements {
		list = append(list, &pb.Movement{ID: movement.ID.String(), Title: movement.Title, Description: movement.Description, URI: movement.URI, FeaturedImage: movement.FeaturedImage})
	}
	return &pb.ListMovementsResponse{Data: list}, nil
}

// UpdateMovement will check the current Movement state.
func (s *Service) UpdateMovement(ctx context.Context, in *pb.UpdateMovementRequest) (*pb.UpdateMovementResponse, error) {
	idFromString, err := uuid.FromString(in.Item.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "unable to create valid UUID from input"+err.Error())
	}

	creatorID, err := uuid.FromString("bd84693a-7191-40e3-90a7-eb77eda808c1")
	if err != nil {
		return nil, status.Error(codes.Unknown, "unable to convert string to UUID of creator"+err.Error())
	}

	updatedMovement := &store.Movement{
		ID:            idFromString,
		Title:         in.Item.Title,
		Description:   in.Item.Description,
		URI:           in.Item.URI,
		FeaturedImage: in.Item.FeaturedImage,
		CreatorID:     creatorID, //This can come from a value injected by interceptor
	}

	movement, err := store.UpdateMovement(updatedMovement)
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to update Movement ->"+err.Error())
	}

	return &pb.UpdateMovementResponse{Item: convertPGtoProto(*movement)}, nil
}

// DeleteMovement will check the current Movement state.
func (s *Service) DeleteMovement(ctx context.Context, in *pb.DeleteMovementRequest) (*pb.DeleteMovementResponse, error) {
	movement, err := store.DeleteMovement(in.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to Delete Movement ->"+err.Error())
	}

	return &pb.DeleteMovementResponse{Item: convertPGtoProto(*movement)}, nil
}

// CreateLandingPage will check the current LandingPage state.
func (s *Service) CreateLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to create LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for CreateLandingPage"}, nil
}

// ReadLandingPage will check the current LandingPage state.
func (s *Service) ReadLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to read LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadLandingPage"}, nil
}

// ListLandingPages will check the current LandingPage state.
func (s *Service) ListLandingPages(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to list LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for ReadLandingPage"}, nil
}

// UpdateLandingPage will check the current LandingPage state.
func (s *Service) UpdateLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.InitializeDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to update LandingPage ->"+err.Error())
	}

	return &pb.StubResponse{Message: "Response for UpdateLandingPage"}, nil
}

// ArchiveLandingPage will check the current LandingPage state.
func (s *Service) ArchiveLandingPage(ctx context.Context, in *pb.StubRequest) (*pb.StubResponse, error) {
	err := store.DestroyDB()
	if err != nil {
		return nil, status.Error(codes.Unknown, "persistent store failed to archive LandingPage ->"+err.Error())
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
