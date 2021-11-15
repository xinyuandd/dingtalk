/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

type SnsUserInfo struct {
	//用户授权的临时授权码，只能使用一次。获取方法请参考：
	//
	//扫码登录第三方网站：开发流程
	//
	//免登第三方网站：开发流程
	//
	//使用钉钉账号登录第三方网站：开发流程
	AutoCode string `json:"tmp_auth_code,omitempty" validate:"required"`
}

func NewSnsUserInfo(code string) *SnsUserInfo {
	return &SnsUserInfo{code}
}
