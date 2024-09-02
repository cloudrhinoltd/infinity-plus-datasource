# InfinityPlus Datasource for Grafana

InfinityPlus is a modified version of the original Grafana Infinity Datasource by Grafana Labs, with additional features and improvements by Cloud Rhino. This plugin allows you to visualize data from various data sources in Grafana.

**Important:** This README references the original documentation for the Infinity Datasource, which largely applies to InfinityPlus due to its shared functionality. For new features added by Cloud Rhino, separate documentation is provided below.

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

## 📥 Download

**Please note:** The download links below will change once InfinityPlus is published in the Grafana plugin catalog. For now, refer to the links below for the original Infinity Datasource until InfinityPlus is officially released.

- From [Grafana plugin catalog](https://grafana.com/grafana/plugins/infinity-plus-datasource/)
- From [GitHub release page](https://github.com/grafana/grafana-infinity-datasource/releases) (includes beta and pre-release versions)
- Using Grafana CLI:
  - `grafana-cli plugins install infinity-plus-datasource`
- Using Docker:
  - `docker run -p 3000:3000 -e "GF_INSTALL_PLUGINS=infinity-plus-datasource" grafana/grafana:latest`

Once InfinityPlus is published, the correct installation instructions and download links will be updated here.

## 📖 Documentation

The documentation for InfinityPlus is primarily based on the original Infinity Datasource by Grafana Labs. The full documentation for the original datasource, which covers core features and usage, can be found here:

- [Original Plugin Documentation](https://grafana.com/docs/plugins/yesoreyeram-infinity-datasource) — applicable to both InfinityPlus and the original plugin.

### **Credits:**

The original documentation and plugin were developed by Grafana Labs. InfinityPlus builds upon this foundation, with new features and updates added by Cloud Rhino.

For documentation specific to the new features added by Cloud Rhino, please refer to the additional resources provided within this repository or contact Cloud Rhino directly.

## ⚡️ Useful Links

- [Plugin documentation](https://grafana.com/docs/plugins/yesoreyeram-infinity-datasource) (original Infinity Datasource by Grafana Labs)
- [Demo video](https://youtu.be/Wmgs1E9Ry-s)

## 👍 Contributing

Contributions to InfinityPlus are welcome! Here are some ways you can contribute:

- Showcase how you are using APIs with InfinityPlus Datasource so that other community members can benefit.
- Test different APIs and report bugs if things don't work as expected.
- Refer to the [contributing guide](https://github.com/grafana/grafana-infinity-datasource/blob/main/CONTRIBUTING.md) from the original project for setup and contribution guidelines.

## ⭐️ Project Assistance

If you want to say **thank you** or support the active development of `InfinityPlus Datasource`:

- Add a [GitHub Star](https://github.com/grafana/grafana-infinity-datasource) to the original project.
- Share the project on [Twitter](https://twitter.com/intent/tweet?text=Checkout%20this%20cool%20%23grafana%20datasource%20%40grafanainfinity.%20%0A%0ALiterally,%20get%20your%20data%20from%20anywhere%20into%20%23grafana.%20JSON,%20CSV,%20XML,%20GraphQL,%20OAuth2,%20RSS%20feed,%20%23kubernetes,%20%23azure,%20%23aws,%20%23gcp%20and%20more%20stuff.%0A%0Ahttps%3A//grafana.com/docs/plugins/yesoreyeram-infinity-datasource%0A).
- Write articles about the project on [Dev.to](https://dev.to/), [Medium](https://medium.com/), or your personal blog.

## ⚠️ License

InfinityPlus Datasource is a derivative of the Grafana Infinity Datasource by Grafana Labs, which is licensed under [Apache 2.0](https://github.com/grafana/grafana-infinity-datasource/blob/main/LICENSE). In compliance with the license:

- This project continues to follow the terms of the Apache 2.0 license.
- Credits and documentation links are provided to acknowledge the original authors and maintain transparency.

InfinityPlus Datasource includes modifications and additional features implemented by Cloud Rhino.
