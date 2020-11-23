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
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// NoopStore is a no-op.
type NoopStore struct{}

// ListBrands lists available end device vendors.
func (*NoopStore) ListBrands(ListBrandsRequest) ([]*ttnpb.EndDeviceBrand, error) {
	return nil, nil
}

// ListModels lists available end device definitions.
func (*NoopStore) ListModels(ListModelsRequest) ([]*ttnpb.EndDeviceModel, error) {
	return nil, nil
}

// GetTemplate retrieves an end device template for an end device definition.
func (*NoopStore) GetTemplate(*ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	return nil, nil
}

// GetFormatters retrieves the message payload formatters for an end device template
func (*NoopStore) GetFormatters(*ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatters, error) {
	return nil, nil
}
