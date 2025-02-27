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

// Code generated by Kitex v0.12.0. DO NOT EDIT.

package versionservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	version "github.com/west2-online/fzuhelper-server/kitex_gen/version"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newVersionServiceLoginArgs,
		newVersionServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadVersion": kitex.NewMethodInfo(
		uploadVersionHandler,
		newVersionServiceUploadVersionArgs,
		newVersionServiceUploadVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadParams": kitex.NewMethodInfo(
		uploadParamsHandler,
		newVersionServiceUploadParamsArgs,
		newVersionServiceUploadParamsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DownloadReleaseApk": kitex.NewMethodInfo(
		downloadReleaseApkHandler,
		newVersionServiceDownloadReleaseApkArgs,
		newVersionServiceDownloadReleaseApkResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DownloadBetaApk": kitex.NewMethodInfo(
		downloadBetaApkHandler,
		newVersionServiceDownloadBetaApkArgs,
		newVersionServiceDownloadBetaApkResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetReleaseVersion": kitex.NewMethodInfo(
		getReleaseVersionHandler,
		newVersionServiceGetReleaseVersionArgs,
		newVersionServiceGetReleaseVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetBetaVersion": kitex.NewMethodInfo(
		getBetaVersionHandler,
		newVersionServiceGetBetaVersionArgs,
		newVersionServiceGetBetaVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetSetting": kitex.NewMethodInfo(
		getSettingHandler,
		newVersionServiceGetSettingArgs,
		newVersionServiceGetSettingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetTest": kitex.NewMethodInfo(
		getTestHandler,
		newVersionServiceGetTestArgs,
		newVersionServiceGetTestResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetCloud": kitex.NewMethodInfo(
		getCloudHandler,
		newVersionServiceGetCloudArgs,
		newVersionServiceGetCloudResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetCloud": kitex.NewMethodInfo(
		setCloudHandler,
		newVersionServiceSetCloudArgs,
		newVersionServiceSetCloudResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetDump": kitex.NewMethodInfo(
		getDumpHandler,
		newVersionServiceGetDumpArgs,
		newVersionServiceGetDumpResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AndroidGetVersion": kitex.NewMethodInfo(
		androidGetVersionHandler,
		newVersionServiceAndroidGetVersionArgs,
		newVersionServiceAndroidGetVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	versionServiceServiceInfo                = NewServiceInfo()
	versionServiceServiceInfoForClient       = NewServiceInfoForClient()
	versionServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return versionServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return versionServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return versionServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "VersionService"
	handlerType := (*version.VersionService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "version",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.0",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceLoginArgs)
	realResult := result.(*version.VersionServiceLoginResult)
	success, err := handler.(version.VersionService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceLoginArgs() interface{} {
	return version.NewVersionServiceLoginArgs()
}

func newVersionServiceLoginResult() interface{} {
	return version.NewVersionServiceLoginResult()
}

func uploadVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceUploadVersionArgs)
	realResult := result.(*version.VersionServiceUploadVersionResult)
	success, err := handler.(version.VersionService).UploadVersion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceUploadVersionArgs() interface{} {
	return version.NewVersionServiceUploadVersionArgs()
}

func newVersionServiceUploadVersionResult() interface{} {
	return version.NewVersionServiceUploadVersionResult()
}

func uploadParamsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceUploadParamsArgs)
	realResult := result.(*version.VersionServiceUploadParamsResult)
	success, err := handler.(version.VersionService).UploadParams(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceUploadParamsArgs() interface{} {
	return version.NewVersionServiceUploadParamsArgs()
}

func newVersionServiceUploadParamsResult() interface{} {
	return version.NewVersionServiceUploadParamsResult()
}

func downloadReleaseApkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceDownloadReleaseApkArgs)
	realResult := result.(*version.VersionServiceDownloadReleaseApkResult)
	success, err := handler.(version.VersionService).DownloadReleaseApk(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceDownloadReleaseApkArgs() interface{} {
	return version.NewVersionServiceDownloadReleaseApkArgs()
}

func newVersionServiceDownloadReleaseApkResult() interface{} {
	return version.NewVersionServiceDownloadReleaseApkResult()
}

func downloadBetaApkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceDownloadBetaApkArgs)
	realResult := result.(*version.VersionServiceDownloadBetaApkResult)
	success, err := handler.(version.VersionService).DownloadBetaApk(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceDownloadBetaApkArgs() interface{} {
	return version.NewVersionServiceDownloadBetaApkArgs()
}

func newVersionServiceDownloadBetaApkResult() interface{} {
	return version.NewVersionServiceDownloadBetaApkResult()
}

func getReleaseVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceGetReleaseVersionArgs)
	realResult := result.(*version.VersionServiceGetReleaseVersionResult)
	success, err := handler.(version.VersionService).GetReleaseVersion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceGetReleaseVersionArgs() interface{} {
	return version.NewVersionServiceGetReleaseVersionArgs()
}

func newVersionServiceGetReleaseVersionResult() interface{} {
	return version.NewVersionServiceGetReleaseVersionResult()
}

func getBetaVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceGetBetaVersionArgs)
	realResult := result.(*version.VersionServiceGetBetaVersionResult)
	success, err := handler.(version.VersionService).GetBetaVersion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceGetBetaVersionArgs() interface{} {
	return version.NewVersionServiceGetBetaVersionArgs()
}

func newVersionServiceGetBetaVersionResult() interface{} {
	return version.NewVersionServiceGetBetaVersionResult()
}

func getSettingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceGetSettingArgs)
	realResult := result.(*version.VersionServiceGetSettingResult)
	success, err := handler.(version.VersionService).GetSetting(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceGetSettingArgs() interface{} {
	return version.NewVersionServiceGetSettingArgs()
}

func newVersionServiceGetSettingResult() interface{} {
	return version.NewVersionServiceGetSettingResult()
}

func getTestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceGetTestArgs)
	realResult := result.(*version.VersionServiceGetTestResult)
	success, err := handler.(version.VersionService).GetTest(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceGetTestArgs() interface{} {
	return version.NewVersionServiceGetTestArgs()
}

func newVersionServiceGetTestResult() interface{} {
	return version.NewVersionServiceGetTestResult()
}

func getCloudHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceGetCloudArgs)
	realResult := result.(*version.VersionServiceGetCloudResult)
	success, err := handler.(version.VersionService).GetCloud(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceGetCloudArgs() interface{} {
	return version.NewVersionServiceGetCloudArgs()
}

func newVersionServiceGetCloudResult() interface{} {
	return version.NewVersionServiceGetCloudResult()
}

func setCloudHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceSetCloudArgs)
	realResult := result.(*version.VersionServiceSetCloudResult)
	success, err := handler.(version.VersionService).SetCloud(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceSetCloudArgs() interface{} {
	return version.NewVersionServiceSetCloudArgs()
}

func newVersionServiceSetCloudResult() interface{} {
	return version.NewVersionServiceSetCloudResult()
}

func getDumpHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceGetDumpArgs)
	realResult := result.(*version.VersionServiceGetDumpResult)
	success, err := handler.(version.VersionService).GetDump(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceGetDumpArgs() interface{} {
	return version.NewVersionServiceGetDumpArgs()
}

func newVersionServiceGetDumpResult() interface{} {
	return version.NewVersionServiceGetDumpResult()
}

func androidGetVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*version.VersionServiceAndroidGetVersionArgs)
	realResult := result.(*version.VersionServiceAndroidGetVersionResult)
	success, err := handler.(version.VersionService).AndroidGetVersion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVersionServiceAndroidGetVersionArgs() interface{} {
	return version.NewVersionServiceAndroidGetVersionArgs()
}

func newVersionServiceAndroidGetVersionResult() interface{} {
	return version.NewVersionServiceAndroidGetVersionResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *version.LoginRequest) (r *version.LoginResponse, err error) {
	var _args version.VersionServiceLoginArgs
	_args.Req = req
	var _result version.VersionServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadVersion(ctx context.Context, req *version.UploadRequest) (r *version.UploadResponse, err error) {
	var _args version.VersionServiceUploadVersionArgs
	_args.Req = req
	var _result version.VersionServiceUploadVersionResult
	if err = p.c.Call(ctx, "UploadVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadParams(ctx context.Context, req *version.UploadParamsRequest) (r *version.UploadParamsResponse, err error) {
	var _args version.VersionServiceUploadParamsArgs
	_args.Req = req
	var _result version.VersionServiceUploadParamsResult
	if err = p.c.Call(ctx, "UploadParams", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadReleaseApk(ctx context.Context, req *version.DownloadReleaseApkRequest) (r *version.DownloadReleaseApkResponse, err error) {
	var _args version.VersionServiceDownloadReleaseApkArgs
	_args.Req = req
	var _result version.VersionServiceDownloadReleaseApkResult
	if err = p.c.Call(ctx, "DownloadReleaseApk", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadBetaApk(ctx context.Context, req *version.DownloadBetaApkRequest) (r *version.DownloadBetaApkResponse, err error) {
	var _args version.VersionServiceDownloadBetaApkArgs
	_args.Req = req
	var _result version.VersionServiceDownloadBetaApkResult
	if err = p.c.Call(ctx, "DownloadBetaApk", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetReleaseVersion(ctx context.Context, req *version.GetReleaseVersionRequest) (r *version.GetReleaseVersionResponse, err error) {
	var _args version.VersionServiceGetReleaseVersionArgs
	_args.Req = req
	var _result version.VersionServiceGetReleaseVersionResult
	if err = p.c.Call(ctx, "GetReleaseVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetBetaVersion(ctx context.Context, req *version.GetBetaVersionRequest) (r *version.GetBetaVersionResponse, err error) {
	var _args version.VersionServiceGetBetaVersionArgs
	_args.Req = req
	var _result version.VersionServiceGetBetaVersionResult
	if err = p.c.Call(ctx, "GetBetaVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSetting(ctx context.Context, req *version.GetSettingRequest) (r *version.GetSettingResponse, err error) {
	var _args version.VersionServiceGetSettingArgs
	_args.Req = req
	var _result version.VersionServiceGetSettingResult
	if err = p.c.Call(ctx, "GetSetting", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTest(ctx context.Context, req *version.GetTestRequest) (r *version.GetTestResponse, err error) {
	var _args version.VersionServiceGetTestArgs
	_args.Req = req
	var _result version.VersionServiceGetTestResult
	if err = p.c.Call(ctx, "GetTest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCloud(ctx context.Context, req *version.GetCloudRequest) (r *version.GetCloudResponse, err error) {
	var _args version.VersionServiceGetCloudArgs
	_args.Req = req
	var _result version.VersionServiceGetCloudResult
	if err = p.c.Call(ctx, "GetCloud", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetCloud(ctx context.Context, req *version.SetCloudRequest) (r *version.SetCloudResponse, err error) {
	var _args version.VersionServiceSetCloudArgs
	_args.Req = req
	var _result version.VersionServiceSetCloudResult
	if err = p.c.Call(ctx, "SetCloud", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDump(ctx context.Context, req *version.GetDumpRequest) (r *version.GetDumpResponse, err error) {
	var _args version.VersionServiceGetDumpArgs
	_args.Req = req
	var _result version.VersionServiceGetDumpResult
	if err = p.c.Call(ctx, "GetDump", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AndroidGetVersion(ctx context.Context, req *version.AndroidGetVersioneRequest) (r *version.AndroidGetVersionResponse, err error) {
	var _args version.VersionServiceAndroidGetVersionArgs
	_args.Req = req
	var _result version.VersionServiceAndroidGetVersionResult
	if err = p.c.Call(ctx, "AndroidGetVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
