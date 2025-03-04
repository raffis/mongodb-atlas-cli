// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package peering

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPeeringConnectionLister(ctrl)

	t.Run("no provider", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
		}

		var expected []mongodbatlas.Peer
		mockStore.
			EXPECT().
			PeeringConnections(listOpts.ProjectID, listOpts.newContainerListOptions()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
