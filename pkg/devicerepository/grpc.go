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
	response, err := dr.store.ListBrands(store.ListBrandsRequest{
		Limit:   request.Limit,
		Offset:  request.Offset,
		OrderBy: request.OrderBy,
		Paths:   request.FieldMask.Paths,
	})
	if err != nil {
		return nil, err
	}
	return &ttnpb.ListEndDeviceBrandsResponse{
		Brands: response.Brands,
		Count:  response.Count,
		Offset: response.Offset,
		Total:  response.Total,
	}, nil
}

// ListModels implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) ListModels(ctx context.Context, request *ttnpb.ListEndDeviceModelsRequest) (*ttnpb.ListEndDeviceModelsResponse, error) {
	response, err := dr.store.ListModels(store.ListModelsRequest{
		BrandID: request.BrandID,
		ModelID: request.ModelID,
		Limit:   request.Limit,
		Offset:  request.Offset,
		Paths:   request.FieldMask.Paths,
	})
	if err != nil {
		return nil, err
	}
	return &ttnpb.ListEndDeviceModelsResponse{
		Models: response.Models,
		Count:  response.Count,
		Offset: response.Offset,
		Total:  response.Total,
	}, nil
}

// GetTemplate implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) GetTemplate(ctx context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	return dr.store.GetTemplate(ids)
}

// GetUplinkDecoder implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) GetUplinkDecoder(ctx context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	s, err := dr.store.GetUplinkDecoder(ids)
	if err != nil {
		return nil, err
	}
	return &ttnpb.MessagePayloadFormatter{
		Codec: s,
	}, nil
}

// GetDownlinkDecoder implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) GetDownlinkDecoder(ctx context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	s, err := dr.store.GetDownlinkDecoder(ids)
	if err != nil {
		return nil, err
	}
	return &ttnpb.MessagePayloadFormatter{
		Codec: s,
	}, nil
}

// GetDownlinkEncoder implements the ttnpb.DeviceRepositoryServer interface.
func (dr *DeviceRepository) GetDownlinkEncoder(ctx context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	s, err := dr.store.GetDownlinkEncoder(ids)
	if err != nil {
		return nil, err
	}
	return &ttnpb.MessagePayloadFormatter{
		Codec: s,
	}, nil
}
