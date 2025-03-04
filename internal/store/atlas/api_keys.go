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

//go:generate mockgen -destination=../../mocks/atlas/mock_api_keys.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectAPIKeyLister,ProjectAPIKeyCreator,OrganizationAPIKeyLister,OrganizationAPIKeyDescriber,OrganizationAPIKeyUpdater,OrganizationAPIKeyCreator,OrganizationAPIKeyDeleter,ProjectAPIKeyDeleter,ProjectAPIKeyAssigner

type ProjectAPIKeyLister interface {
	ProjectAPIKeys(string, *atlas.ListOptions) (*atlasv2.PaginatedApiApiUser, error)
}

type ProjectAPIKeyCreator interface {
	CreateProjectAPIKey(string, *atlasv2.CreateApiKey) (*atlasv2.ApiUser, error)
}

type ProjectAPIKeyDeleter interface {
	DeleteProjectAPIKey(string, string) error
}

type ProjectAPIKeyAssigner interface {
	AssignProjectAPIKey(string, string, *atlasv2.CreateApiKey) error
}

type OrganizationAPIKeyLister interface {
	OrganizationAPIKeys(string, *atlas.ListOptions) (*atlasv2.PaginatedApiApiUser, error)
}

type OrganizationAPIKeyDescriber interface {
	OrganizationAPIKey(string, string) (*atlasv2.ApiUser, error)
}

type OrganizationAPIKeyUpdater interface {
	UpdateOrganizationAPIKey(string, string, *atlasv2.CreateApiKey) (*atlasv2.ApiUser, error)
}

type OrganizationAPIKeyCreator interface {
	CreateOrganizationAPIKey(string, *atlasv2.CreateApiKey) (*atlasv2.ApiUser, error)
}

type OrganizationAPIKeyDeleter interface {
	DeleteOrganizationAPIKey(string, string) error
}

// OrganizationAPIKeys encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKeys(orgID string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiApiUser, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.ListApiKeys(s.ctx, orgID).
		ItemsPerPage(int32(opts.ItemsPerPage)).PageNum(int32(opts.PageNum)).IncludeCount(opts.IncludeCount).Execute()
	return result, err
}

// OrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKey(orgID, apiKeyID string) (*atlasv2.ApiUser, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.GetApiKey(s.ctx, orgID, apiKeyID).Execute()
	return result, err
}

// UpdateOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateOrganizationAPIKey(orgID, apiKeyID string, input *atlasv2.CreateApiKey) (*atlasv2.ApiUser, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.UpdateApiKey(s.ctx, orgID, apiKeyID).CreateApiKey(*input).Execute()
	return result, err
}

// CreateOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOrganizationAPIKey(orgID string, input *atlasv2.CreateApiKey) (*atlasv2.ApiUser, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.CreateApiKey(s.ctx, orgID).CreateApiKey(*input).Execute()
	return result, err
}

// DeleteOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOrganizationAPIKey(orgID, id string) error {
	_, _, err := s.clientv2.ProgrammaticAPIKeysApi.DeleteApiKey(s.ctx, orgID, id).Execute()
	return err
}

// ProjectAPIKeys returns the API Keys for a specific project.
func (s *Store) ProjectAPIKeys(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiApiUser, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.ListProjectApiKeys(s.ctx, projectID).
		PageNum(int32(opts.PageNum)).ItemsPerPage(int32(opts.ItemsPerPage)).IncludeCount(opts.IncludeCount).Execute()
	return result, err
}

// CreateProjectAPIKey creates an API Keys for a project.
func (s *Store) CreateProjectAPIKey(projectID string, apiKeyInput *atlasv2.CreateApiKey) (*atlasv2.ApiUser, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.CreateProjectApiKey(s.ctx, projectID).CreateApiKey(*apiKeyInput).Execute()
	return result, err
}

// AssignProjectAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) AssignProjectAPIKey(projectID, apiKeyID string, input *atlasv2.CreateApiKey) error {
	_, _, err := s.clientv2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(s.ctx, projectID, apiKeyID).CreateApiKey(*input).Execute()
	return err
}

// DeleteProjectAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProjectAPIKey(projectID, id string) error {
	_, _, err := s.clientv2.ProgrammaticAPIKeysApi.RemoveProjectApiKey(s.ctx, projectID, id).Execute()
	return err
}
