// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package xpprovider exposes the Terraform Plugin Framework provider to
// Crossplane providers, which run it in-process (no-fork) instead of running
// the provider binary. The provider implementation lives in internal packages,
// which other modules cannot import.
package xpprovider

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"terraform-provider-infisical/internal/provider"
)

// New returns a new instance of the Infisical Terraform provider.
func New(version string) fwprovider.Provider {
	return provider.New(version)()
}
