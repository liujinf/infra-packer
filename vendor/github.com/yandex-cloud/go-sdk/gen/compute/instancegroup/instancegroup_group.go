// Code generated by sdkgen. DO NOT EDIT.

package instancegroup

import (
	"context"

	"google.golang.org/grpc"
)

// InstanceGroup provides access to "instancegroup" component of Yandex.Cloud
type InstanceGroup struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewInstanceGroup creates instance of InstanceGroup
func NewInstanceGroup(g func(ctx context.Context) (*grpc.ClientConn, error)) *InstanceGroup {
	return &InstanceGroup{g}
}

// InstanceGroup gets InstanceGroupService client
func (i *InstanceGroup) InstanceGroup() *InstanceGroupServiceClient {
	return &InstanceGroupServiceClient{getConn: i.getConn}
}