package api

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Server) GetClusters(ctx context.Context, _ *pb.GetClustersRequest) (*pb.GetClustersResponse, error) {
	clusters, err := s.store.GetAllClusters(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the clusters to the protobuf format
	var pbClusters []*pb.Cluster
	for _, cluster := range clusters {
		pbClusters = append(pbClusters, convertClusterToPb(cluster))
	}

	return &pb.GetClustersResponse{
		Clusters: pbClusters,
	}, nil
}
