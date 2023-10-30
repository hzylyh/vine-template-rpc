package pagehelper

import (
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
	pb "vine-template-rpc/api/system/v1"
	"vine-template-rpc/pkg/pbtrans"
)

// Page 分页信息
type Page struct {
	PageNum  int         `json:"pageNum,omitempty" form:"pageNum"`
	PageSize int         `json:"pageSize,omitempty" form:"pageSize"`
	Total    int         `json:"total,omitempty" form:"total"`
	List     interface{} `json:"list"`
}

type PageHelper interface {
	Paginator() Page
}

func NewMemPage(list interface{}) *MemPage {
	return &MemPage{
		List: list,
	}
}

type MemPage struct {
	List     interface{}
	PageNum  int
	PageSize int
}

func (mp *MemPage) Paginator(pageNum, pageSize int32) *pb.Page {
	s := reflect.ValueOf(mp.List)

	reply := pb.Page{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    int32(s.Len()),
		List:     nil,
	}
	start := (pageNum - 1) * pageSize
	var end int
	if int(pageNum*pageSize) > s.Len() || pageNum*pageSize == 0 {
		end = s.Len()
	} else {
		end = int(pageNum * pageSize)
	}

	for i := int(start); i < end; i++ {
		var tmpVar *structpb.Struct
		pbtrans.InterfaceToPbStruct(s.Index(int(i)).Interface(), &tmpVar)
		reply.List = append(reply.List, tmpVar)
	}
	return &reply
}

type DbPage struct {
}

func (dp *DbPage) Paginator() Page {
	return Page{}
}
