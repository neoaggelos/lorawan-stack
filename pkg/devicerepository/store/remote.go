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
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"gopkg.in/yaml.v2"
)

// remoteStore implements the Store interface using a fetcher.
type remoteStore struct {
	fetcher fetch.Interface
}

// NewRemoteStore initializes a new Store using a fetcher.
func NewRemoteStore(fetcher fetch.Interface) (Store, error) {
	return &remoteStore{fetcher}, nil
}

// paginate returns page start and end indices, and false if the page is invalid.
func paginate(size int, limit, offset *pbtypes.UInt32Value) (uint32, uint32, bool) {
	start, end := uint32(0), uint32(size)
	if offset != nil && offset.Value > 0 {
		start = offset.Value
	}
	if start >= end {
		return 0, 0, false
	}
	if limit != nil && limit.Value > 0 && start+limit.Value < end {
		end = start + limit.Value
	}
	return start, end, true
}

// ListBrands lists available end device vendors from the vendor/index.yaml file.
func (s *remoteStore) ListBrands(req ListBrandsRequest) ([]*ttnpb.EndDeviceBrand, error) {
	b, err := s.fetcher.File("vendor", "index.yaml")
	if err != nil {
		return nil, err
	}
	rawVendors := VendorsIndex{}
	if err := yaml.Unmarshal(b, &rawVendors); err != nil {
		return nil, err
	}

	start, end, ok := paginate(len(rawVendors.Vendors), req.Limit, req.Offset)
	if !ok {
		return []*ttnpb.EndDeviceBrand{}, nil
	}

	brands := make([]*ttnpb.EndDeviceBrand, 0, end-start)
	for idx := start; idx < end; idx++ {
		brands = append(brands, &ttnpb.EndDeviceBrand{
			BrandID:   rawVendors.Vendors[idx].ID,
			BrandName: rawVendors.Vendors[idx].Name,
		})
	}

	return brands, nil
}

var (
	errUnknownBrand = errors.DefineNotFound("unknown_brand", "unknown brand `{brand_id}`")
)

// ListDefinitions lists available end device definitions.
func (s *remoteStore) ListDefinitions(req ListDefinitionsRequest) ([]*ttnpb.EndDeviceDefinition, error) {
	b, err := s.fetcher.File("vendor", req.BrandID, "index.yaml")
	if err != nil {
		return nil, errUnknownBrand.WithAttributes("brand_id", req.BrandID)
	}
	index := VendorEndDevicesIndex{}
	if err := yaml.Unmarshal(b, &index); err != nil {
		return nil, err
	}
	start, end, ok := paginate(len(index.EndDevices), req.Limit, req.Offset)
	if !ok {
		return []*ttnpb.EndDeviceDefinition{}, nil
	}

	defs := make([]*ttnpb.EndDeviceDefinition, 0, end-start)
	for idx := start; idx < end; idx++ {
		definitionID := index.EndDevices[idx]
		if req.ModelID != "" && definitionID != req.ModelID {
			continue
		}
		b, err := s.fetcher.File("vendor", req.BrandID, definitionID+".yaml")
		if err != nil {
			return nil, err
		}
		definition := EndDeviceDefinition{}
		if err := yaml.Unmarshal(b, &definition); err != nil {
			return nil, err
		}
		pb, err := definition.ToPB(definitionID, req.Paths...)
		if err != nil {
			return nil, err
		}
		defs = append(defs, pb)
	}
	return defs, nil
}

// GetTemplate retrieves an end device template for an end device definition.
func (s *remoteStore) GetTemplate(DefinitionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	return nil, nil
}

// GetFormatters retrieves the message payload formatters for an end device template
func (s *remoteStore) GetFormatters(DefinitionIdentifiers) (*ttnpb.MessagePayloadFormatters, error) {
	return nil, nil
}
