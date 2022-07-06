package main

import "fmt"

// ICreateServer 创建云主机
type ICreateServer interface {
	CreateServer(cpu, mem float64)
}

// AWSClient aws sdk
type AWSClient struct{}

// RunInstance 启动实例
func (c *AWSClient) RunInstance(cpu, mem float64) {
	fmt.Printf("aws client run success, cpu： %f, mem: %f\n", cpu, mem)
}

// AwsClientAdapter 适配器
type AwsClientAdapter struct {
	Client AWSClient
}

// CreateServer 启动实例
func (a *AwsClientAdapter) CreateServer(cpu, mem float64) {
	a.Client.RunInstance(cpu, mem)
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) {
	fmt.Printf("aws client run success, cpu： %d, mem: %d\n", cpu, mem)
}

// AliyunClientAdapter 适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

// CreateServer 启动实例
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) {
	a.Client.CreateServer(int(cpu), int(mem))
}
func main() {
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}

	a.CreateServer(1.0, 2.0)

	var b ICreateServer = &AwsClientAdapter{
		Client: AWSClient{},
	}

	b.CreateServer(1.0, 2.0)
}
