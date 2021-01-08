// Copyright 2021 The gqlp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sacloud/gqlp/graph/generated"
	"github.com/sacloud/gqlp/graph/model"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func (r *mutationResolver) ShutdownServer(ctx context.Context, zone string, id int64, option *model.ShutdownOption) (*model.MutationResult, error) {
	caller := r.APICaller()
	serverOp := sacloud.NewServerOp(caller)
	opt := &sacloud.ShutdownOption{}
	if option != nil && option.Force {
		opt.Force = true
	}
	return r.MutationResult(serverOp.Shutdown(ctx, zone, types.ID(id), opt))
}

func (r *queryResolver) Servers(ctx context.Context, zone string) ([]*model.Server, error) {
	caller := r.APICaller()
	serverOp := sacloud.NewServerOp(caller)
	results, err := serverOp.Find(ctx, zone, &sacloud.FindCondition{})
	if err != nil {
		return nil, err
	}
	var servers []*model.Server
	for _, s := range results.Servers {
		servers = append(servers, &model.Server{
			ID:               s.ID.Int64(),
			Name:             s.Name,
			Tags:             s.Tags,
			Description:      s.Description,
			Availability:     string(s.Availability),
			HostName:         s.HostName,
			InterfaceDriver:  s.InterfaceDriver.String(),
			PlanID:           s.ServerPlanID.Int64(),
			PlanName:         s.ServerPlanName,
			CPU:              s.CPU,
			Memory:           s.GetMemoryGB(),
			Commitment:       s.ServerPlanCommitment.String(),
			PlanGeneration:   int(s.ServerPlanGeneration),
			InstanceHostName: s.InstanceHostName,
			InstanceStatus:   string(s.InstanceStatus),
		})
	}
	return servers, nil
}

func (r *serverResolver) Disks(ctx context.Context, obj *model.Server) ([]*model.Disk, error) {
	rezCtx := graphql.GetFieldContext(ctx)
	zone := rezCtx.Parent.Parent.Args["zone"].(string)

	caller := r.APICaller()
	diskOp := sacloud.NewDiskOp(caller)

	results, err := diskOp.Find(ctx, zone, &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Server.ID"): search.OrEqual(obj.ID),
		},
	})
	if err != nil {
		return nil, err
	}
	var disks []*model.Disk
	for _, d := range results.Disks {
		if d.ServerID.Int64() == obj.ID {
			disks = append(disks, &model.Disk{
				ID:          d.ID.Int64(),
				Name:        d.Name,
				Tags:        d.Tags,
				Description: d.Description,
				Size:        d.GetSizeGB(),
			})
		}
	}
	return disks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Server returns generated.ServerResolver implementation.
func (r *Resolver) Server() generated.ServerResolver { return &serverResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type serverResolver struct{ *Resolver }
