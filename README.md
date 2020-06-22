calendar-notifier
===

![](https://github.com/ww24/calendar-notifier/workflows/Test%20on%20master/badge.svg)

Calendar Notifier is Google Calendar event handler and actions.

## Features

- Events
  - [x] [Calendar Events List API](https://developers.google.com/calendar/v3/reference/events/list)
  - [ ] [Calendar Events Watch API](https://developers.google.com/calendar/v3/reference/events/watch)
- Actions
  - [x] HTTP Action
  - [x] [Cloud Pub/Sub](https://cloud.google.com/pubsub/) Action
  - [x] [Cloud Tasks](https://cloud.google.com/tasks/) Action

## Setup

- Create GCP Project
  - https://cloud.google.com/resource-manager/docs/creating-managing-projects
- Create GCP Service Account
  - Role assignment is not required if it does not use pubsub action.
  - https://cloud.google.com/iam/docs/creating-managing-service-accounts
- Share Google Calendar with service account
  - Use service account e-mail address. (e.g. `my-service-account@project-id.iam.gserviceaccount.com`)
  - https://support.google.com/calendar/answer/37082
- Set `SERVICE_ACCOUNT` environment. Use base64 encoded GCP Service Account JSON.
  - `echo "SERVICE_ACCOUNT=$(base64 < service_account.json)" > .env`
- Edit `config.sample.yml` and save as `config.yml`.
- Set `CONFIG` environment. Use base64 encoded config.yml.
  - `echo "CONFIG=$(base64 < config.yml)" > .env`

## Usage

### For the docker-compose

- Run `docker-compose up`

### For the docker

- Run `docker build -t calendar-notifier .`
- Run `docker run -e SERVICE_ACCOUNT=$(base64 < service_account.json) -e CONFIG=$(base64 < config.yml) calendar-notifier`
