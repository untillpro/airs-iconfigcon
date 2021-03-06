/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"github.com/untillpro/airs-iconfig"
	"testing"
)

func Test_Impl(t *testing.T) {
	ctx := start(t)
	defer stop(ctx, t)

	iconfig.TestImpl(ctx, t)
}
