package officesdk

const ()

// Config 配置
type Config struct {
	// 文件接口
	FileProvider
	// AI
	AIProvider
	// 路由前缀（上传用）
	Prefix string
}

// Res 返回参数结构体
type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}

type FileResponse struct {
	ID      string `json:"id"`       // 文档 ID，必须与传入的 file_id 一致，最大长度 47
	Name    string `json:"name"`     // 带有后缀的文件名，如：hi.docx
	Version uint32 `json:"version"`  // 文档版本号，无符号 int32 位，从 0 开始
	FromSDK bool   `json:"from_sdk"` // 等于 false，预览会先下载文件进行格式转换，再上传保存为 sdk 格式文件。在下次预览时 from_sdk 为 true

	CreateTime int64  `json:"create_time,omitempty"` // 文档创建时间戳，单位纪元秒
	ModifyTime int64  `json:"modify_time,omitempty"` // 文档最后修改时间戳，单位纪元秒
	CreatorID  string `json:"creator_id,omitempty"`  // 文档创建者 Id
	ModifierID string `json:"modifier_id,omitempty"` // 文档最后修改者 Id
}

type DownloadResponse struct {
	URL     string            `json:"url"`               // 文档下载地址，需确保外网可访问
	Headers map[string]string `json:"headers,omitempty"` // 请求文档下载地址所需的额外请求头
}

type WatermarkResponse struct {
	Type       int     `json:"type"`                  // 水印类型，0表示无水印，1表示文字水印
	Value      string  `json:"value,omitempty"`       // type = 1 时必须，水印显示的文字内容
	FillStyle  string  `json:"fill_style,omitempty"`  // 水印透明度，示例： rgba( 192, 192, 192, 0.6 )
	Font       string  `json:"font,omitempty"`        // 水印字体设置，示例： bold 20px Serif
	Rotate     float64 `json:"rotate,omitempty"`      // 水印旋转度，示例： -0.7853982
	Horizontal int     `json:"horizontal,omitempty"`  // type = 1 时必须，水印水平间距
	Vertical   int     `json:"vertical,omitempty"`    // type = 1 时必须，水印垂直间距
	TextAlign  string  `json:"text_align,omitempty"`  // 文字水平对其，参考值 center
	LineHeight int     `json:"line_height,omitempty"` // 行间距，参考值 2
}

type UploadURLResponse struct {
	URL              string            `json:"url"`                         // 上传文件的 URL
	Method           string            `json:"method"`                      // 上传文档的 HTTP Method，暂只支持 PUT
	Headers          map[string]string `json:"headers,omitempty"`           // 上传时需要携带的额外请求头
	Params           map[string]string `json:"params,omitempty"`            // 上传时需要携带的额外参数
	CompletionParams map[string]string `json:"completion_params,omitempty"` // 上传后请求完成上传接口需要原样带回的额外参数
}

type UploadCompletionResponse struct {
	ID         string `json:"id"`                    // 文档 ID，必须与传入的 file_id 一致
	Version    int    `json:"version"`               // 文档版本号，从 0 开始，每次保存后递增
	CreateTime int64  `json:"create_time,omitempty"` // 文档创建时间戳，单位纪元秒
	ModifyTime int64  `json:"modify_time,omitempty"` // 文档最后修改时间戳，单位纪元秒
	CreatorID  string `json:"creator_id,omitempty"`  // 文档创建者 Id
	ModifierID string `json:"modifier_id,omitempty"` // 文档最后修改者 Id
}

type AssetUploadURLResponse struct {
	URL              string            `json:"url"`                         // 上传文件的 URL
	Method           string            `json:"method"`                      // 上传文档的 HTTP Method，暂只支持 PUT
	FileFieldKey     string            `json:"file_field_key,omitempty"`    // 使用 http multipart/form-data  上传文件时，文件实体参数名，默认为 file
	Headers          map[string]string `json:"headers,omitempty"`           // 上传文件时需要携带的额外请求头
	Params           map[string]string `json:"params,omitempty"`            // 上传文件时需要携带的额外参数
	CompletionParams map[string]string `json:"completion_params,omitempty"` // 上传文件完成后请求完成上传接口时需要原样带回的额外参数
}

type LLMConfig struct {
	Name           string `json:"name"`                 // 配置名称，会展示在 AI 聊天框配置中
	BaseURL        string `json:"baseURL"`              // 大模型接口地址
	TextModel      string `json:"textModel"`            // 模型名称
	Token          string `json:"token"`                // API Key
	InputMaxToken  int    `json:"inputMaxToken"`        // 输入 token 限制
	OutputMaxToken int    `json:"outputMaxToken"`       // 输出 token 限制
	ProxyURL       string `json:"proxyURL,omitempty"`   // 代理 URL
	Subservice     string `json:"subservice,omitempty"` // 子服务
}

type AIConfigResponse struct {
	AiIcon  string      `json:"aiIcon,omitempty"` // AI 图标
	LLMList []LLMConfig `json:"llmList"`          // 大模型配置列表
}

// UserInfo 用户信息
type UserInfo struct {
	ID          string          `json:"id"`
	Name        string          `json:"name,omitempty"`
	Email       string          `json:"email,omitempty"`
	Avatar      string          `json:"avatar,omitempty"`
	Permissions map[string]bool `json:"permissions,omitempty"`
}

// VerifyResponse 文件鉴权返回结果
type VerifyResponse struct {
	CurrentUserInfo UserInfo `json:"currentUserInfo"`
}

// IsBrokenResponse 是否已终端返回结果
type IsBrokenResponse struct {
	Broken bool `json:"broken"`
}

// ChatConversation AI对话
type ChatConversation struct {
	ConversationId string          `json:"conversation_id"`
	System         string          `json:"system"`
	FileGuid       string          `json:"file_guid"`
	UserId         string          `json:"user_id"`
	Messages       []ChatMessageDO `json:"messages,omitempty"`
}

// ChatMessageDO AI对话消息
type ChatMessageDO struct {
	ChatMessage
	NeedAIChat bool `json:"need_ai_chat,omitempty"`
}

type ChatMessage struct {
	MessageId      string                 `json:"message_id,omitempty"`
	Role           string                 `json:"role,omitempty"`
	Type           string                 `json:"type,omitempty"`
	SentenceId     int                    `json:"sentence_id"`
	IsEnd          bool                   `json:"is_end,omitempty"`
	IsTruncated    bool                   `json:"is_truncated,omitempty"`
	Content        ChatContents           `json:"content,omitempty"`
	Text           string                 `json:"text,omitempty"`
	Created        int64                  `json:"created,omitempty"` // 目前用来给消息排序
	ErrorCode      int                    `json:"error_code,omitempty"`
	ErrorMsg       string                 `json:"error_msg,omitempty"`
	Usage          *ChatMessageTokenUsage `json:"usage,omitempty"`
	RegenMessageId string                 `json:"regen_message_id,omitempty"`
}

type ChatContents []ChatContent

// ChatContent 前端传参，表示本次 chat 的行为
type ChatContent struct {
	Type string       `json:"type,omitempty"`
	File *ContentFile `json:"file,omitempty"`

	ShimoFileGenerator *ContentShimoFileGenerator `json:"shimo_file_generator,omitempty"` // 生成文件的请求
	Text               *string                    `json:"text,omitempty"`                 // prompt
	TextProcessor      *ContentTextProcessor      `json:"text_processor,omitempty"`       // 生成 ppt 约定的入参
	TextContext        *string                    `json:"text_context,omitempty"`         // 和前端约定的 prompt 模板
}

type ContentShimoFileGenerator struct {
	InputContent          string                 `json:"input_content,omitempty"`
	InputContentMessageId string                 `json:"input_content_message_id,omitempty"`
	TargetFileType        string                 `json:"target_file_type,omitempty"`
	TargetFolder          string                 `json:"target_folder,omitempty"`
	FileName              string                 `json:"file_name,omitempty"`
	TemplateId            string                 `json:"template_id,omitempty"`
	FileInfo              map[string]interface{} `json:"file_info,omitempty"`
}

type ContentTextProcessor struct {
	Text   string                 `json:"text,omitempty"`
	Action string                 `json:"action,omitempty"`
	Guid   string                 `json:"guid,omitempty"`
	Ext    map[string]interface{} `json:"ext,omitempty"`
}

type ContentFile struct {
	Guid     string                 `json:"guid,omitempty"`
	FileKey  string                 `json:"file_key,omitempty"`
	Text     string                 `json:"text,omitempty"`
	FileInfo map[string]interface{} `json:"file_info,omitempty"`
}

type ChatMessageTokenUsage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`     // 问题tokens数
	CompletionTokens int `json:"completion_tokens,omitempty"` // 回答tokens数
	TotalTokens      int `json:"total_tokens,omitempty"`      // tokens总数 prompt_tokens + completion_tokens
}

type CreateAssetsResponse struct {
	ID        string      `json:"id"`
	Size      int         `json:"size"`
	UserQuery interface{} `json:"user_query"`
}
