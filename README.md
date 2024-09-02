# InfinityPlus Datasource

InfinityPlus is a Grafana data source plugin based on the original [Grafana Infinity Datasource](https://grafana.com/grafana/plugins/yesoreyeram-infinity-datasource/). It extends the functionality of the original plugin with additional features while maintaining the core capabilities of visualizing data from JSON, CSV, XML, GraphQL, and HTML endpoints.

> **Important Note**: This plugin, InfinityPlus, is a separate plugin based on the original Grafana Infinity Datasource, developed by Cloud Rhino. The majority of the documentation for InfinityPlus is based on the original documentation, which largely applies here as well. However, please note that any new features introduced by InfinityPlus are not covered in the original documentation.

## 🎯 Key Features

### Original Features (from Grafana Infinity Datasource):

- Get data from multiple sources into Grafana.
- Supports various data formats:
  - JSON
  - CSV / TSV / any delimited format
  - XML
  - GraphQL
  - HTML
  - RSS/ATOM
- Supports various authentication methods:
  - Basic authentication
  - Bearer token authentication
  - API Key authentication
  - OAuth passthrough
  - OAuth2 client credentials
  - OAuth2 JWT authentication
  - AWS/Azure/GCP authentication
  - Digest authentication
- Flexible data manipulation with UQL, GROQ, JSONata.
- Supports alerting, recorded queries, public dashboards, and query caching.
- Utility variable functions.
- Supports Grafana node graph panel, annotations, and more.

### New Features Added by Cloud Rhino:

- **Azure Managed Identity authentication**: Enhanced security and seamless integration with Azure services using managed identities for authentication.

## 📖 Documentation

For comprehensive documentation, refer to the [original plugin documentation](https://grafana.com/docs/plugins/yesoreyeram-infinity-datasource). This documentation covers most of the features and functionality that are also present in InfinityPlus.

### Key Points

- The linked documentation is for the original Grafana Infinity Datasource and not specifically for InfinityPlus.
- Most features are shared, but the additional features introduced in InfinityPlus are not covered in the original documentation. For information on these new features, refer to the additional resources provided by Cloud Rhino or reach out to Cloud Rhino directly.

### [Demo video](https://youtu.be/Wmgs1E9Ry-s)

## ⚠️ Known Limitations

- Backend features such as Alerting, Recorded Queries, Enterprise query caching, and public dashboards work only if the `backend` parser option is selected in queries, similar to the original plugin.

## 📜 License

InfinityPlus is licensed under the Apache License 2.0, just like the original Grafana Infinity Datasource. This plugin includes modifications and additional features by Cloud Rhino and is distributed in compliance with the original license terms.

For more information about the licensing, see the [LICENSE](./LICENSE) file in this repository.

## 🌟 Credits

- **Original Plugin**: [Grafana Infinity Datasource](https://grafana.com/grafana/plugins/yesoreyeram-infinity-datasource/) by Grafana Labs.
- **Author of InfinityPlus**: Cloud Rhino.

InfinityPlus builds upon the foundation laid by Grafana Labs, adding new capabilities while respecting the original work and licensing.
