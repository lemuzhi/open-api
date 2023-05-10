package utils

// Copyright 2022 Beijing Douyin Information Service Co., Ltd.
// SPDX-License-Identifier: BSD-3-Clause
// Package utils sign.go 签名算法
//
// Package utils sign.go implement the signature algorithm

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

//var (
//	otherSettleParams = "other_settle_params"                      // 其他分账方参数 (Other settle params)
//	appId             = global.CONFIG.Douyin.AppId                     // 小程序appID (Applets appID)
//	thirdpartyId      = ""                                     // 代小程序进行该笔交易调用的第三方平台服务商 id (The id of the third-party platform service provider that calls the transaction on behalf of the Applets)
//	sign              = "" // 签名 (sign)
//)

// RequestSign 担保支付请求签名算法
// 参数："paramsMap" 所有的请求参数
// 返回：签名字符串
//
// RequestSign Guaranteed Payment Request Signature Algorithm
// Param: "paramsMap" all request parameters
// Return: signature string
func RequestSign(paramsMap map[string]interface{}, salt string) string {
	var paramsArr []string
	for k, v := range paramsMap {
		//fmt.Println("k = ", k)
		if k == "OtherSettleParams" || k == "app_id" || k == "thirdparty_id" || k == "sign" {
			continue
		}
		value := strings.TrimSpace(fmt.Sprintf("%v", v))
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") && len(value) > 1 {
			value = value[1 : len(value)-1]
		}
		value = strings.TrimSpace(value)
		if value == "" || value == "null" {
			continue
		}
		paramsArr = append(paramsArr, value)
	}

	// 支付密钥值，需要替换为自己的密钥(完成进件后，开发者可在字节开放平台-【某小程序】-【功能】-【支付】-【担保交易设置】中查看支付系统秘钥 SALT)
	paramsArr = append(paramsArr, salt)
	sort.Strings(paramsArr)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(paramsArr, "&"))))
}

// CallbackSign 担保支付回调签名算法
// 参数："strArr" 所有字段（验证时注意不包含 sign 签名本身，不包含空字段与 type 常量字段）内容与平台上配置的 token
// 返回：签名字符串
//
// CallbackSign Guaranteed payment callback signature algorithm
// Param: "strArr" The content of all fields (note that the sign signature itself is not included during verification, and does not include empty fields and type constant fields) content and the token configured on the platform
// Return: signature string
func CallbackSign(strArr []string) string {
	fmt.Println(strArr)
	sort.Strings(strArr)
	h := sha1.New()
	h.Write([]byte(strings.Join(strArr, "")))
	return fmt.Sprintf("%x", h.Sum(nil))
}
