/**
 * @Description: 用户相关 service
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/16 14:09
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/16 14:09
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/system/v1"
	"vine-template-rpc/internal/page"
)

func (s *SystemService) AddUser(ctx context.Context, request *v1.AddUserRequest) (*v1.AddUserReply, error) {
	//conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("couldn't connect: %v", err)
	//}
	//defer conn.Close()
	//
	//c := pb.NewCmsClient(conn)
	//r, err := c.CreateCms(context.Background(), &pb.CreateCmsRequest{})
	//if err != nil {
	//	//e := errors.FromError(err)
	//	//fmt.Printf("%s", e.)
	//	fmt.Printf("err: %v", err.Error())
	//	return nil, err
	//}
	//
	//fmt.Printf("Content: %s", r.Name)
	return s.ub.AddUser(ctx, request)
}

func (s *SystemService) UpdateUser(ctx context.Context, request *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}

func (s *SystemService) DeleteUser(ctx context.Context, request *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}

func (s *SystemService) GetUser(ctx context.Context, request *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return s.ub.GetUser(ctx, request)
}

func (s *SystemService) ListUser(ctx context.Context, request *v1.ListUserRequest) (*page.Page, error) {
	return s.ub.ListUser(ctx, request)
}
