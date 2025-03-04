// Copyright 2023 MongoDB Inc
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

package store

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_backup_serverless.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ServerlessSnapshotsLister,ServerlessSnapshotsDescriber,ServerlessRestoreJobsLister,ServerlessRestoreJobsDescriber,ServerlessRestoreJobsCreator

type ServerlessSnapshotsLister interface {
	ServerlessSnapshots(string, string, *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error)
}

type ServerlessSnapshotsDescriber interface {
	ServerlessSnapshot(string, string, string) (*atlas.CloudProviderSnapshot, error)
}

type ServerlessRestoreJobsLister interface {
	ServerlessRestoreJobs(string, string, *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob, error)
}

type ServerlessRestoreJobsDescriber interface {
	ServerlessRestoreJob(string, string, string) (*atlasv2.ServerlessBackupRestoreJob, error)
}

type ServerlessRestoreJobsCreator interface {
	ServerlessCreateRestoreJobs(string, string, *atlasv2.ServerlessBackupRestoreJob) (*atlasv2.ServerlessBackupRestoreJob, error)
}

// ServerlessSnapshots encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessSnapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.CloudBackupsApi.ListServerlessBackups(s.ctx, projectID, clusterName).PageNum(int32(opts.PageNum)).ItemsPerPage(int32(opts.ItemsPerPage)).IncludeCount(opts.IncludeCount).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessSnapshot(projectID, instanceName, snapshotID string) (*atlas.CloudProviderSnapshot, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:      projectID,
		SnapshotID:   snapshotID,
		InstanceName: instanceName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetOneServerlessSnapshot(s.ctx, o)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessRestoreJobs(projectID, instanceName string, _ *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.CloudBackupsApi.ListServerlessBackupRestoreJobs(s.ctx, projectID, instanceName).Execute() // TODO:List options missing in call
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessRestoreJob(projectID, instanceName string, jobID string) (*atlasv2.ServerlessBackupRestoreJob, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.CloudBackupsApi.GetServerlessBackupRestoreJob(s.ctx, projectID, instanceName, jobID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessCreateRestoreJobs(projectID, clusterName string, request *atlasv2.ServerlessBackupRestoreJob) (*atlasv2.ServerlessBackupRestoreJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.CreateServerlessBackupRestoreJob(s.ctx, projectID, clusterName).ServerlessBackupRestoreJob(*request).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
