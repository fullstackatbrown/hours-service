package service

import (
	pb "github.com/fullstackatbrown/hours/pb/out"
	"github.com/google/uuid"
	"github.com/nthnluu/aether/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository interface {
	GetQueue(queueId string) (*pb.Queue, error)
	CreateQueue(createQueueRequest *pb.CreateQueueRequest) (*pb.Queue, error)
	UpdateQueue(updateQueueRequest *pb.UpdateQueueRequest) error
	DeleteQueue(queueId string) (*pb.Queue, error)
}

type InMemoryRepository struct {
	queueMapping map[string]*pb.Queue
}

func (r *InMemoryRepository) UpdateQueue(updateQueueRequest *pb.UpdateQueueRequest) error {
	//TODO implement me
	return errors.NotYetImplementedError("Implement me")
}

func (r *InMemoryRepository) DeleteQueue(queueId string) (*pb.Queue, error) {
	//TODO implement me
	return nil, errors.NotYetImplementedError("Implement me")
}

func (r *InMemoryRepository) GetQueue(queueId string) (*pb.Queue, error) {
	queue, ok := r.queueMapping[queueId]
	if !ok {
		return nil, status.Error(codes.NotFound, "queue not found")
	}
	return queue, nil
}

func (r *InMemoryRepository) CreateQueue(createQueueRequest *pb.CreateQueueRequest) (*pb.Queue, error) {
	id := uuid.NewString()
	newQueue := &pb.Queue{
		QueueId: id,
		CreationMetadata: &pb.CreationMetadata{
			CreatedBy: "",
			CreatedAt: nil,
			UpdatedAt: nil,
		},
		Course:       "",
		QueueOptions: createQueueRequest.QueueOptions,
	}
	r.queueMapping[id] = newQueue
	return newQueue, nil
}

func CreateInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		queueMapping: make(map[string]*pb.Queue),
	}
}
