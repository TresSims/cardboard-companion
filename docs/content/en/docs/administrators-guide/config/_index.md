---
title: Config
description: How do I configure my CardboardCompanion?
categories: [Documentation]
weight: 4
---

There are currently 4 valid keys for cardboard companion config. Two are required.

## Required

```yaml
Token: <your bot token>
GuildID: <The guild to register your bot to>
```

## Optional

```yaml
PollSchedule: "0 9 * * 4" # CronTab formatted schedule to run the format poll.
PollChannel: "1357492784424292416" # Does nothing if PollSchedule is not defined.

```

{{% alert title="Warning" color="warning" %}}
PollChannel is *required* if PollSchedule is defined.
{{% /alert %}}
