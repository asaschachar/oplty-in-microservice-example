# Example of Optimizely Running in a Microservice Environment

1. Create a free [Optimizely Account](https://www.optimizely.com/rollouts-signup/?utm_source=youtube&utm_campaign=microservice-feature-flags)
2. Setup and run one or more of the following example microservices:
 - [JavaScript Service](./javascript-service)
 - [Python Service](./python-service)
 - [Go Service](./go-service)

3. Run [Optimizely as a service](https://docs.developers.optimizely.com/full-stack/docs/setup-optimizely-agent):
 - Install [Golang v1.13+](https://golang.org/dl/)
 - Clone the [Optimizely Agent repo](https://github.com/optimizely/agent).
 - For debugging during development, change the `pollingInterval` in the `config.yaml` to `10s`
 - From the Agent repo directory, setup & run Optimizely Agent with:
 ```bash
 make setup
 ```
 Then
 ```bash
 make run
 ```

4. Create a feature flag in [Optimizely](https://app.optimizely.com)

5. Update one or more of the microservices to evaluate the feature flag:
 - [JavaScript Handler](./javascript-service/README.md#evaluate-the-feature-flag)
 - [Python Handler](./python-service/README.md#evaluate-the-feature-flag)
 - [Go Handler](./go-service/README.md#evaluate-the-feature-flag)

6. Turn the feature flag on.
7. Profit 🎉 !
