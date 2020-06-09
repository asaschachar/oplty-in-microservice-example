# Example of Optimizely Running in a Microservice Environment

1. Create a free [Optimizely Account](https://www.optimizely.com/rollouts-signup/?utm_source=youtube&utm_campaign=microservice-feature-flags)
2. Run one or more of the following example microservices:
 - [JavaScript Service](./javascript-service)
 - [Python Service](./python-service)
 - [Go Service](./go-service)

3. Run [Optimizely as a service](https://docs.developers.optimizely.com/full-stack/docs/setup-optimizely-agent):
 - Install [Golang v1.13+](https://golang.org/dl/)
 - Clone the [Optimizely Agent repo](https://github.com/optimizely/agent).
 - For debugging during development, change the polling interval in the config.yaml to `10s`
 - From the Agent repo directory, setup Optimizely Agent with:
 ```bash
 make setup
 ```
 - Run Optimizely agent with:
 ```bash
 make run
 ```

4. Create a feature flag in [Optimizely](https://app.optimizely.com)

5. Update one or more of the microservices to evaluate the feature flag:
 - [JavaScript Handler](./javascript-service/handler.js)
 - [Python Handler](./python-service/handler.js)
 - [Go Handler](./go-service/handler.js)

6. Profit 🎉 !
