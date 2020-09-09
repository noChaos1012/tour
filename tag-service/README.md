##操作指令
###1、创建pb文件
`protoc --go_out=plugins=grpc:. ./proto/*.proto`

###2、使用grpcurl进行接口调试
`grpcurl -plaintext localhost:8999 list`
`grpcurl -plaintext localhost:8999 list proto.TagService`
`grpcurl -plaintext -d '{"name":"Go"}' localhost:8999  proto.TagService.GetTagList`

##工具使用
###1、grpccurl 
grpc 接口调试工具
`go get github.com/fullstorydev/grpcurl`
`go install github.com/fullstorydev/grpcurl/cmd/grpcurl`
