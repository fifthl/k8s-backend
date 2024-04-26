package config

import "time"

const (
	//监听地址
	ListenAddr = "0.0.0.0:9090"
	WsAddr     = "0.0.0.0:8081"
	//kubeconfig路径
	//Kubeconfig = "F:\\goproject\\config"  windows路径
	Kubeconfigs = `{"TST-1":"/opt/config"}`
	//pod日志tail显示行数
	PodLogTailLine = 2000
	//登录账号密码
	AdminUser = "admin"
	AdminPwd  = "123456"

	//数据库配置
	DbType = "mysql"
	DbHost = "10.40.0.10"
	DbPort = 3306
	DbName = "k8s_demo"
	DbUser = "root"
	DbPwd  = "Abc123456."
	//打印mysql debug sql日志
	LogMode = false
	//连接池配置
	MaxIdleConns = 10               //最大空闲连接
	MaxOpenConns = 100              //最大连接数
	MaxLifeTime  = 30 * time.Second //最大生存时间
	//helm配置
	UploadPath = "/Users/adoo/chart"
)
