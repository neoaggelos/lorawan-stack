---
title: "Gateway APIs"
description: ""
weight: 5
---

{{< proto/message package="ttn.lorawan.v3" message="GatewayIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="Gateway" >}}

# The Gateway Registry (`GatewayRegistry` service)

{{< proto/method package="ttn.lorawan.v3" service="GatewayRegistry" method="Create" >}}

{{< proto/message package="ttn.lorawan.v3" message="CreateGatewayRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayRegistry" method="Get" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetGatewayRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayRegistry" method="List" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListGatewaysRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayRegistry" method="Update" >}}

{{< proto/message package="ttn.lorawan.v3" message="UpdateGatewayRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayRegistry" method="Delete" >}}

# Gateway Access Management (`GatewayAccess` service)

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="ListRights" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="CreateAPIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="CreateGatewayAPIKeyRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="ListAPIKeys" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListGatewayAPIKeysRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="GetAPIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetGatewayAPIKeyRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="UpdateAPIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="UpdateGatewayAPIKeyRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="GetCollaborator" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetGatewayCollaboratorRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetCollaboratorResponse" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="SetCollaborator" >}}

{{< proto/message package="ttn.lorawan.v3" message="SetGatewayCollaboratorRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="GatewayAccess" method="ListCollaborators" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListGatewayCollaboratorsRequest" >}}

# Listing available frequency plans (`Configuration` service)

The Gateway Server exposes the list of available frequency plans with the `Configuration` service.

{{< proto/method package="ttn.lorawan.v3" service="Configuration" method="ListFrequencyPlans" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListFrequencyPlansRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListFrequencyPlansResponse" >}}

{{< proto/message package="ttn.lorawan.v3" message="FrequencyPlanDescription" >}}

# Other

{{< proto/message package="ttn.lorawan.v3" message="OrganizationOrUserIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="OrganizationIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="UserIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="ContactInfo" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="ContactType" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="ContactMethod" >}}

{{< proto/message package="ttn.lorawan.v3" message="APIKeys" >}}

{{< proto/message package="ttn.lorawan.v3" message="APIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="Collaborators" >}}

{{< proto/message package="ttn.lorawan.v3" message="Collaborator" >}}

{{< proto/message package="ttn.lorawan.v3" message="Rights" >}}

{{< proto/enum package="ttn.lorawan.v3" enum="Right" >}}
