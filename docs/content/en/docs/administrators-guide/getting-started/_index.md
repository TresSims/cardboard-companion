---
title: Getting Started
description: How do I get a CardboardCompanion running as fast as humanly possible?
categories: [Examples]
weight: 2
---

The CardboardCompanion is published as a docker image, and the CardboardCompanion team
reccomends that you launch it using docker-compose. An example docker compose file is
shown below:

docker-compose.yaml

```yaml
services:
  cardboard-companion:
    image: ghcr.io/tressims/cardboard-companion:latest
    volumes:
      - ./cardboard-companion:/code/cardboard-companion.yml
    environment:
      - TZ=America/New_York # Helpful if you are scheduling cron-jobs
```

The `cardboard-companion.yml` file contains the following minimal structure:

```yaml
Token: <your-bot-token>
GuildID: <your-guild-id>
PollSchedule: "0 9 * * 4" # Send a poll at 9:00 am on Thursday
PollChannel: "<your-channel-to-post-a-cron-poll>"
```

Once you have those two files nice and cozy (in the same directory, wherever), all that's
left is to run:

```bash
docker-compose up --build -d
```
