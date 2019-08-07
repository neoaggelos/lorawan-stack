---
title: "Application Pub-Sub APIs"
description: ""
weight: 4
---

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSubIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSub" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSub.NATSProvider" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSub.MQTTProvider" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="ApplicationPubSub.MQTTProvider.QoS" >}}

# The Application Pub-Sub Registry (`ApplicationPubSubRegistry` service)

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSub.Message" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationPubSubRegistry" method="GetFormats" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSubFormats" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationPubSubRegistry" method="Set" >}}

{{< proto/message package="ttn.lorawan.v3" message="SetApplicationPubSubRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationPubSubRegistry" method="Get" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetApplicationPubSubRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationPubSubRegistry" method="List" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListApplicationPubSubsRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationPubSubs" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationPubSubRegistry" method="Delete" >}}
