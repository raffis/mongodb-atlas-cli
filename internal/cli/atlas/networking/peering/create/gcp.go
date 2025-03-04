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

package create

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

type GCPOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	atlasCIDRBlock string
	gcpProjectID   string
	network        string
	regions        []string
	store          store.GCPPeeringConnectionCreator
}

func (opts *GCPOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GCPOpts) Run() error {
	container, err := opts.containerExists()
	if err != nil {
		return err
	}

	if container == nil {
		var err2 error
		r, err2 := opts.store.CreateContainer(opts.ConfigProjectID(), opts.newContainer())
		container = r.(*atlasv2.GCPCloudProviderContainer)
		if err2 != nil {
			return err2
		}
	}
	r, err := opts.store.CreatePeeringConnection(opts.ConfigProjectID(), opts.newPeer(*container.Id))
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *GCPOpts) containerExists() (*atlasv2.GCPCloudProviderContainer, error) {
	r, err := opts.store.GCPContainers(opts.ConfigProjectID())
	if err != nil {
		return nil, err
	}
	if len(r) > 0 {
		return r[0], nil
	}
	return nil, nil
}

func (opts *GCPOpts) newGCPContainer() *atlasv2.GCPCloudProviderContainer {
	c := &atlasv2.GCPCloudProviderContainer{
		AtlasCidrBlock: opts.atlasCIDRBlock,
		Regions:        opts.regions,
		ProviderName:   pointer.Get("GCP"),
	}
	return c
}

func (opts *GCPOpts) newContainer() *atlasv2.CloudProviderContainer {
	w := atlasv2.GCPCloudProviderContainerAsCloudProviderContainer(opts.newGCPContainer())
	return &w
}

func (opts *GCPOpts) newPeer(containerID string) *atlas.Peer {
	a := &atlas.Peer{
		ContainerID:  containerID,
		GCPProjectID: opts.gcpProjectID,
		NetworkName:  opts.network,
		ProviderName: "GCP",
	}
	return a
}

// mongocli atlas networking peering create gcp [--atlasCidrBlock atlasCidrBlock][--gcpProjectId gcpProjectId][--network networkName]
// [--regions region][--projectId projectId]
// --atlasCidrBlock atlasCidrBlock: CIDR block that Atlas uses for the Network Peering containers in your project.
// --gcpProjectId gcpProjectId: GCP project ID of the owner of the network peer.
// --network networkName: Name of the network peer to which Atlas connects.
// --regions region: "List of Atlas regions where the container resides."
// --projectId projectId: ID of the project
// Create a network peering with GCP, this command will internally check if a container already exists for the provider and if it does then we’ll use that,
// if it does not exists we’ll try to create one and use it, there can only be one container per GCP provider.
func GCPBuilder() *cobra.Command {
	opts := &GCPOpts{}
	cmd := &cobra.Command{
		Use:   "gcp",
		Short: "Create a network peering connection between the Atlas VPC and your Google Cloud VPC.",
		Long: `The network peering create command checks if a VPC exists in the region you specify for your Atlas project. If one exists, this command creates the peering connection between that VPC and your VPC. If an Atlas VPC doesn't exist, this command creates one and creates a connection between it and your VPC.
		
To learn more about network peering connections, see https://www.mongodb.com/docs/atlas/security-vpc-peering/.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: fmt.Sprintf(`  # Create a network peering connection between the Atlas VPC in CIDR block 192.168.0.0/24 and your GCP VPC with the GCP project ID grandiose-branch-256701 in the network named cli-test:
  %s networking peering create gcp --atlasCidrBlock 192.168.0.0/24 --gcpProjectId grandiose-branch-256701 --network cli-test --output json`, cli.ExampleAtlasEntryPoint()),
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.gcpProjectID, flag.GCPProjectID, "", usage.GCPProjectID)
	cmd.Flags().StringVar(&opts.network, flag.Network, "", usage.Network)
	cmd.Flags().StringVar(&opts.atlasCIDRBlock, flag.AtlasCIDRBlock, "", usage.AtlasCIDRBlock)
	cmd.Flags().StringSliceVar(&opts.regions, flag.Region, []string{}, usage.ContainerRegions)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.AtlasCIDRBlock)

	return cmd
}
