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

package customdbroles

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

const createTemplate = "Custom database role '{{.RoleName}}' successfully created.\n"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	action         []string
	roleName       string
	inheritedRoles []string
	store          store.DatabaseRoleCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	role := opts.newCustomDBRole()

	r, err := opts.store.CreateDatabaseRole(opts.ConfigProjectID(), role)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCustomDBRole() *atlasv2.CustomDBRole {
	return &atlasv2.CustomDBRole{
		RoleName:       opts.roleName,
		Actions:        joinActions(convert.BuildAtlasActions(opts.action)),
		InheritedRoles: convert.BuildAtlasInheritedRoles(opts.inheritedRoles),
	}
}

func (opts *CreateOpts) validate() error {
	if len(opts.action) == 0 && opts.inheritedRoles == nil {
		return errors.New("you must provide either actions or inherited roles")
	}

	return nil
}

// mongocli atlas dbrole(s) create <roleName> --privilege action[@dbName.collection] --inheritedRole role@db.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:     "create <roleName>",
		Short:   "Create a custom database role for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  %s customDbRoles create customRole --privilege FIND@database,UPDATE@database`, cli.ExampleAtlasEntryPoint()),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"roleNameDesc": "Name of the custom role to create.",
			"output":       createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.roleName = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.inheritedRoles, flag.InheritedRole, []string{}, usage.InheritedRoles)
	cmd.Flags().StringSliceVar(&opts.action, flag.Privilege, []string{}, usage.PrivilegeAction)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
