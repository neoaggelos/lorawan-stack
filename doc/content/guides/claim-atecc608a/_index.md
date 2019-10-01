---
title: "Claim ATECC608A Secure Elements"
description: ""
weight: 1000
---

{{< cli-only >}}

The Things Industries and Microchip developed a security solution for LoRaWAN that enables secure key provisioning and secures cryptographic operations carried out by devices. [Learn more](https://www.thethingsindustries.com/technology/security-solution)

This guide helps device makers to claim secure elements on TTI Join Server. This guide also makes devices available for claiming by device owners to onboard them on supported LoRaWAN Network and Application Servers.

## Prerequisites

1. ATECC608A-TNGLORA devices. [Product details](https://www.microchip.com/wwwproducts/en/ATECC608A-TNGLORA)
2. Device security (manifest) file. You can obtain this from your [Microchip Direct order history](https://www.microchipdirect.com/orders)
3. Account on TTI Join Server. [Create account](https://join.thethings.industries/oauth/register)
4. Application on TTI Join Server. [Create application](https://join.thethings.industries/applications/add)
5. TTI Join Server CLI configuration. [Download configuration file](.ttn-lw-cli.yml)
6. The Things Stack CLI on your local machine. [Installation instructions]({{< ref "/guides/getting-started/installation" >}})

## Working directory

For this guide, we assume you have a working directory with the following files.

```
.
├── .ttn-lw-cli.yml     # TTI Join Server CLI configuration (prerequisite #5)
└── manifest.json       # Manifest downloaded from Microchip Direct (prerequisite #2)
```

## Login with CLI

Open a terminal session, go to the working directory and login:

```bash
$ ttn-lw-cli login
```

This opens the TTI Join Server login screen. Login with your username and password. Once you logged in in the browser, return to the terminal session to proceed.

## Convert the manifest

Convert the manifest to device templates:

```bash
$ ttn-lw-cli end-device template from-data \
  --format-id microchip-atecc608a-tnglora \
  --local-file manifest.json > templates.json
```

This creates a `templates.json` with end device templates for the devices to be created. [Learn more about end device templates]({{< ref "/concepts/end-device-templates" >}})

## Create the devices

You can execute these templates to create the end devices in your application:

```bash
$ templates.json \
  | ttn-lw-cli end-device template execute \
  | ttn-lw-cli end-device create --application-id test-app --with-claim-authentication-code
```

>Note: Replace `test-app` with the application ID of your application (prerequisite #4).

This creates the devices in your application on TTI Join Server with a claim authentication code. This is a secret code that device owners can use to claim the devices, which transfers the device from your application to the device owner's application.

## Creating QR codes

```bash
$ ttn-lw-cli end-device list test-app \
  | ttn-lw-cli end-device generate-qr
```
