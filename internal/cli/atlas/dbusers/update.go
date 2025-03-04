// Copyright 2021 MongoDB Inc
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

package dbusers

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const updateTemplate = "Successfully updated database user '{{.Username}}'.\n"

type UpdateOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	username string
	password string
	roles    []string
	scopes   []string
	store    store.DatabaseUserUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	current := new(atlas.DatabaseUser)
	opts.update(current)

	r, err := opts.store.UpdateDatabaseUser(current)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) update(out *atlas.DatabaseUser) {
	out.GroupID = opts.ConfigProjectID()
	out.Username = opts.username
	if opts.password != "" {
		out.Password = opts.password
	}

	out.Scopes = convert.BuildAtlasScopes(opts.scopes)
	out.Roles = convert.BuildAtlasRoles(opts.roles)
}

// mongocli atlas dbuser(s) update <username> [--password password] [--role roleName@dbName] [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <username>",
		Short: "Modify the details for a database user for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  # Update roles for a database user named myUser for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s dbuser update myUser --role readWriteAnyDatabase --projectId 5e2211c17a3e5a48f5497de3

  # Update scopes for a database user named myUser for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s dbuser update myUser --scope resourceName:resourceType --projectId 5e2211c17a3e5a48f5497de3`,
			cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"usernameDesc": "Username to update in the MongoDB database.",
			"output":       updateTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.username = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.username, flag.Username, flag.UsernameShort, "", usage.DBUsername)
	cmd.Flags().StringVarP(&opts.password, flag.Password, flag.PasswordShort, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.Roles)
	cmd.Flags().StringSliceVar(&opts.scopes, flag.Scope, []string{}, usage.Scopes)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
