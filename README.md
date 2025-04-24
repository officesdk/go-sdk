# How to Use

```
import (
	"github.com/gin-gonic/gin"
	"github.com/officesdk/go-sdk/officesdk"
	"time"
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
```
