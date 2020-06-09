# Example of Optimizely Running in a Microservice Environment

1. Create a free [Optimizely Account](https://www.optimizely.com/rollouts-signup/?utm_source=youtube&utm_campaign=microservice-feature-flags)
2. Clone this repository
3. Navigate to the root directory of the repository:
```bash
cd optly-in-microservice-example
```
4. Setup and run one or more of the following example microservices by cloning this repo and then following instructions below:
 - [JavaScript Service](./javascript-service)
 - [Python Service](./python-service)
 - [Go Service](./go-service)

5. Run Optimizely as a service using [Optimizely Agent](https://docs.developers.optimizely.com/full-stack/docs/setup-optimizely-agent):
 - Install [Golang v1.13+](https://golang.org/dl/)
 - Clone the [Optimizely Agent repo](https://github.com/optimizely/agent).
 - Change to the agent directory:
 ```bash
 cd agent
 ```
 - Setup Optimizely Agent with:
 ```bash
 make setup
 ```
 - Run Optimizely Agent with:
 ```bash
 make run
 ```
 - For debugging during development, change the `pollingInterval` in the `config.yaml` to `10s`


6. Create a feature flag in [Optimizely](https://app.optimizely.com)

7. Follow the instructions below to update one or more of the microservices to evaluate the feature flag using Optimizley Agent's [activate endpoint](https://docs.developers.optimizely.com/full-stack/docs/use-optimizely-agent#section-manage-features): `POST /v1/activate?featureKey={featureKey}`
 - [JavaScript Changes](./javascript-service/README.md#evaluate-the-feature-flag)
 - [Python Changes](./python-service/README.md#evaluate-the-feature-flag)
 - [Go Changes](./go-service/README.md#evaluate-the-feature-flag)

8. Turn the feature flag on.
9. Profit 🎉 !
