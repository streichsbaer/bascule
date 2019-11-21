// Code generated by "stringer -type=ErrorResponseReason"; DO NOT EDIT.

package basculehttp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MissingHeader-0]
	_ = x[InvalidHeader-1]
	_ = x[KeyNotSupported-2]
	_ = x[ParseFailed-3]
	_ = x[GetURLFailed-4]
	_ = x[MissingAuthentication-5]
	_ = x[ChecksNotFound-6]
	_ = x[ChecksFailed-7]
}

const _ErrorResponseReason_name = "MissingHeaderInvalidHeaderKeyNotSupportedParseFailedGetURLFailedMissingAuthenticationChecksNotFoundChecksFailed"

var _ErrorResponseReason_index = [...]uint8{0, 13, 26, 41, 52, 64, 85, 99, 111}

func (i ErrorResponseReason) String() string {
	if i < 0 || i >= ErrorResponseReason(len(_ErrorResponseReason_index)-1) {
		return "ErrorResponseReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorResponseReason_name[_ErrorResponseReason_index[i]:_ErrorResponseReason_index[i+1]]
}
