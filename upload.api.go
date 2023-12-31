// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.5). DO NOT EDIT.
// source: bilibili-webapp/upload.proto

package bilibili_webapp

import (
	context "context"
	fmt "fmt"
	io "io"
	json "encoding/json"
	bytes "bytes"
	http "net/http"
	strings "strings"
	url "net/url"
	multipart "mime/multipart"
)
// Reference imports to suppress errors if they are not otherwise used.
var _ = context.Background
var _ = http.NewRequest
var _ = io.Copy
var _ = bytes.Compare
var _ = json.Marshal
var _ = strings.Compare
var _ = fmt.Errorf
var _ = url.Parse
var _ = multipart.ErrMessageTooLarge


// Client API for Upload service

type UploadService interface {
	// InitPart  文件上传预处理 https://openhome.bilibili.com/doc/4/0c532c6a-e6fb-0aff-8021-905ae2409095
	InitPart(ctx context.Context, in *InitPartReq, opts ...Option) (*InitPartRes, error)
	// UploadPart  文件分片上传 域名用 https://openupos.bilivideo.com https://openhome.bilibili.com/doc/4/733a520a-c50f-7bb4-17cb-35338ba20500
	UploadPart(ctx context.Context, in *UploadPartReq, opts ...Option) (*UploadPartRes, error)
	// CompletePart  分片合并 https://openhome.bilibili.com/doc/4/0828e499-38d8-9e58-2a70-a7eaebf9dd64
	CompletePart(ctx context.Context, in *CompletePartReq, opts ...Option) (*CompletePartRes, error)
	// UploadCover  封面上传 https://openhome.bilibili.com/doc/4/8243399e-50e3-4058-7f01-1ebe4c632cf8
	UploadCover(ctx context.Context, in *UploadCoverReq, opts ...Option) (*UploadCoverRes, error)
	// Upload  单次上传 域名用 https://openupos.bilivideo.com https://openhome.bilibili.com/doc/4/f22a0eee-c92d-0f1d-f69c-be170cf533c7
	Upload(ctx context.Context, in *UploadReq, opts ...Option) (*UploadRes, error)
	// UploadImage  上传图片 https://openhome.bilibili.com/doc/4/0eaa4d3e-c4c0-f874-6f3c-e083aa939a1b
	UploadImage(ctx context.Context, in *UploadImageReq, opts ...Option) (*UploadImageRes, error)
}

type uploadService struct {
	// opts
	opts *Options
}

func NewUploadService(opts ...Option) UploadService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://member.bilibili.com"
	}
	return &uploadService {
		opts: opt,
	}
}


func (c *uploadService) InitPart(ctx context.Context, in *InitPartReq, opts ...Option) (*InitPartRes, error) {
	var res InitPartRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/arcopen/fn/archive/video/init", opt.addr)

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAccessToken() != "" {
		params.Add("access_token", fmt.Sprintf("%v", in.GetAccessToken()))
	}
	if in.GetClientId() != "" {
		params.Add("client_id", fmt.Sprintf("%v", in.GetClientId()))
	}	
	req.URL.RawQuery = params.Encode()
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *uploadService) UploadPart(ctx context.Context, in *UploadPartReq, opts ...Option) (*UploadPartRes, error) {
	var res UploadPartRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/video/v2/part/upload", opt.addr)

	// body
	req, err := http.NewRequest("POST", rawURL, nil)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetPartNumber() != 0 {
		params.Add("part_number", fmt.Sprintf("%v", in.GetPartNumber()))
	}
	if in.GetUploadToken() != "" {
		params.Add("upload_token", fmt.Sprintf("%v", in.GetUploadToken()))
	}	
	req.URL.RawQuery = params.Encode()
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *uploadService) CompletePart(ctx context.Context, in *CompletePartReq, opts ...Option) (*CompletePartRes, error) {
	var res CompletePartRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/arcopen/fn/archive/video/complete", opt.addr)

	// body
	req, err := http.NewRequest("POST", rawURL, nil)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetUploadToken() != "" {
		params.Add("upload_token", fmt.Sprintf("%v", in.GetUploadToken()))
	}	
	req.URL.RawQuery = params.Encode()
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *uploadService) UploadCover(ctx context.Context, in *UploadCoverReq, opts ...Option) (*UploadCoverRes, error) {
	var res UploadCoverRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/arcopen/fn/archive/cover/upload", opt.addr)

	// body
	body := new(bytes.Buffer)
	bodyForms := multipart.NewWriter(body) 
	if in.GetBody() != nil && in.GetBody().GetFile() != nil {
		bodyForms.WriteField("file", fmt.Sprintf("%v", in.GetBody().GetFile()))
	}
	defer func() { _ =  bodyForms.Close() } ()
	headers["Content-Type"] = "multipart/form-data"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAccessToken() != "" {
		params.Add("access_token", fmt.Sprintf("%v", in.GetAccessToken()))
	}
	if in.GetClientId() != "" {
		params.Add("client_id", fmt.Sprintf("%v", in.GetClientId()))
	}	
	req.URL.RawQuery = params.Encode()
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *uploadService) Upload(ctx context.Context, in *UploadReq, opts ...Option) (*UploadRes, error) {
	var res UploadRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/video/v2/upload", opt.addr)

	// body
	body := bytes.NewReader(in.GetBody())
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetUploadToken() != "" {
		params.Add("upload_token", fmt.Sprintf("%v", in.GetUploadToken()))
	}	
	req.URL.RawQuery = params.Encode()
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *uploadService) UploadImage(ctx context.Context, in *UploadImageReq, opts ...Option) (*UploadImageRes, error) {
	var res UploadImageRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/arcopen/fn/article/upload/image", opt.addr)

	// body
	body := new(bytes.Buffer)
	bodyForms := multipart.NewWriter(body) 
	if in.GetBody() != nil && in.GetBody().GetFile() != nil {
		bodyForms.WriteField("file", fmt.Sprintf("%v", in.GetBody().GetFile()))
	}
	if in.GetBody() != nil && in.GetBody().GetWatermark() {
		bodyForms.WriteField("watermark", fmt.Sprintf("%v", in.GetBody().GetWatermark()))
	}
	defer func() { _ =  bodyForms.Close() } ()
	headers["Content-Type"] = "multipart/form-data"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAccessToken() != "" {
		params.Add("access_token", fmt.Sprintf("%v", in.GetAccessToken()))
	}
	if in.GetClientId() != "" {
		params.Add("client_id", fmt.Sprintf("%v", in.GetClientId()))
	}	
	req.URL.RawQuery = params.Encode()
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}
