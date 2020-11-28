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

package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	errCorruptedIndex = errors.DefineCorruption("corrupted_index", "corrupted index file")
)

// ListBrands lists available end device vendors.
func (bl *bleveStore) ListBrands(req store.ListBrandsRequest) (*store.ListBrandsResponse, error) {
	queries := []query.Query{
		bleve.NewMatchAllQuery(),
	}
	if q := req.Search; q != "" {
		queries = append(queries, bleve.NewQueryStringQuery(q))
	}
	if q := req.BrandID; q != "" {
		queries = append(queries, bleve.NewPhraseQuery([]string{q}, "BrandID"))
	}

	searchRequest := bleve.NewSearchRequest(bleve.NewConjunctionQuery(queries...))
	if limit := req.Limit; limit != nil && limit.Value > 0 {
		searchRequest.Size = int(limit.Value)
	}
	if offset := req.Offset; offset != nil && offset.Value > 0 {
		searchRequest.From = int(offset.Value)
	}
	searchRequest.Fields = []string{"BrandPB"}
	switch req.OrderBy {
	case "brand_id":
		searchRequest.SortBy([]string{"BrandID"})
	case "-brand_id":
		searchRequest.SortBy([]string{"-BrandID"})
	case "name":
		searchRequest.SortBy([]string{"BrandName"})
	case "-name":
		searchRequest.SortBy([]string{"-BrandName"})
	}

	bl.brandsIndexMu.RLock()
	result, err := bl.brandsIndex.Search(searchRequest)
	bl.brandsIndexMu.RUnlock()
	if err != nil {
		return nil, err
	}

	brands := make([]*ttnpb.EndDeviceBrand, 0, len(result.Hits))
	for _, hit := range result.Hits {
		s, ok := hit.Fields["BrandPB"].(string)
		if !ok {
			return nil, errCorruptedIndex.New()
		}
		brand := &ttnpb.EndDeviceBrand{}
		if err := jsonpb.TTN().Unmarshal([]byte(s), brand); err != nil {
			return nil, err
		}
		pb := &ttnpb.EndDeviceBrand{}
		if err := pb.SetFields(brand, req.Paths...); err != nil {
			return nil, err
		}
		brands = append(brands, pb)
	}
	return &store.ListBrandsResponse{
		Count:  uint32(len(result.Hits)),
		Total:  uint32(result.Total),
		Offset: uint32(searchRequest.From),
		Brands: brands,
	}, nil
}

// ListModels lists available end device definitions.
func (bl *bleveStore) ListModels(req store.ListModelsRequest) (*store.ListModelsResponse, error) {
	queries := []query.Query{
		bleve.NewMatchAllQuery(),
	}
	if q := req.Search; q != "" {
		queries = append(queries, bleve.NewQueryStringQuery(q))
	}
	if q := req.BrandID; q != "" {
		queries = append(queries, bleve.NewPhraseQuery([]string{q}, "BrandID"))
	}
	if q := req.ModelID; q != "" {
		queries = append(queries, bleve.NewPhraseQuery([]string{q}, "ModelID"))
	}

	searchRequest := bleve.NewSearchRequest(bleve.NewConjunctionQuery(queries...))
	if limit := req.Limit; limit != nil && limit.Value > 0 {
		searchRequest.Size = int(limit.Value)
	}
	if offset := req.Offset; offset != nil && offset.Value > 0 {
		searchRequest.From = int(offset.Value)
	}
	searchRequest.Fields = []string{"ModelPB"}
	switch req.OrderBy {
	case "brand_id":
		searchRequest.SortBy([]string{"BrandID"})
	case "-brand_id":
		searchRequest.SortBy([]string{"-BrandID"})
	case "model_id":
		searchRequest.SortBy([]string{"ModelID"})
	case "-model_id":
		searchRequest.SortBy([]string{"-ModelID"})
	case "name":
		searchRequest.SortBy([]string{"ModelName"})
	case "-name":
		searchRequest.SortBy([]string{"-ModelName"})
	}

	bl.modelsIndexMu.RLock()
	result, err := bl.modelsIndex.Search(searchRequest)
	bl.modelsIndexMu.RUnlock()
	if err != nil {
		return nil, err
	}

	models := make([]*ttnpb.EndDeviceModel, 0, len(result.Hits))
	for _, hit := range result.Hits {
		s, ok := hit.Fields["ModelPB"].(string)
		if !ok {
			return nil, errCorruptedIndex.New()
		}
		model := &ttnpb.EndDeviceModel{}
		if err := jsonpb.TTN().Unmarshal([]byte(s), model); err != nil {
			return nil, err
		}
		pb := &ttnpb.EndDeviceModel{}
		if err := pb.SetFields(model, req.Paths...); err != nil {
			return nil, err
		}
		models = append(models, pb)
	}
	return &store.ListModelsResponse{
		Count:  uint32(len(result.Hits)),
		Total:  uint32(result.Total),
		Offset: uint32(searchRequest.From),
		Models: models,
	}, nil
}

// GetTemplate retrieves an end device template for an end device definition.
func (bl *bleveStore) GetTemplate(ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	return bl.store.GetTemplate(ids)
}

// GetUplinkDecoder retrieves the codec for decoding uplink messages.
func (bl *bleveStore) GetUplinkDecoder(ids *ttnpb.EndDeviceVersionIdentifiers) (string, error) {
	return bl.store.GetUplinkDecoder(ids)
}

// GetDownlinkDecoder retrieves the codec for decoding downlink messages.
func (bl *bleveStore) GetDownlinkDecoder(ids *ttnpb.EndDeviceVersionIdentifiers) (string, error) {
	return bl.store.GetDownlinkDecoder(ids)
}

// GetDownlinkEncoder retrieves the codec for encoding downlink messages.
func (bl *bleveStore) GetDownlinkEncoder(ids *ttnpb.EndDeviceVersionIdentifiers) (string, error) {
	return bl.store.GetDownlinkEncoder(ids)
}
