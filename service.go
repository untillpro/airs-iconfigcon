/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"context"
	gc "github.com/untillpro/gochips"
)

type service struct {
	params *Params
}

type contextKeyType string

const contextKey = contextKeyType("contextKey")

// Init s.e.
func (s *service) Init(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (s *service) Start(ctx context.Context) error {
	return nil
}

// Stop s.e.
func (s *service) Stop(ctx context.Context) {

}

// Finit s.e.
func (s *service) Finit(ctx context.Context) {
}
