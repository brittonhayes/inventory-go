// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	empty "github.com/golang/protobuf/ptypes/empty"
	ent "github.com/open-farms/inventory/ent"
	category "github.com/open-farms/inventory/ent/category"
	location "github.com/open-farms/inventory/ent/location"
	tool "github.com/open-farms/inventory/ent/tool"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ToolService implements ToolServiceServer
type ToolService struct {
	client *ent.Client
	UnimplementedToolServiceServer
}

// NewToolService returns a new ToolService
func NewToolService(client *ent.Client) *ToolService {
	return &ToolService{
		client: client,
	}
}

// toProtoTool transforms the ent type to the pb type
func toProtoTool(e *ent.Tool) (*Tool, error) {
	v := &Tool{}
	createtime := timestamppb.New(e.CreateTime)
	v.CreateTime = createtime
	id := int32(e.ID)
	v.Id = id
	name := e.Name
	v.Name = name
	powered := e.Powered
	v.Powered = powered
	updatetime := timestamppb.New(e.UpdateTime)
	v.UpdateTime = updatetime
	if edg := e.Edges.Category; edg != nil {
		id := int32(edg.ID)
		v.Category = &Category{
			Id: id,
		}
	}
	if edg := e.Edges.Location; edg != nil {
		id := int32(edg.ID)
		v.Location = &Location{
			Id: id,
		}
	}
	return v, nil
}

// Create implements ToolServiceServer.Create
func (svc *ToolService) Create(ctx context.Context, req *CreateToolRequest) (*Tool, error) {
	tool := req.GetTool()
	m := svc.client.Tool.Create()
	toolCreateTime := runtime.ExtractTime(tool.GetCreateTime())
	m.SetCreateTime(toolCreateTime)
	toolName := tool.GetName()
	m.SetName(toolName)
	toolPowered := tool.GetPowered()
	m.SetPowered(toolPowered)
	toolUpdateTime := runtime.ExtractTime(tool.GetUpdateTime())
	m.SetUpdateTime(toolUpdateTime)
	toolCategory := int(tool.GetCategory().GetId())
	m.SetCategoryID(toolCategory)
	toolLocation := int(tool.GetLocation().GetId())
	m.SetLocationID(toolLocation)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTool(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements ToolServiceServer.Get
func (svc *ToolService) Get(ctx context.Context, req *GetToolRequest) (*Tool, error) {
	var (
		err error
		get *ent.Tool
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetToolRequest_VIEW_UNSPECIFIED, GetToolRequest_BASIC:
		get, err = svc.client.Tool.Get(ctx, id)
	case GetToolRequest_WITH_EDGE_IDS:
		get, err = svc.client.Tool.Query().
			Where(tool.ID(id)).
			WithCategory(func(query *ent.CategoryQuery) {
				query.Select(category.FieldID)
			}).
			WithLocation(func(query *ent.LocationQuery) {
				query.Select(location.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoTool(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements ToolServiceServer.Update
func (svc *ToolService) Update(ctx context.Context, req *UpdateToolRequest) (*Tool, error) {
	tool := req.GetTool()
	toolID := int(tool.GetId())
	m := svc.client.Tool.UpdateOneID(toolID)
	toolCreateTime := runtime.ExtractTime(tool.GetCreateTime())
	m.SetCreateTime(toolCreateTime)
	toolName := tool.GetName()
	m.SetName(toolName)
	toolPowered := tool.GetPowered()
	m.SetPowered(toolPowered)
	toolUpdateTime := runtime.ExtractTime(tool.GetUpdateTime())
	m.SetUpdateTime(toolUpdateTime)
	toolCategory := int(tool.GetCategory().GetId())
	m.SetCategoryID(toolCategory)
	toolLocation := int(tool.GetLocation().GetId())
	m.SetLocationID(toolLocation)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTool(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements ToolServiceServer.Delete
func (svc *ToolService) Delete(ctx context.Context, req *DeleteToolRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Tool.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}
