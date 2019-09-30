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

package commands

import (
	"os"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/io"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

var (
	selectApplicationPackageAssociationsFlags = util.FieldMaskFlags(&ttnpb.ApplicationPackageAssociation{})
	setApplicationPackageAssociationsFlags    = util.FieldFlags(&ttnpb.ApplicationPackageAssociation{})
)

func applicationPackageAssociationIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("application-id", "", "")
	flagSet.String("device-id", "", "")
	flagSet.String("f-port", "", "")
	return flagSet
}

func getApplicationPackageAssociationID(flagSet *pflag.FlagSet, args []string) (*ttnpb.ApplicationPackageAssociationIdentifiers, error) {
	applicationID, _ := flagSet.GetString("application-id")
	deviceID, _ := flagSet.GetString("device-id")
	fport, _ := flagSet.GetUint32("f-port")
	switch len(args) {
	case 0:
	case 1:
	case 2:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 3:
		applicationID = args[0]
		deviceID = args[1]
		fport64, _ := strconv.ParseUint(args[2], 10, 32)
		fport = uint32(fport64)
	default:
		logger.Warn("multiple IDs found in arguments, considering the first")
		applicationID = args[0]
		deviceID = args[1]
		fport64, _ := strconv.ParseUint(args[2], 10, 32)
		fport = uint32(fport64)
	}
	if applicationID == "" {
		return nil, errNoApplicationID
	}
	if deviceID == "" {
		return nil, errNoEndDeviceID
	}
	return &ttnpb.ApplicationPackageAssociationIdentifiers{
		EndDeviceIdentifiers: ttnpb.EndDeviceIdentifiers{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{ApplicationID: applicationID},
			DeviceID:               deviceID,
		},
		FPort: fport,
	}, nil
}

var (
	applicationsPackagesCommand = &cobra.Command{
		Use:     "packages",
		Aliases: []string{"package"},
		Short:   "Application packages commands",
	}
	applicationsPackagesGetAvailableCommand = &cobra.Command{
		Use:     "get-available-packages",
		Aliases: []string{"available-packages"},
		Short:   "Get the available packages for the device",
		RunE: func(cmd *cobra.Command, args []string) error {
			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).GetPackages(ctx, devID)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackagesGetCommand = &cobra.Command{
		Use:     "get [application-id] [device-id] [f-port]",
		Aliases: []string{"info"},
		Short:   "Get the properties of an application package association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPackageAssociationsFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPackageAssociationsFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).Get(ctx, &ttnpb.GetApplicationPackageAssociationRequest{
				ApplicationPackageAssociationIdentifiers: *assocID,
				FieldMask:                                types.FieldMask{Paths: paths},
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackagesListCommand = &cobra.Command{
		Use:     "list [application-id] [device-id]",
		Aliases: []string{"ls"},
		Short:   "List application package associations",
		RunE: func(cmd *cobra.Command, args []string) error {
			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPackageAssociationsFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPackageAssociationsFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).List(ctx, &ttnpb.ListApplicationPackageAssociationRequest{
				EndDeviceIdentifiers: *devID,
				FieldMask:            types.FieldMask{Paths: paths},
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackagesSetCommand = &cobra.Command{
		Use:     "set [application-id] [device-id] [f-port]",
		Aliases: []string{"update"},
		Short:   "Set the properties of an application package association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.UpdateFieldMask(cmd.Flags(), setApplicationPackageAssociationsFlags, headersFlags())

			var association ttnpb.ApplicationPackageAssociation
			if err = util.SetFields(&association, setApplicationPackageAssociationsFlags); err != nil {
				return err
			}
			association.ApplicationPackageAssociationIdentifiers = *assocID

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).Set(ctx, &ttnpb.SetApplicationPackageAssociationRequest{
				ApplicationPackageAssociation: association,
				FieldMask:                     types.FieldMask{Paths: paths},
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackagesDeleteCommand = &cobra.Command{
		Use:   "delete [application-id] [device-id] [f-port]",
		Short: "Delete an application package association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewApplicationPackageRegistryClient(as).Delete(ctx, assocID)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	applicationsPackagesCommand.AddCommand(applicationsPackagesGetAvailableCommand)
	applicationsPackagesGetCommand.Flags().AddFlagSet(applicationPackageAssociationIDFlags())
	applicationsPackagesGetCommand.Flags().AddFlagSet(selectApplicationPackageAssociationsFlags)
	applicationsPackagesCommand.AddCommand(applicationsPackagesGetCommand)
	applicationsPackagesListCommand.Flags().AddFlagSet(endDeviceIDFlags())
	applicationsPackagesListCommand.Flags().AddFlagSet(selectApplicationPackageAssociationsFlags)
	applicationsPackagesCommand.AddCommand(applicationsPackagesListCommand)
	applicationsPackagesSetCommand.Flags().AddFlagSet(applicationPackageAssociationIDFlags())
	applicationsPackagesSetCommand.Flags().AddFlagSet(setApplicationPackageAssociationsFlags)
	applicationsPackagesCommand.AddCommand(applicationsPackagesSetCommand)
	applicationsPackagesDeleteCommand.Flags().AddFlagSet(applicationPackageAssociationIDFlags())
	applicationsPackagesCommand.AddCommand(applicationsPackagesDeleteCommand)
	applicationsCommand.AddCommand(applicationsPackagesCommand)
}
