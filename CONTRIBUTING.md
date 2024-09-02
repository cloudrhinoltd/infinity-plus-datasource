# Contributing to InfinityPlus

Thank you for your interest in contributing to InfinityPlus! This plugin is based on the original Grafana Infinity Datasource. While we have added new features, much of the core functionality remains similar. Below are some ways you can contribute:

- Test different APIs and create issues if things are not working as expected. Since InfinityPlus is based on the original, many existing API tests might still be relevant.
- If you discover interesting APIs or use cases, consider sharing how you are using InfinityPlus so that other community members can benefit from your experience.

## Setting Up Locally for Development

You will need the following tools installed on your local machine:

- NodeJS v20.0+
- Go 1.22
- Mage

To set up the development environment, follow these steps after cloning the repo locally in Grafana's plugin folder:

- Run `yarn` to install the frontend dependencies (only needed once).
- Run `yarn test` to ensure all existing tests pass.
- Run `go test -v ./...` to ensure all backend tests pass.
- Use `yarn dev` for continuous frontend build monitoring.
- Use `yarn build` to build the frontend components.
- Run `mage -v` to build the backend part of the plugin (only needed if contributing to the backend).
- Use `docker-compose up` to run the plugin with Grafana locally (use `infinity:infinity` as the credentials). Debugging and tracing can be enabled; refer to the **Setting up Grafana in Debug Mode** section below.

## Submitting PRs

If you are submitting a pull request, make sure to:

- Run `yarn changeset` from your branch to create a changelog entry.
- Provide the necessary details; this will generate a markdown file inside the `./.changeset` folder.
- During the release, the changelog will be updated and the package version will be bumped based on these changesets.

## Releasing & Bumping Version

To create a new release, execute `yarn changeset version`. This will update the changelog and bump the version in the `package.json` file. Be sure to commit these changes.

## Setting Up Grafana in Debug Mode

- Install the Loki Docker plugin using: `docker plugin install grafana/loki-docker-driver:2.9.1 --alias loki --grant-all-permissions`
- Start Docker using the debug configuration: `docker compose -f docker-compose-debug.yaml up`

## Testing the PDC

To test the PDC functionality with InfinityPlus, you can use the debug Docker setup: `docker compose -f docker-compose-debug.yaml up`. This configuration includes **microsocks** proxy, PDC enabled, and configured settings. Example datasource instances with secure socks proxy enabled and various authentication mechanisms can be found in the [provisioned datasources](./provisioning/datasources/default.yml) file. Look for the PDC-enabled datasources with the prefix **PDC**.

---

**Credits**: InfinityPlus is based on the [Grafana Infinity Datasource](https://github.com/grafana/grafana-infinity-datasource) by Grafana Labs. For specific API testing and showcases, you may still refer to the discussions and community resources of the original plugin.
