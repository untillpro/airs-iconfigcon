/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"github.com/untillpro/airs-iconfig"
	"github.com/untillpro/godif"
	"github.com/untillpro/godif/iservices"
)

// Declare s.e.
func Declare(service Service) {
	godif.ProvideSliceElement(&iservices.Services, &service)
	godif.Provide(&iconfig.GetConfig, getConfig)
	godif.Provide(&iconfig.PutConfig, putConfig)
}
