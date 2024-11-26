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

type ProjectServer struct {
	projectpb.UnimplementedProjectServiceServer

	svc service.IProjectService
}

func NewProjectServer(svc service.IProjectService) *ProjectServer {
	return &ProjectServer{svc: svc}
}

func (p *ProjectServer) GetProjectByID(ctx context.Context, req *projectpb.GetProjectByIDReq) (*projectpb.GetProjectByIDRep, error) {
	if req.Id == 0 {
		return nil, errcode.ErrInvalidParams
	}

	resp, err := p.svc.GetProjectByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &projectpb.GetProjectByIDRep{
		Id:          resp.ID,
		Uid:         resp.UID,
		Name:        resp.Name,
		Desc:        resp.Desc,
		CreatedTime: resp.CreatedTime.Unix(),
		UpdatedTime: resp.UpdatedTime.Unix(),
	}, nil
}

func (p *ProjectServer) GetProjectByUID(ctx context.Context, req *projectpb.GetProjectByUIDReq) (*projectpb.GetProjectByUIDRep, error) {
	if req.Uid == 0 {
		return nil, errcode.ErrInvalidParams
	}

	projects, total, err := p.svc.GetProjectsByUID(ctx, req.Uid, req.PageNo, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*projectpb.ProjectInfo, 0, len(projects))
	for _, v := range projects {
		list = append(list, &projectpb.ProjectInfo{
			Id:          v.ID,
			Uid:         v.UID,
			Name:        v.Name,
			Desc:        v.Desc,
			CreatedTime: v.CreatedTime.Unix(),
			UpdatedTime: v.UpdatedTime.Unix(),
		})
	}
	return &projectpb.GetProjectByUIDRep{List: list, Total: total}, nil
}

func (p *ProjectServer) CreateProject(ctx context.Context, req *projectpb.CreateProjectReq) (*commonpb.Empty, error) {
	if len(req.Name) == 0 {
		return nil, errcode.ErrInvalidParams
	}

	if err := p.svc.CreateProject(ctx, domain.Project{
		UID:         req.Uid,
		Name:        req.Name,
		Desc:        req.Desc,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (p *ProjectServer) UpdateProjectByID(ctx context.Context, req *projectpb.UpdateProjectByIDReq) (*commonpb.Empty, error) {
	if req.Id == 0 {
		return nil, errcode.ErrInvalidParams
	}

	if err := p.svc.UpdateProjectByID(ctx, req.Id, req.Name, req.Desc); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (p *ProjectServer) DeleteProjectByID(ctx context.Context, req *projectpb.DeleteProjectByIDReq) (*commonpb.Empty, error) {
	if req.Id == 0 {
		return nil, errcode.ErrInvalidParams
	}

	if err := p.svc.DeleteProjectByID(ctx, req.Id); err != nil {
		return nil, err
	}

	return &commonpb.Empty{}, nil
}

func (p *ProjectServer) RegisterServer(server *grpc.Server) {
	projectpb.RegisterProjectServiceServer(server, p)
}
