package common

import "strings"

// 错误码
const (
	OK                              = 0
	TokenInvalid                    = 101
	TokenDeny                       = 102
	TokenExpired                    = 103
	TokenExpiredSoon                = 104 // Deprecated: use TokenExpired instead
	TokenBanned                     = 105
	SignInvalid                     = 106
	ACLConnectServerFailed          = 201
	ACLCallFailed                   = 202
	ACLObjectLevelLimit             = 203
	ACLFieldLevelLimit              = 204
	ACLRecordLevelLimit             = 205
	ACLActionTypeUnsupported        = 206
	ACLRecordAccessLevelUnsupported = 207
	ACLFunctionPermissionLimit      = 208
	AccountTokenInvalid             = 301
	UmbrellaMysqlError              = 200000
	UmbrellaVideoThemeNotFound      = 200001
	UmbrellaVideoInternalError      = 200002
	UmbrellaVideoNotFound           = 200003
	UmbrellaVideoRequestIOError     = 200004
	UmbrellaVideoJsonUnmarshalErr   = 200005
)

// ErrorMap 错误码与错误信息map
var ErrorMap = map[int]map[string]string{
	0: map[string]string{
		"EN-US": "OK",
		"ZH-CN": "操作成功",
	},
	101: map[string]string{
		"EN-US": "Token Invalid",
		"ZH-CN": "",
	},
	102: map[string]string{
		"EN-US": "Token Deny",
		"ZH-CN": "",
	},
	103: map[string]string{
		"EN-US": "Token Expired",
		"ZH-CN": "",
	},
	104: map[string]string{
		"EN-US": "Token Expired Soon",
		"ZH-CN": "",
	},
	105: map[string]string{
		"EN-US": "Token Banned",
		"ZH-CN": "",
	},
	106: map[string]string{
		"EN-US": "Sign Invalid",
		"ZH-CN": "",
	},
	201: map[string]string{
		"EN-US": "ACL Connect Server Failed",
		"ZH-CN": "",
	},
	202: map[string]string{
		"EN-US": "ACL Call Failed",
		"ZH-CN": "",
	},
	203: map[string]string{
		"EN-US": "ACL Object Level Limit",
		"ZH-CN": "",
	},
	204: map[string]string{
		"EN-US": "ACL Field Level Limit",
		"ZH-CN": "",
	},
	205: map[string]string{
		"EN-US": "ACL Record Level Limit",
		"ZH-CN": "",
	},
	206: map[string]string{
		"EN-US": "ACL Action Type Unsupported",
		"ZH-CN": "",
	},
	207: map[string]string{
		"EN-US": "ACL Record Access Level Unsupported",
		"ZH-CN": "",
	},
	208: map[string]string{
		"EN-US": "ACL Function Permission Limit",
		"ZH-CN": "",
	},
	301: map[string]string{
		"EN-US": "Account Token Invalid",
		"ZH-CN": "",
	},
	200000: map[string]string{
		"EN-US": "Umbrella Mysql Error",
		"ZH-CN": "",
	},
	200001: map[string]string{
		"EN-US": "Umbrella Video Theme Not Found",
		"ZH-CN": "",
	},
	200002: map[string]string{
		"EN-US": "Umbrella Video Internal Error",
		"ZH-CN": "",
	},
	200003: map[string]string{
		"EN-US": "Umbrella Video Not Found",
		"ZH-CN": "",
	},
	200004: map[string]string{
		"EN-US": "Umbrella Video Request IO Error",
		"ZH-CN": "",
	},
	200005: map[string]string{
		"EN-US": "Umbrella Video Json Unmarshal Error",
		"ZH-CN": "",
	},
}

// GetMsg 错误码转各个语言的错误信息
func GetMsg(code int, languages []string) string {
	msgMap, ok := ErrorMap[code]
	if !ok {
		return "Unknown error"
	}
	for _, lang := range languages {
		if msg, ok := msgMap[strings.ToUpper(lang)]; ok {
			if msg != "" {
				return msg
			}
		}
	}
	return "Unknown error"
}
