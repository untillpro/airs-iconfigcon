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
	"github.com/stretchr/testify/require"
	iconfig "github.com/untillpro/airs-iconfig"
	"github.com/untillpro/godif/services"
	"testing"

	"github.com/untillpro/godif"
)

var host = "127.0.0.1"
var port uint16 = 8500

func Test_StartStop(t *testing.T) {
	ctx := start(t)
	defer stop(ctx, t)

	srv := getService(ctx)

	require.Equal(t, host, srv.Host)
	require.Equal(t, port, srv.Port)
}

func start(t *testing.T) context.Context {
	godif.Require(&iconfig.GetConfig)
	godif.Require(&iconfig.PutConfig)

	// Declare own service
	Declare(Service{host, port})

	ctx, err := services.ResolveAndStart()
	require.Nil(t, err)
	return ctx
}

func stop(ctx context.Context, t *testing.T) {
	services.StopServices(ctx)
	godif.Reset()
}
