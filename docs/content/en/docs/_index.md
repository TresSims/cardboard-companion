---
title: Documentation
linkTitle: Docs
menu: {main: {weight: 20}}
---

The CardboardCompanion in it 0.1.0 state has few features, but huge potential!

Right now it has two main features types, commands and cron jobs.

## Commands

Right now, there are three commands:

- /format : Send a poll
- /wheel : Pick a format from a pre-defined list
- /ping : A basic health check

## Cron Jobs

The only cron job available right now is the poll. It can be scheduled using
unix crontab formatting.

e.g.

`0 9 * * 4` = Send a poll at 9 am every Thursday
