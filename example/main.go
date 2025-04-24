package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/officesdk/go-sdk/officesdk"
)

func main() {
	// 初始化路由
	e := gin.Default()

	officesdk.NewServer(officesdk.Config{
		FileProvider: &FileProvider{},
		AIProvider:   &AIProvider{},
		Prefix:       "/api",
	}, e)

	_ = e.Run(":8080")
}

// PreviewProvider 实现预览相关接口
type FileProvider struct{}

func (p *FileProvider) GetFile(c *gin.Context, fileId string) (*officesdk.FileResponse, error) {
	return &officesdk.FileResponse{
		ID:         fileId,
		Name:       "example.docx",
		Version:    1,
		CreateTime: time.Now().Unix(),
		ModifyTime: time.Now().Unix(),
		CreatorID:  "creator123",
		ModifierID: "modifier123",
		FromSDK:    false,
	}, nil
}

func (p *FileProvider) GetFileDownload(c *gin.Context, fileId string) (*officesdk.DownloadResponse, error) {
	return &officesdk.DownloadResponse{
		URL: "https://example.com/download/" + fileId,
		Headers: map[string]string{
			"Authorization": "Bearer token",
		},
	}, nil
}

func (p *FileProvider) GetFileWatermark(c *gin.Context, fileId string) (*officesdk.WatermarkResponse, error) {
	return &officesdk.WatermarkResponse{
		Type:       1,
		Value:      "示例水印",
		FillStyle:  "rgba(192, 192, 192, 0.6)",
		Font:       "bold 20px Serif",
		Rotate:     -0.7853982,
		Horizontal: 10,
		Vertical:   20,
	}, nil
}

func (p *FileProvider) GetUploadURL(c *gin.Context, fileId string) (*officesdk.UploadURLResponse, error) {
	return &officesdk.UploadURLResponse{
		URL:    "https://example.com/upload/" + fileId,
		Method: "PUT",
		Headers: map[string]string{
			"Authorization": "Bearer token",
		},
		Params: map[string]string{
			"param1": "value1",
		},
		CompletionParams: map[string]string{
			"completion_param1": "value1",
		},
	}, nil
}

func (p *FileProvider) CompleteUpload(c *gin.Context, fileId string) (*officesdk.UploadCompletionResponse, error) {
	return &officesdk.UploadCompletionResponse{
		ID:         fileId,
		Version:    1,
		CreateTime: time.Now().Unix(),
		ModifyTime: time.Now().Unix(),
		CreatorID:  "creator123",
		ModifierID: "modifier123",
	}, nil
}

func (p *FileProvider) GetDownloadURL(c *gin.Context, fileId string) (*officesdk.DownloadResponse, error) {
	return &officesdk.DownloadResponse{
		URL: "https://example.com/download/" + fileId,
		Headers: map[string]string{
			"Authorization": "Bearer token",
		},
	}, nil
}

func (p *FileProvider) GetAssetUploadURL(c *gin.Context, fileId string) (*officesdk.AssetUploadURLResponse, error) {
	return &officesdk.AssetUploadURLResponse{
		URL:          "https://example.com/upload/" + fileId,
		Method:       "PUT",
		FileFieldKey: "file",
		Headers: map[string]string{
			"Authorization": "Bearer token",
		},
		Params: map[string]string{
			"param1": "value1",
		},
		CompletionParams: map[string]string{
			"completion_param1": "value1",
		},
	}, nil
}

func (p *FileProvider) AssetCompleteUpload(c *gin.Context, fileId string) (*officesdk.UploadCompletionResponse, error) {
	return &officesdk.UploadCompletionResponse{
		ID:         fileId,
		Version:    1,
		CreateTime: time.Now().Unix(),
		ModifyTime: time.Now().Unix(),
		CreatorID:  "creator123",
		ModifierID: "modifier123",
	}, nil
}

func (p *FileProvider) GetAssetDownloadURL(c *gin.Context, fileId string) (*officesdk.DownloadResponse, error) {
	return &officesdk.DownloadResponse{
		URL: "https://example.com/download/" + fileId,
		Headers: map[string]string{
			"Authorization": "Bearer token",
		},
	}, nil
}

func (p *FileProvider) VerifyFile(c *gin.Context, fileId string) (*officesdk.VerifyResponse, error) {
	return &officesdk.VerifyResponse{
		CurrentUserInfo: officesdk.UserInfo{
			ID:     "1",
			Name:   "名称",
			Email:  "a@b.com",
			Avatar: "",
		},
	}, nil
}

// AIProvider 实现 AI 相关接口
type AIProvider struct{}

func (p *AIProvider) AIConfig(c *gin.Context) (*officesdk.AIConfigResponse, error) {
	return &officesdk.AIConfigResponse{
		LLMList: []officesdk.LLMConfig{
			{
				Name:           "GPT-4",
				BaseURL:        "https://api.example.com/gpt-4",
				TextModel:      "gpt-4",
				Token:          "your_api_key",
				InputMaxToken:  2048,
				OutputMaxToken: 150,
				ProxyURL:       "https://proxy.example.com",
				Subservice:     "default",
			},
		},
	}, nil
}
