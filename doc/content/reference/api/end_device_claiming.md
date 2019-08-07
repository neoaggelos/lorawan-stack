---
title: "End Device Claiming APIs"
description: ""
weight: 8
---

# End Device Claiming (`EndDeviceClaimingServer` service)

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceClaimingServer" method="AuthorizeApplication" >}}

{{< proto/message package="ttn.lorawan.v3" message="AuthorizeApplicationRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceClaimingServer" method="UnauthorizeApplication" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationIdentifiers" >}}

{{< proto/method package="ttn.lorawan.v3" service="EndDeviceClaimingServer" method="Claim" >}}

{{< proto/message package="ttn.lorawan.v3" message="ClaimEndDeviceRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="ClaimEndDeviceRequest.AuthenticatedIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="EndDeviceIdentifiers" >}}
