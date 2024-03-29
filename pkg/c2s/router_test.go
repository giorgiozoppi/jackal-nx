// Copyright 2022 The jackal Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package c2s

import (
	"context"
	"testing"

	"github.com/jackal-xmpp/stravaganza"
	"github.com/jackal-xmpp/stravaganza/jid"
	"github.com/ortuman/jackal/pkg/cluster/instance"
	"github.com/ortuman/jackal/pkg/hook"
	c2smodel "github.com/ortuman/jackal/pkg/model/c2s"
	"github.com/ortuman/jackal/pkg/router"
	"github.com/stretchr/testify/suite"
)

type routerSuite struct {
	suite.Suite
	localRouterMock   *localRouterMock
	clusterRouterMock *clusterRouterMock
	resMngMock        *resourceManagerMock
	repositoryMock    *repositoryMock
	router            *c2sRouter
}

func (s *routerSuite) SetupTest() {
	s.localRouterMock = &localRouterMock{}
	s.clusterRouterMock = &clusterRouterMock{}
	s.resMngMock = &resourceManagerMock{}
	s.repositoryMock = &repositoryMock{}
	s.router = &c2sRouter{
		local:   s.localRouterMock,
		cluster: s.clusterRouterMock,
		resMng:  s.resMngMock,
		rep:     s.repositoryMock,
		hk:      hook.NewHooks(),
	}
}

func (s *routerSuite) TestRouter_NotExistingAccount() {
	// given
	s.repositoryMock.UserExistsFunc = func(_ context.Context, _ string) (bool, error) {
		return false, nil
	}

	// when
	msg := testMessageStanza()
	_, err := s.router.Route(context.Background(), msg, router.CheckUserExistence)

	// then
	s.Require().Equal(router.ErrNotExistingAccount, err)
}

func (s *routerSuite) TestRouter_NotAuthenticated() {
	// given
	s.repositoryMock.UserExistsFunc = func(_ context.Context, _ string) (bool, error) {
		return false, nil
	}
	s.resMngMock.GetResourcesFunc = func(ctx context.Context, username string) ([]c2smodel.ResourceDesc, error) {
		return nil, nil
	}

	// when
	msg := testMessageStanza()
	_, err := s.router.Route(context.Background(), msg, router.RoutingOptions(0))

	// then
	s.Require().Equal(router.ErrUserNotAvailable, err)
}

func (s *routerSuite) TestRouter_ResourceNotFound() {
	// given
	jd, _ := jid.New("ortuman", "jackal.im", "yard", true)

	s.repositoryMock.UserExistsFunc = func(_ context.Context, _ string) (bool, error) {
		return false, nil
	}
	s.resMngMock.GetResourcesFunc = func(ctx context.Context, username string) ([]c2smodel.ResourceDesc, error) {
		return []c2smodel.ResourceDesc{
			c2smodel.NewResourceDesc(instance.ID(), jd, nil, c2smodel.NewInfoMap()),
		}, nil
	}

	// when
	msg := testMessageStanza()
	_, err := s.router.Route(context.Background(), msg, router.RoutingOptions(0))

	// then
	s.Require().Equal(router.ErrResourceNotFound, err)
}

func (s *routerSuite) TestRouter_LocalRoute() {
	// given
	jd, _ := jid.New("ortuman", "jackal.im", "balcony", true)

	s.repositoryMock.UserExistsFunc = func(_ context.Context, _ string) (bool, error) {
		return false, nil
	}
	s.resMngMock.GetResourcesFunc = func(_ context.Context, _ string) ([]c2smodel.ResourceDesc, error) {
		return []c2smodel.ResourceDesc{
			c2smodel.NewResourceDesc(instance.ID(), jd, nil, c2smodel.NewInfoMap()),
		}, nil
	}
	var routed bool
	s.localRouterMock.RouteFunc = func(stanza stravaganza.Stanza, username string, resource string) error {
		routed = true
		return nil
	}

	// when
	msg := testMessageStanza()
	_, err := s.router.Route(context.Background(), msg, router.RoutingOptions(0))

	// then
	s.Require().Nil(err)
	s.Require().True(routed)
}

func (s *routerSuite) TestRouter_ClusterRoute() {
	// given
	jd, _ := jid.New("ortuman", "jackal.im", "balcony", true)

	s.repositoryMock.UserExistsFunc = func(_ context.Context, _ string) (bool, error) {
		return false, nil
	}
	s.resMngMock.GetResourcesFunc = func(_ context.Context, _ string) ([]c2smodel.ResourceDesc, error) {
		return []c2smodel.ResourceDesc{
			c2smodel.NewResourceDesc("abcd1234", jd, nil, c2smodel.NewInfoMap()),
		}, nil
	}
	var routed bool
	s.clusterRouterMock.RouteFunc = func(ctx context.Context, stanza stravaganza.Stanza, username string, resource string, instanceID string) error {
		s.Require().Equal("abcd1234", instanceID)
		routed = true
		return nil
	}

	// when
	msg := testMessageStanza()
	_, err := s.router.Route(context.Background(), msg, router.RoutingOptions(0))

	// then
	s.Require().Nil(err)
	s.Require().True(routed)
}

func TestC2SRouterSuite(t *testing.T) {
	suite.Run(t, new(routerSuite))
}
