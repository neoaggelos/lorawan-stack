---
title: "Application Server APIs"
description: ""
weight: 2
---

{{< proto/message package="ttn.lorawan.v3" message="ApplicationLink" >}}

## Link Configuration (`As` service)

{{< proto/method package="ttn.lorawan.v3" service="As" method="SetLink" >}}

{{< proto/message package="ttn.lorawan.v3" message="SetApplicationLinkRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="As" method="GetLink" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetApplicationLinkRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="As" method="GetLinkStats" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationLinkStats" >}}

{{< proto/method package="ttn.lorawan.v3" service="As" method="DeleteLink" >}}

## Downlink Queue Management (`AppAs` service)

{{< proto/message package="ttn.lorawan.v3" message="DownlinkQueueRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationDownlink" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationDownlink.ClassBC" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="TxSchedulePriority" >}}

{{< proto/method package="ttn.lorawan.v3" service="AppAs" method="DownlinkQueuePush" >}}

{{< proto/method package="ttn.lorawan.v3" service="AppAs" method="DownlinkQueueReplace" >}}

{{< proto/method package="ttn.lorawan.v3" service="AppAs" method="DownlinkQueueList" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationDownlinks" >}}

# Other

{{< proto/message package="ttn.lorawan.v3" message="ApplicationIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="EndDeviceIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="GatewayAntennaIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="GatewayIdentifiers" >}}
