// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package devicetemplates

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
)

type euiCSV struct {
}

func (euiCSV) Format() *ttnpb.EndDeviceTemplateFormat {
	return &ttnpb.EndDeviceTemplateFormat{
		Name:        "EUI CSV",
		Description: "Converts comma-delimited JoinEUI and DevEUI",
	}
}

var errEUICSVFormat = errors.DefineInvalidArgument("eui_csv_format", "invalid EUI CSV format")

func (euiCSV) Convert(ctx context.Context, r io.Reader, ch chan<- *ttnpb.EndDeviceTemplate) error {
	defer close(ch)

	s := bufio.NewScanner(r)
	for s.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		parts := strings.Split(s.Text(), ",")
		if len(parts) < 2 {
			return errEUICSVFormat
		}

		var joinEUI, devEUI types.EUI64
		if err := joinEUI.UnmarshalText([]byte(parts[0])); err != nil {
			return errEUICSVFormat.WithCause(err)
		}
		if err := devEUI.UnmarshalText([]byte(parts[1])); err != nil {
			return errEUICSVFormat.WithCause(err)
		}

		ch <- &ttnpb.EndDeviceTemplate{
			EndDevice: ttnpb.EndDevice{
				EndDeviceIdentifiers: ttnpb.EndDeviceIdentifiers{
					DeviceID: strings.ToLower(fmt.Sprintf("eui-%s", devEUI)),
					JoinEUI:  &joinEUI,
					DevEUI:   &devEUI,
				},
			},
			FieldMask: pbtypes.FieldMask{
				Paths: []string{
					"ids.device_id",
					"ids.dev_eui",
					"ids.join_eui",
				},
			},
			MappingKey: devEUI.String(),
		}
	}
	return s.Err()
}

func init() {
	RegisterConverter("eui-csv", new(euiCSV))
}
