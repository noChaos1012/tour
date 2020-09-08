1、构建指令 
`go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"`

2、文件部署时，需要将configs文件夹、storage文件夹拷贝在同一目录下运行

3、`swag init` 根据注解更新swagger文档

4、执行程序 `./blog-service -version `可查看构建版本