package officesdk

const ()

// Config 配置
type Config struct {
	// 预览接口
	PreviewProvider
	// 编辑接口
	EditProvider
	// AI
	AIProvider
	// 路由前缀（上传用）
	Prefix string
}
