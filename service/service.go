package service

import (
	"context"
	pb "github.com/fullstackatbrown/hours/pb/out"
	"github.com/nthnluu/aether/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	repository Repository
	*pb.UnimplementedHoursServiceServer
}

// CreateCourse creates a new course with the given options.
func (s *service) CreateCourse(ctx context.Context, createCourseRequest *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	return nil, errors.NotYetImplementedError("Implement me!")
}

// UpdateCourse updates an existing course's options.
func (s *service) UpdateCourse(ctx context.Context, updateCourseRequest *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	return nil, errors.NotYetImplementedError("Implement me!")
}

// DeleteCourse deletes a given course and all data associated with it.
func (s *service) DeleteCourse(ctx context.Context, deleteCourseRequest *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	return nil, errors.NotYetImplementedError("Implement me!")
}

// ArchiveCourse marks a given course as archived, which hides it but doesn't delete its data.
func (s *service) ArchiveCourse(ctx context.Context, archiveCourseRequest *pb.ArchiveCourseRequest) (*pb.ArchiveCourseResponse, error) {
	return nil, errors.NotYetImplementedError("Implement me!")
}

// ArchiveCoursesByTerm archives all courses with the given term. Useful for archiving all courses
// at the end of the semester.
func (s *service) ArchiveCoursesByTerm(ctx context.Context, archiveCoursesByTermRequest *pb.ArchiveCoursesByTermRequest) (*pb.ArchiveCoursesByTermResponse, error) {
	return nil, errors.NotYetImplementedError("Implement me!")
}

// GetQueue look up a queue by id.
func (s *service) GetQueue(ctx context.Context, getQueueRequest *pb.GetQueueRequest) (*pb.GetQueueResponse, error) {
	queue, err := s.repository.GetQueue(getQueueRequest.GetQueueId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.GetQueueResponse{Queue: queue}, nil
}

// CreateQueue creates a queue.
func (s *service) CreateQueue(ctx context.Context, createQueueRequest *pb.CreateQueueRequest) (*pb.CreateQueueResponse, error) {
	queue, err := s.repository.CreateQueue(createQueueRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.CreateQueueResponse{Queue: queue}, nil
}

// UpdateQueue updates an existing queue.
func (s *service) UpdateQueue(ctx context.Context, updateQueueRequest *pb.UpdateQueueRequest) (*pb.UpdateQueueResponse, error) {
	err := s.repository.UpdateQueue(updateQueueRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.UpdateQueueResponse{
		OldQueueOptions: updateQueueRequest.QueueOptions,
	}, nil
}

// DeleteQueue deletes a queue permanently.
func (s *service) DeleteQueue(ctx context.Context, deleteQueueRequest *pb.DeleteQueueRequest) (*pb.DeleteQueueResponse, error) {
	deletedQueue, err := s.repository.DeleteQueue(deleteQueueRequest.QueueId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.DeleteQueueResponse{
		QueueId: deletedQueue.QueueId,
	}, nil
}

func CreateService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}
