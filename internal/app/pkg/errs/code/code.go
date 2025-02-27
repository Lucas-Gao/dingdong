//go:generate stringer -type=ErrorCode --linecomment

package code

type ErrorCode int

const (
	Unexpected            ErrorCode = 1000 + iota // 意外错误
	OutOfRange                                    // 数组索引越界
	SignFailed                                    // 签名失败
	AssertFailed                                  // 断言失败
	ParseFailed                                   // 解析失败
	RequestFailed                                 // 请求失败
	InvalidResponse                               // 无效的响应
	NoValidAddress                                // 当前没有可用的收货地址
	NoValidProduct                                // 当前没有可购商品
	NoReserveTime                                 // 当前没有可用的配送时段
	NoReserveTimeAndRetry                         // 当前没有可用的配送时段, 请稍后再试
)

func (i ErrorCode) Int() int {
	return int(i)
}

func (i ErrorCode) Uint() uint {
	return uint(i)
}
