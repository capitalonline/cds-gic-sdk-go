[English](./README.md) | 简体中文

<p align="center">
<a href=" https://www.alibabacloud.com"><img src="https://www.capitalonline.net/templets/default/icon/logo_header.png"></a>
</p>

<h1 align="center">CapitalOnline Cloud SDK for Go</h1>

欢迎使用`CapitalOnline Cloud SDK for Go`，它可以管理[首都在线](https://www.capitalonline.net)多个全球服务，如`首云全球云服务器`、`首云块存储`、`首云全球互通网络`等，帮您轻松管理所有在线资源。基于首云官方[OpenAPI文档](https://github.com/capitalonline/openapi/blob/master/README.md)。

## 功能特点

你可以通过SDK管理各种资源，[这里](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md)可以看到所有可用接口清单。下面列出部分接口列表：

- [X] 实例管理
  - [X] CreateInstance
  - [X] DeleteInstance
  - [X] StopInstance
  - [X] RebootInstance
  - [X] ModifyInstanceChargeType
- [X] 虚拟数据中心管理
  - [X] DescribeVdc
  - [X] CreateVdc
  - [X] DeleteVdc
  - [X] CreatePublicNetwork
  - [X] CreatePrivateNetwork
- [X] 安全组管理
  - [X] CreateSecurityGroup
  - [X] DeleteSecurityGroup
  - [X] ForceDeleteSecurityGroup
  - [X] DescribeSecurityGroupAttribute
  - [X] ModifySecurityGroupAttribute
- [X] 还有更多

## 使用指南

使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/capitalonline/cds-gic-sdk-go
```

## 使用示例

```go
    	// Init a credential with Access Key Id and Secret Access Key
	// You can apply them from the CDS web portal
	credential := common.NewCredential(
		os.Getenv("CDS_SECRET_ID"),
		os.Getenv("CDS_SECRET_KEY"),
	)

	// init a client profile with method type
	cpf := profile.NewClientProfile()

	// Example: Get vdc
	cpf.HttpProfile.ReqMethod = "GET"
	vdcClient, _ := vdc.NewClient(credential, "", cpf)
	descVdcRequest := vdc.DescribeVdcRequest()

	// comment out the line below, you can get all the vdc data
	descVdcRequest.RegionId = common.StringPtr(Shanghai)
	descVdcResponse, err := vdcClient.DescribeVdc(descVdcRequest)
	if err != nil {
		fmt.Println("API request fail:", err.Error())
	} else {
		fmt.Println(descVdcResponse.ToJsonString())
	}

	// Example: Create vdc
	cpf.HttpProfile.ReqMethod = "POST"
	vdcClient, _ = vdc.NewClient(credential, "", cpf)
	createVdcRequest := vdc.NewAddVdcRequest()
	createVdcRequest.RegionId = common.StringPtr(Beijing)
	createVdcRequest.VdcName = common.StringPtr("beijing-vdc")
	createVdcRequest.PublicNetwork = &vdc.PublicNetwork{
		// the account of public IP
		IPNum: common.IntPtr(8),
		// the bandwidth of public network
		Qos:           common.IntPtr(10),
		Name:          common.StringPtr("pubnet"),
		BillingMethod: common.StringPtr("Bandwidth"),
		Type:          common.StringPtr("Bandwidth_BGP"),
	}
	createVdcResponse, err := vdcClient.CreateVdc(createVdcRequest)
	if err != nil {
		fmt.Println("API request fail:", err.Error())
	} else {
		fmt.Printf("Task: %v, code: %v", *createVdcResponse.TaskId, *createVdcResponse.Code)
	}
```

更多示例参见 [example](./example) 目录。

## 如何贡献

欢迎提交 Issue 或 Pull Request。

## 相关参考

- [CDS OpenAPI Explorer](https://github.com/capitalonline/openapi)

## 许可证

[Apache License v2.0](./LICENSE)
