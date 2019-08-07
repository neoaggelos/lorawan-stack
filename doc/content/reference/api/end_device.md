---
title: "End Device APIs"
description: ""
weight: 7
---

{{< proto/message package="ttn.lorawan.v3" message="EndDeviceIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="EndDevice" >}}

{{< proto/message package="ttn.lorawan.v3" message="EndDeviceVersionIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="Location" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="MACVersion" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="PHYVersion" >}}

# Shared Messages

{{< proto/message package="ttn.lorawan.v3" message="GetEndDeviceRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="SetEndDeviceRequest" >}}

# The Global End Device Registry (`EndDeviceRegistry` service)

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceRegistry" method="Create" >}}

{{< proto/message package="ttn.lorawan.v3" message="CreateEndDeviceRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceRegistry" method="Get" >}}

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceRegistry" method="List" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListEndDevicesRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceRegistry" method="Update" >}}

{{< proto/message package="ttn.lorawan.v3" message="UpdateEndDeviceRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceRegistry" method="Delete" >}}

# The Join Server (`JsEndDeviceRegistry` service)

{{< proto/method package="ttn.lorawan.v3" service="JsEndDeviceRegistry" method="Set" >}}

> TODO: What fields to set to the JS.

{{< proto/method package="ttn.lorawan.v3" service="JsEndDeviceRegistry" method="Get" >}}

> TODO: What fields to get from the JS.

{{< proto/method package="ttn.lorawan.v3" service="JsEndDeviceRegistry" method="Delete" >}}

{{< proto/message package="ttn.lorawan.v3" message="RootKeys" >}}

{{< proto/message package="ttn.lorawan.v3" message="EndDeviceAuthenticationCode" >}}

# The Network Server (`NsEndDeviceRegistry` service)

{{< proto/method package="ttn.lorawan.v3" service="NsEndDeviceRegistry" method="Set" >}}

> TODO: What fields to set to the NS.

{{< proto/method package="ttn.lorawan.v3" service="NsEndDeviceRegistry" method="Get" >}}

> TODO: What fields to get from the NS.

{{< proto/method package="ttn.lorawan.v3" service="NsEndDeviceRegistry" method="Delete" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACSettings" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACSettings.PingSlotPeriodValue" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="PingSlotPeriod" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACSettings.DataRateIndexValue" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="DataRateIndex" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACSettings.RxDelayValue" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="RxDelay" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACSettings.AggregatedDutyCycleValue" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="AggregatedDutyCycle" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACState" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACParameters" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="RejoinTimeExponent" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="RejoinCountExponent" >}}

{{< proto/message package="ttn.lorawan.v3" message="MACParameters.Channel" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="Class" >}}

{{< proto/message package="ttn.lorawan.v3" message="Session" >}}

{{< proto/message package="ttn.lorawan.v3" message="SessionKeys" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="PowerState" >}}

# The Application Server (`AsEndDeviceRegistry` service)

{{< proto/method package="ttn.lorawan.v3" service="AsEndDeviceRegistry" method="Set" >}}

> TODO: What fields to set to the AS.

{{< proto/method package="ttn.lorawan.v3" service="AsEndDeviceRegistry" method="Get" >}}

> TODO: What fields to get from the AS.

{{< proto/method package="ttn.lorawan.v3" service="AsEndDeviceRegistry" method="Delete" >}}

{{< proto/message package="ttn.lorawan.v3" message="MessagePayloadFormatters" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="PayloadFormatter" >}}

# Other

{{< proto/enum package="ttn.lorawan.v3" enum="LocationSource" >}}

{{< proto/message package="ttn.lorawan.v3" message="KeyEnvelope" >}}
