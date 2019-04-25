/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package iconfigcon

import (
	"context"
	"fmt"
	"testing"

	gc "github.com/untillpro/gochips"
	"github.com/untillpro/godif"
	"github.com/untillpro/godif/iservices"
	"github.com/untillpro/godif/services"
)


func start(t *testing.T) (context.Context, error) {

	// Require iservices interface

	godif.Require(&iservices.InitAndStart)
	godif.Require(&iservices.StopAndFinit)

	// Provice iservices interface
	services.Declare()

	// Declare own service
	Declare(args)

	errs := godif.ResolveAll()
	gc.FatalIfError(t, errs, "Resolve problem")

	ctx := context.Background()
	ctx, err := iservices.InitAndStart(ctx)
	return ctx, err

}

func stop(ctx context.Context, t *testing.T) {
	iservices.StopAndFinit(ctx)
	godif.Reset()
}
