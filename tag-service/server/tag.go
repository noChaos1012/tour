package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/noChaos1012/tour/tag-service/pkg/bapi"
	"github.com/noChaos1012/tour/tag-service/pkg/errcode"
	pb "github.com/noChaos1012/tour/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	fmt.Println("start")
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ERROR_GET_TAG_LIST_FAIL)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	fmt.Printf("list is %v", tagList)
	return &tagList, nil
}
