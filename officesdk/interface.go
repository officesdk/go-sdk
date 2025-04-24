package officesdk

import "github.com/gin-gonic/gin"

type FileProvider interface {
	GetFile(c *gin.Context, fileId string) (*FileResponse, error)
	GetFileDownload(c *gin.Context, fileId string) (*DownloadResponse, error)
	GetFileWatermark(c *gin.Context, fileId string) (*WatermarkResponse, error)

	GetUploadURL(c *gin.Context, fileId string) (*UploadURLResponse, error)
	CompleteUpload(c *gin.Context, fileId string) (*UploadCompletionResponse, error)
	GetDownloadURL(c *gin.Context, fileId string) (*DownloadResponse, error)

	GetAssetUploadURL(c *gin.Context, fileId string) (*AssetUploadURLResponse, error)
	AssetCompleteUpload(c *gin.Context, fileId string) (*UploadCompletionResponse, error)
	GetAssetDownloadURL(c *gin.Context, fileId string) (*DownloadResponse, error)

	VerifyFile(c *gin.Context, fileId string) (*VerifyResponse, error)
}

type AIProvider interface {
	AIConfig(c *gin.Context) (*AIConfigResponse, error)
}
