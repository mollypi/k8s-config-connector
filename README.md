# GCP Config Connector

[![Git](https://app.soluble.cloud/api/v1/public/badges/3fc34cb8-ca5e-4893-8fe8-48968309ffad.svg?orgId=561911742905)](https://app.soluble.cloud/repos/details/github.com/mollypi/k8s-config-connector?orgId=561911742905)  [![CIS](https://app.soluble.cloud/api/v1/public/badges/343fc6db-6911-4e9c-8ddc-fd60b25147ed.svg?orgId=561911742905)](https://app.soluble.cloud/repos/details/github.com/mollypi/k8s-config-connector?orgId=561911742905)  [![IaC](https://app.soluble.cloud/api/v1/public/badges/a0de1bac-2c33-40c0-aae3-018e61dc2f0a.svg?orgId=561911742905)](https://app.soluble.cloud/repos/details/github.com/mollypi/k8s-config-connector?orgId=561911742905)  

Config Connector is a Kubernetes add-on that allows customers to manage GCP
resources, such as Cloud Spanner or Cloud Storage, through your cluster's API.

With Config Connector, now you can describe GCP resources declaratively using
Kubernetes-style configuration. Config Connector will create any new GCP
resources and update any existing ones to the state specified by your
configuration, and continuously makes sure GCP is kept in sync. The same
resource model is the basis of Istio, Knative, Kubernetes, and the Google Cloud
Services Platform.

As a result, developers can manage their whole application, including both its
Kubernetes components as well as any GCP dependencies, using the same
configuration, and -- more importantly -- *tooling*. For example, the same
customization or templating tool can be used to manage test vs. production
versions of an application across both Kubernetes and GCP.

This repository contains Config Connector CRDs, install bundles, and sample
resource configurations.

For simple starter
examples, please refer to [Cloud Foundation Toolkit Config Connector Solutions](https://github.com/GoogleCloudPlatform/cloud-foundation-toolkit/tree/master/config-connector/solutions).

For more information about Config Connector, see
https://cloud.google.com/config-connector/docs/overview.
