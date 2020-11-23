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

package devicerepository

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	errNotImplemented = errors.DefineUnimplemented("not_implemented", "not implemented")
)

// ListBrands implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) ListBrands(ctx context.Context, request *ttnpb.ListEndDeviceBrandsRequest) (*ttnpb.ListEndDeviceBrandsResponse, error) {
	brands, err := dr.store.ListBrands(store.ListBrandsRequest{
		Limit:   request.Limit,
		Offset:  request.Offset,
		OrderBy: request.OrderBy,
		Paths:   request.FieldMask.Paths,
	})
	if err != nil {
		return nil, err
	}
	return &ttnpb.ListEndDeviceBrandsResponse{Brands: brands}, nil
}

// ListModels implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) ListModels(ctx context.Context, request *ttnpb.ListEndDeviceModelsRequest) (*ttnpb.ListEndDeviceModelsResponse, error) {
	defs, err := dr.store.ListModels(store.ListModelsRequest{
		BrandID: request.BrandID,
		ModelID: request.ModelID,
		Limit:   request.Limit,
		Offset:  request.Offset,
		Paths:   request.FieldMask.Paths,
	})
	if err != nil {
		return nil, err
	}
	return &ttnpb.ListEndDeviceModelsResponse{Models: defs}, nil
}

// GetTemplate implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) GetTemplate(ctx context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	return dr.store.GetTemplate(ids)
}

// GetMessagePayloadFormatters implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) GetMessagePayloadFormatters(ctx context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatters, error) {
	return dr.store.GetFormatters(ids)
}
