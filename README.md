<!--
Copyright 2022 Orix Au Yeung

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

*⚠️Note: This is a fork from [google/gcp-fwrule-open-domain](https://github.com/google/gcp-fwrule-open-domain), which is now archived. Any unmodified content from the original fork is owned by Google LLC and made public for use under Apache 2.0 License.⚠️*

# GCP Cloud Run Jobs module for using a domain name as a f/w rule source

Containerized BASH script for creating and updating a firewall rule using a domain name as the source ip for a firewall rule.

## Improvements made in this fork
1. **Lower CPU/RAM consumption**: Make use of Cloud Run Jobs (in Preview), remove the need to create a server listening to incoming HTTP requests.
2. **More efficient**: Cloud Run Job containers will self-terminate immediately after completion, no waiting period for container termination.
3. **Greater flexibilities**: Removed hardcoded configurations such as deployment location and timezone.

![Architecture](Architecture.png)

## Disclaimer

1. This is not an officially supported Google product.
1. Proper testing should be done before running this tool in production.
1. Should the DNS servers being queried become exploited or not within the control of the customer, the firewall would then become at risk.

## License Summary

This sample code is made available under Apache 2 license. See the LICENSE file.

## GCP Costs

GCP Cloud Run Invocation and Cloud Scheduler Jobs  
(Usually Free <https://cloud.google.com/run/pricing#tables>, <https://cloud.google.com/scheduler/pricing>)

## Prerequisites

1. Setup GCP SDK or use the Cloud Shell - <https://cloud.google.com/sdk>, <https://cloud.google.com/shell>

1. Initialize the SDK for the target account

   * `gcloud init` - <https://cloud.google.com/sdk/docs/initializing>

1. Run prerequisite.sh to create the required IAM objects and enable the required APIs.  

    `sh prerequisite.sh`

## Building

1. Set the ENV variable for the Cloud Run Service name so that N different rules can be made for different domain names.  

    `DOMAIN_FW_NAME='openfwusingdomain'`  
    `export DOMAIN_FW_NAME`

1. Export the PROJECT_ID env variable to make it available for build and deploy shell scripts.  

    `export PROJECT_ID`  

1. Run build.sh to build a new cloud run container image.

    `sh build.sh`

## Deploying

1. Update the env values and schedule frequency in the deploy.sh for your f/w rules settings.

    example:

    PRIORITY='1000'  
    RULES='tcp:80,tcp:8080,udp:8000'  
    DOMAIN='reddit.com'  
    TARGETTAGS='tags1'  
    NETWORK='default'  
    FREQUENCY='0 */1 * * *'  

1. Run deploy.sh to deploy to Cloud Run and create a new Cloud Scheduler Job.

    `sh deploy.sh`
