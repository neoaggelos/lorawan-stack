---
title: "Application Webhook APIs"
description: ""
weight: 3
---

{{< proto/message package="ttn.lorawan.v3" message="ApplicationWebhookIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationWebhook" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationWebhookTemplateIdentifiers" >}}

# The Application Webhook Registry (`ApplicationWebhookRegistry` service)

{{< proto/message package="ttn.lorawan.v3" message="ApplicationWebhook.Message" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationWebhookRegistry" method="GetFormats" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationWebhookFormats" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationWebhookRegistry" method="Set" >}}

{{< proto/message package="ttn.lorawan.v3" message="SetApplicationWebhookRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationWebhookRegistry" method="Get" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetApplicationWebhookRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationWebhookRegistry" method="List" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListApplicationWebhooksRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="ApplicationWebhooks" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationWebhookRegistry" method="Delete" >}}
