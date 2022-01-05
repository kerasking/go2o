/**
 * This file is auto generated by tto v0.4.5 !
 * If you want to modify this code, please read the guide
 * to modify code template.
 *
 * Get started: https://github.com/ixre/tto
 *
 * Copyright (C) 2022 56X.NET, All rights reserved.
 *
 * name : sys_search_word_dao.go
 * author : jarrysix
 * date : 2022/01/05 10:13:57
 * description :
 * history :
 */
package dao

import (
	"github.com/ixre/go2o/core/dao/model"
)

type ISysSearchWordDao interface {
	// GetSearchWord Get 热搜词
	GetSearchWord(primary interface{}) *model.SearchWord
	// GetSearchWordBy GetBy 热搜词
	GetSearchWordBy(where string, v ...interface{}) *model.SearchWord
	// CountSearchWord Count 热搜词 by condition
	CountSearchWord(where string, v ...interface{}) (int, error)
	// SelectSearchWord Select 热搜词
	SelectSearchWord(where string, v ...interface{}) []*model.SearchWord
	// SaveSearchWord Save 热搜词
	SaveSearchWord(v *model.SearchWord) (int, error)
	// DeleteSearchWord Delete 热搜词
	DeleteSearchWord(primary interface{}) error
	// BatchDeleteSearchWord Batch Delete 热搜词
	BatchDeleteSearchWord(where string, v ...interface{}) (int64, error)
}
