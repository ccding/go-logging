// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
//
// This file defines GetGoId function, which is used to get the id of the
// current goroutine. More details about this function are availeble in the
// runtime.c file of golang source code.
#include <runtime.h>

void GetGoId(int32 ret) {
	ret = g->goid;
	USED(&ret);
}
