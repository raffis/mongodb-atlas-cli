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

package clusters

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	name  string
	store store.AtlasClusterDescriber
}

var watchTemplate = "\nCluster available.\n"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func isRetryable(err error) bool {
	var atlasErr *atlas.ErrorResponse
	return errors.As(err, &atlasErr) && atlasErr.ErrorCode == "CLUSTER_NOT_FOUND"
}

func (opts *WatchOpts) watcher() (bool, error) {
	result, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return false, err
	}
	if result.StateName == "UPDATING" {
		opts.IsRetryableErr = isRetryable
	}
	return result.StateName == "IDLE", nil
}

func (opts *WatchOpts) Run() error {
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// mongocli atlas cluster(s) watch <clusterName> [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <clusterName>",
		Short: "Watch the specified cluster in your project until it becomes available.",
		Long: `This command checks the cluster's status periodically until it reaches an IDLE state. 
Once the cluster reaches the expected state, the command prints "Cluster available."
If you run the command in the terminal, it blocks the terminal session until the resource state changes to IDLE.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: fmt.Sprintf(`  # Watch for the cluster named myCluster to become available for the project with ID 5e2211c17a3e5a48f5497de3:
  %s clusters watch myCluster --projectId 5e2211c17a3e5a48f5497de3`, cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to watch.",
			"output":          watchTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
