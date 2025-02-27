/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"fmt"

	"github.com/west2-online/fzuhelper-server/kitex_gen/course"
	login_model "github.com/west2-online/fzuhelper-server/kitex_gen/model"
	"github.com/west2-online/fzuhelper-server/pkg/base"
	"github.com/west2-online/fzuhelper-server/pkg/base/context"
	"github.com/west2-online/fzuhelper-server/pkg/logger"
	"github.com/west2-online/fzuhelper-server/pkg/utils"
	"github.com/west2-online/jwch"
)

// GetTermsList 会返回当前用户含有课表的学期信息
func (s *CourseService) GetTermsList(req *course.TermListRequest) ([]string, error) {
	var loginData *login_model.LoginData
	var err error

	loginData, err = context.GetLoginData(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("service.GetTermList: Get login data fail: %w", err)
	}

	if s.cache.IsKeyExist(s.ctx, loginData.GetId()) {
		terms, err := s.cache.Course.GetTermsCache(s.ctx, loginData.GetId())
		if err = base.HandleJwchError(err); err != nil {
			return nil, fmt.Errorf("service.GetTermList: Get terms cache fail: %w", err)
		}
		return terms, nil
	}

	stu := jwch.NewStudent().WithLoginData(loginData.GetId(), utils.ParseCookies(loginData.GetCookies()))
	terms, err := stu.GetTerms()
	if err = base.HandleJwchError(err); err != nil {
		return nil, fmt.Errorf("service.GetTermList: Get terms fail: %w", err)
	}
	go func() {
		err = s.cache.Course.SetTermsCache(s.ctx, loginData.GetId(), terms.Terms)
		if err = base.HandleJwchError(err); err != nil {
			logger.Errorf("service.GetTermList: set cache fail: %v", err)
		}
	}()

	return terms.Terms, nil
}
