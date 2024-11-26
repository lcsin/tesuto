package rpc

import (
	"context"
	"time"

	"github.com/lcsin/tesuto/pkg/errcode"
	"github.com/lcsin/tesuto/proto/v1/rpc/commonpb"
	"github.com/lcsin/tesuto/proto/v1/rpc/projectpb"
	"github.com/lcsin/tesuto/tesuto-project/internel/domain"
	"github.com/lcsin/tesuto/tesuto-project/internel/service"
	"google.golang.org/grpc"
)

type ModuleServer struct {
	projectpb.UnimplementedModuleServiceServer

	svc service.IModuleService
}

func NewModuleServer(svc service.IModuleService) *ModuleServer {
	return &ModuleServer{svc: svc}
}

func (m *ModuleServer) GetModuleByID(ctx context.Context, req *projectpb.GetModuleByIDReq) (*projectpb.GetModuleByIDRep, error) {
	if req.Id == 0 || req.Uid == 0 {
		return nil, errcode.ErrInvalidParams
	}

	resp, err := m.svc.GetModuleByID(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	return &projectpb.GetModuleByIDRep{
		Id:          resp.ID,
		Uid:         resp.UID,
		ProjectId:   resp.ProjectID,
		Name:        resp.Name,
		Desc:        resp.Desc,
		CreatedTime: resp.CreatedTime.Unix(),
		UpdatedTime: resp.UpdatedTime.Unix(),
	}, nil
}

func (m *ModuleServer) GetModulesByProjectID(ctx context.Context, req *projectpb.GetModulesByProjectIDReq) (*projectpb.GetModulesByProjectIDRep, error) {
	if req.ProjectId == 0 || req.Uid == 0 {
		return nil, errcode.ErrInvalidParams
	}

	modules, err := m.svc.GetModuleByProjectID(ctx, req.ProjectId, req.Uid)
	if err != nil {
		return nil, err
	}
	list := make([]*projectpb.ModuleInfo, 0, len(modules))
	for _, v := range modules {
		list = append(list, &projectpb.ModuleInfo{
			Id:          v.ID,
			Uid:         v.UID,
			ProjectId:   v.ProjectID,
			Name:        v.Name,
			Desc:        v.Desc,
			CreatedTime: v.CreatedTime.Unix(),
			UpdatedTime: v.UpdatedTime.Unix(),
		})
	}
	return &projectpb.GetModulesByProjectIDRep{List: list}, nil
}

func (m *ModuleServer) CreateModule(ctx context.Context, req *projectpb.CreateModuleReq) (*commonpb.Empty, error) {
	if req.Uid == 0 || req.ProjectId == 0 || len(req.Name) == 0 {
		return nil, errcode.ErrInvalidParams
	}

	if err := m.svc.CreateModule(ctx, domain.Module{
		UID:         req.Uid,
		ProjectID:   req.ProjectId,
		Name:        req.Name,
		Desc:        req.Desc,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (m *ModuleServer) UpdateModuleByID(ctx context.Context, req *projectpb.UpdateModuleByIDReq) (*commonpb.Empty, error) {
	if req.Uid == 0 || req.Id == 0 || len(req.Name) == 0 {
		return nil, errcode.ErrInvalidParams
	}

	if err := m.svc.UpdateModuleByID(ctx, req.Id, req.Uid, req.Name, req.Desc); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (m *ModuleServer) DeleteModuleByID(ctx context.Context, req *projectpb.DeleteModuleByIDReq) (*commonpb.Empty, error) {
	if req.Id == 0 || req.Uid == 0 {
		return nil, errcode.ErrInvalidParams
	}

	if err := m.svc.DeleteModuleByID(ctx, req.Id, req.Uid); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (m *ModuleServer) RegisterServer(server *grpc.Server) {
	projectpb.RegisterModuleServiceServer(server, m)
}
