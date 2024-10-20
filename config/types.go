/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type snowflake struct {
	WorkerID      int64 `mapstructure:"worker-id"`
	DatancenterID int64 `mapstructure:"datancenter-id"`
}

type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}

type mySQL struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type jaeger struct {
	Addr string
}

type etcd struct {
	Addr string
}

type rabbitMQ struct {
	Addr     string
	Username string
	Password string
}

type redis struct {
	Addr     string
	Password string
}

type oss struct {
	Endpoint        string
	AccessKeyID     string `mapstructure:"accessKey-id"`
	AccessKeySecret string `mapstructure:"accessKey-secret"`
	BucketName      string
	MainDirectory   string `mapstructure:"main-directory"`
}

type elasticsearch struct {
	Addr string
	Host string
}

type kafka struct {
	Address string
	Network string
}

type defaultUser struct {
	Account  string `mapstructure:"account"`
	Password string `mapstructure:"password"`
}
type upcloud struct {
	Service    string
	User       string
	Pass       string
	DomainName string
	Path       string
}
type config struct {
	Server        server
	Snowflake     snowflake
	MySQL         mySQL
	Jaeger        jaeger
	Etcd          etcd
	RabbitMQ      rabbitMQ
	Redis         redis
	OSS           oss
	Elasticsearch elasticsearch
	Kafka         kafka
	DefaultUser   defaultUser
	Upcloud       upcloud
}
