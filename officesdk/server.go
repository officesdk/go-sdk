package officesdk

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/cetus/l"
	"github.com/gotomicro/ego/core/elog"
)

// Server 服务配置
type Server struct {
	config Config
	engine *gin.Engine
	root   gin.IRouter
}

// NewServer 创建服务
func NewServer(config Config, e *gin.Engine) {
	if config.PreviewProvider == nil {
		log.Panic("PreviewProvider must not nil")
	}
	srv := &Server{
		engine: e,
		config: config,
	}
	if config.Prefix == "" {
		srv.root = srv.engine
	} else {
		srv.root = srv.engine.Group(config.Prefix)
	}

	// 注册路由到主engine
	srv.registerRoutes(srv.root)
}

// attachHeaders 将请求头中的信息附加到上下文中
func attachHeaders(c *gin.Context) {
	if v, err := url.ParseQuery(c.Request.Header.Get("X-User-Query")); err == nil {
		c.Header("X-User-Query", v.Encode())
	}
}

// wrapHandlerFunc 包装handler
func (srv *Server) wrapHandlerFunc(f func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		data, err := f(c)
		cost := time.Since(begin)
		if err == nil {
			if cost > time.Second {
				elog.Warn("wrapHandlerFunc", l.S("method", c.Request.Method), l.S("uri", c.Request.URL.RequestURI()), l.S("cost", cost.String()))
			}
			c.JSON(http.StatusOK, &Res{Code: OK, Data: data})
			return
		}
		var e *Error
		if errors.As(err, &e) {
			elog.Error("wrapHandlerFunc", l.S("method", c.Request.Method), l.S("uri", c.Request.URL.RequestURI()), l.I("code", e.Code()), l.E(e), l.S("cost", cost.String()))
			c.JSON(e.HttpStatusCode(), &Res{Code: e.Code(), Message: e.Message()})
			return
		}
		elog.Error("wrapHandlerFunc", l.S("method", c.Request.Method), l.S("uri", c.Request.URL.RequestURI()), l.E(e), l.S("cost", cost.String()))
		c.JSON(http.StatusInternalServerError, &Res{Code: Failed, Message: err.Error()})
	}
}

func (srv *Server) registerRoutes(router gin.IRouter) {
	rg := router.Group("/v1/thirdparty")

	// 预览接口
	rg.GET("/files/:file_id", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
		attachHeaders(c)
		return srv.config.GetFile(c, c.Param("file_id"))
	}))

	rg.GET("/files/:file_id/download", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
		attachHeaders(c)
		// 调用获取下载地址的逻辑
		return srv.config.GetFileDownload(c, c.Param("file_id"))
	}))

	rg.GET("/files/:file_id/watermark", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
		attachHeaders(c)
		// 调用获取水印信息的逻辑
		return srv.config.GetFileWatermark(c, c.Param("file_id"))
	}))

	// 编辑接口
	if srv.config.EditProvider != nil {
		rg.POST("/files/:file_id/content/upload-url", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用获取上传 URL 的逻辑
			return srv.config.GetUploadURL(c, c.Param("file_id"))
		}))

		rg.POST("/files/:file_id/content/upload-completion", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用完成上传的逻辑
			return srv.config.CompleteUpload(c, c.Param("file_id"))
		}))

		rg.GET("/files/:file_id/content/url", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用获取下载 URL 的逻辑
			return srv.config.GetDownloadURL(c, c.Param("file_id"))
		}))

		rg.POST("/files/:file_id/assets/upload-url", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用获取上传 URL 的逻辑
			return srv.config.GetAssetUploadURL(c, c.Param("file_id"))
		}))

		rg.POST("/files/:file_id/assets/upload-completion", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用完成上传的逻辑
			return srv.config.AssetCompleteUpload(c, c.Param("file_id"))
		}))

		rg.GET("/files/:file_id/assets/url", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用获取下载 URL 的逻辑
			return srv.config.GetAssetDownloadURL(c, c.Param("file_id"))
		}))
	}

	// AI 接口
	if srv.config.AIProvider != nil {

		rg.GET("/chat/ai-config", srv.wrapHandlerFunc(func(c *gin.Context) (any, error) {
			attachHeaders(c)
			// 调用完成上传的逻辑
			return srv.config.AIConfig(c)
		}))

	}
}
