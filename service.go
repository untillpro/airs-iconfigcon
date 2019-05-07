/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"context"
	"errors"
	"fmt"
)

type Service struct {
	Host string
	Port uint16
}

type contextKeyType string

const (
	consul            = contextKeyType("consul")
	DefaultConsulHost = "127.0.0.1"
	DefaultConsulPort = 8500
)

func getService(ctx context.Context) *Service {
	return ctx.Value(consul).(*Service)
}

// Start s.e.
func (s *Service) Start(ctx context.Context) (context.Context, error) {
	if s.Host == "" {
		return ctx, errors.New("host can't be empty")
	}
	if s.Port == 0 {
		return ctx, fmt.Errorf("passed port is invalid: %d", s.Port)
	}
	if ctx == nil {
		return ctx, errors.New("passed ctx can't be nil, pass context.TODO instead")
	}
	return context.WithValue(ctx, consul, s), nil
}

// Stop s.e.
func (s *Service) Stop(ctx context.Context) {

}
