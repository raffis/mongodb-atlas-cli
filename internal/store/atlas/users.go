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

package atlas

import (
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_users.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas UserCreator,UserDescriber,UserLister,TeamUserLister

type UserCreator interface {
	CreateUser(user *atlasv2.AppUser) (*atlasv2.AppUser, error)
}

type UserLister interface {
	OrganizationUsers(string, *atlas.ListOptions) (*atlasv2.PaginatedAppUser, error)
}

type TeamUserLister interface {
	TeamUsers(string, string) (*atlasv2.PaginatedApiAppUser, error)
}

type UserDescriber interface {
	UserByID(string) (*atlasv2.AppUser, error)
	UserByName(string) (*atlasv2.AppUser, error)
}

// CreateUser encapsulates the logic to manage different cloud providers.
func (s *Store) CreateUser(user *atlasv2.AppUser) (*atlasv2.AppUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.CreateUser(s.ctx).AppUser(*user).Execute()
	return result, err
}

// UserByID encapsulates the logic to manage different cloud providers.
func (s *Store) UserByID(userID string) (*atlasv2.AppUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.GetUser(s.ctx, userID).Execute()
	return result, err
}

// UserByName encapsulates the logic to manage different cloud providers.
func (s *Store) UserByName(username string) (*atlasv2.AppUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.GetUserByUsername(s.ctx, username).Execute()
	return result, err
}

// OrganizationUsers encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationUsers(organizationID string, opts *atlas.ListOptions) (*atlasv2.PaginatedAppUser, error) {
	result, _, err := s.clientv2.OrganizationsApi.ListOrganizationUsers(s.ctx, organizationID).
		IncludeCount(opts.IncludeCount).ItemsPerPage(int32(opts.ItemsPerPage)).PageNum(int32(opts.PageNum)).Execute()
	return result, err
}

// TeamUsers encapsulates the logic to manage different cloud providers.
func (s *Store) TeamUsers(orgID, teamID string) (*atlasv2.PaginatedApiAppUser, error) {
	result, _, err := s.clientv2.TeamsApi.ListTeamUsers(s.ctx, orgID, teamID).Execute()
	return result, err
}
