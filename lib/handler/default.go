package handler

import "smooth-bear.live/lib/database"

type _default struct {
	accessManage *database.AccessorManage
}

func NewDefault(manage *database.AccessorManage) _default {
	return _default{accessManage: manage}
}
