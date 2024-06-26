// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/internal/util/apiclient/server"
	"github.com/daytonaio/daytona/pkg/views"
	"github.com/daytonaio/daytona/pkg/views/provider"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var providerUninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Uninstall provider",
	Args:    cobra.NoArgs,
	Aliases: []string{"u"},
	Run: func(cmd *cobra.Command, args []string) {
		providerList, err := server.GetProviderList()
		if err != nil {
			log.Fatal(err)
		}

		providerToUninstall, err := provider.GetProviderFromPrompt(providerList, "Choose a provider to uninstall", false)
		if err != nil {
			log.Fatal(err)
		}

		if providerToUninstall == nil {
			return
		}

		apiClient, err := server.GetApiClient(nil)
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()

		res, err := apiClient.ProviderAPI.UninstallProvider(ctx, *providerToUninstall.Name).Execute()

		if err != nil {
			log.Fatal(apiclient.HandleErrorResponse(res, err))
		}

		views.RenderInfoMessageBold(fmt.Sprintf("Provider %s has been successfully uninstalled", *providerToUninstall.Name))
	},
}
