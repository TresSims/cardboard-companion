---
title: Interactions
description: The basic building blocks of a CardboardCompanion
categories: [Documentation]
weight: 6
---

Interactions are the component blocks that CardboardCompanion uses to communicate with
the server. Basically, it's how you can expect the [/format](/docs/users-guide/commands#format)
command and the poll [cron job](/docs/users-guide/cron-jobs/) will do the exact same thing.

There are two interaction components right now.

## Format Poll

The format poll creates an 8 hour long poll with the following options:

- 60-Card Format
- EDH
- The Wheel
- Don't Care/Suggestion
- Not Playing

It also contains the text

```
Another Thursday? Another Poll!
```

## The Wheel

The wheel interaction randomly picks a game mode from the following list:

- "Planechase",
- "Share the Spoils",
- "Quiz Commander",
- "Truly Random",
- "Monochrome Matchup",
- "Mechanic Mashup",
- "Any # of Fools",
- "Bidget Battle",
- "Keywork Klash",
- "Partner Up",
- "Guild Wars",
- "Oathbreaker",
- "Tribal Throwdown",
- "Color Swap",
- "Pauper EDH",
- "Oldies but Goodies",
- "Framed Fight",
- "Alt Win Con",
- "Set Showdown",
- "Teeny Weenies",
- "Big Chungies",
- "Planeswalekr Party",
- "Precon Party",
- "Usurper/Kingdoms",
- "Archenemy",
- "Deck Swap",
- "Gifts Only",
- "Booster Pack Madness",
- "Bounties",
- "Two-Headed Giant",
