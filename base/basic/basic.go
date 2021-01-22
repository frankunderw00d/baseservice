package basic

type (
	// 组合字符串
	ComposeString string
)

const ()

var ()

// 拼接字符串
func (cs ComposeString) Compose(v string) string {
	return string(cs) + v
}

// 字符串化
func (cs ComposeString) String() string {
	return string(cs)
}
