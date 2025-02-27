// Code generated by "stringer -type=ErrorCode --linecomment"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unexpected-1000]
	_ = x[OutOfRange-1001]
	_ = x[SignFailed-1002]
	_ = x[AssertFailed-1003]
	_ = x[ParseFailed-1004]
	_ = x[RequestFailed-1005]
	_ = x[InvalidResponse-1006]
	_ = x[NoValidAddress-1007]
	_ = x[NoValidProduct-1008]
	_ = x[NoReserveTime-1009]
	_ = x[NoReserveTimeAndRetry-1010]
}

const _ErrorCode_name = "意外错误数组索引越界签名失败断言失败解析失败请求失败无效的响应当前没有可用的收货地址当前没有可购商品当前没有可用的配送时段当前没有可用的配送时段, 请稍后再试"

var _ErrorCode_index = [...]uint8{0, 12, 30, 42, 54, 66, 78, 93, 126, 150, 183, 233}

func (i ErrorCode) String() string {
	i -= 1000
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i+1000), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
