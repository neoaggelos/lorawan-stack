---
title: "Application APIs"
description: ""
weight: 1
---

{{< proto/message package="ttn.lorawan.v3" message="ApplicationIdentifiers" >}}

{{< proto/message package="ttn.lorawan.v3" message="Application" >}}

# The Application Registry (`ApplicationRegistry` service)

{{< proto/method package="ttn.lorawan.v3" service="ApplicationRegistry" method="Create" >}}

{{< proto/message package="ttn.lorawan.v3" message="CreateApplicationRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationRegistry" method="Get" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetApplicationRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationRegistry" method="List" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListApplicationsRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationRegistry" method="Update" >}}

{{< proto/message package="ttn.lorawan.v3" message="UpdateApplicationRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationRegistry" method="Delete" >}}

# Application Access Management (`ApplicationAccess` service)

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="ListRights" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="CreateAPIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="CreateApplicationAPIKeyRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="ListAPIKeys" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListApplicationAPIKeysRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="GetAPIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetApplicationAPIKeyRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="UpdateAPIKey" >}}

{{< proto/message package="ttn.lorawan.v3" message="UpdateApplicationAPIKeyRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="GetCollaborator" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetApplicationCollaboratorRequest" >}}

{{< proto/message package="ttn.lorawan.v3" message="GetCollaboratorResponse" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="SetCollaborator" >}}

{{< proto/message package="ttn.lorawan.v3" message="SetApplicationCollaboratorRequest" >}}

{{< proto/method package="ttn.lorawan.v3" service="ApplicationAccess" method="ListCollaborators" >}}

{{< proto/message package="ttn.lorawan.v3" message="ListApplicationCollaboratorsRequest" >}}

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
