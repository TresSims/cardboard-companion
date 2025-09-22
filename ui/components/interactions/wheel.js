'use client';

import { useState } from 'react';
import { ListBox, SortableItem } from '../list'

const wheelList = [
  "Planechase",
  "Share the Spoils",
  "Quiz Commander",
  "Truly Random",
  "Monochrome Matchup",
  "Mechanic Mashup",
  "Any # of Fools",
  "Bidget Battle",
  "Keywork Klash",
  "Partner Up",
  "Guild Wars",
  "Oathbreaker",
  "Tribal Throwdown",
  "Color Swap",
  "Pauper EDH",
  "Oldies but Goodies",
  "Framed Fight",
  "Alt Win Con",
  "Set Showdown",
  "Teeny Weenies",
  "Big Chungies",
  "Planeswalekr Party",
  "Precon Party",
  "Usurper/Kingdoms",
  "Archenemy",
  "Deck Swap",
  "Gifts Only",
  "Booster Pack Madness",
  "Bounties",
  "Two-Headed Giant",
]

export default function Wheel() {
  const [items, setItems] = useState(wheelList)

  return (
    <ListBox
      items={items}
      setItems={setItems}
    >
      {items.map(id => <SortableItem key={id} id={id}>{id}</SortableItem>)}
    </ListBox>
  )
}
