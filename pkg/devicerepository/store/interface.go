// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// ListBrandsRequest is a request to list available end device vendors, with pagination and sorting.
type ListBrandsRequest struct {
	Limit,
	Offset *pbtypes.UInt32Value
	OrderBy string
	Paths   []string
}

// ListModelsRequest is a request to list available end device model definitions.
type ListModelsRequest struct {
	BrandID,
	ModelID string
	Limit,
	Offset *pbtypes.UInt32Value
	Paths []string
}

// DefinitionIdentifiers is a request to retrieve an end device template for an end device definition.
type DefinitionIdentifiers struct {
	BrandID,
	ModelID,
	FirmwareVersion,
	BandID string
}

// Store contains end device definitions.
type Store interface {
	// ListBrands lists available end device vendors.
	ListBrands(ListBrandsRequest) ([]*ttnpb.EndDeviceBrand, error)
	// ListModels lists available end device definitions.
	ListModels(ListModelsRequest) ([]*ttnpb.EndDeviceModel, error)
	// GetTemplate retrieves an end device template for an end device definition.
	GetTemplate(*ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error)
	// GetFormatters retrieves the message payload formatters for an end device template.
	GetFormatters(*ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatters, error)
}
