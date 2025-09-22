'use client';

import { useState } from 'react';
import { ListBox, SortableItem } from '../list'
import CronBox from '../cronBox'

const formatList = [
  "60 Card Format",
  "EDH",
  "The Wheel",
  "Other/Don't Care",
  "Can't Play Today"
]

export default function Format() {
  const [items, setItems] = useState(formatList)

  return (
    <div>
      <h3 className="m-2">
        Title: Another Thursday? Another Poll!
      </h3>
      <p className="m-2">
        It&apos;s game time folks! Please vote in the poll, it&apos;s how we know who&apos;s coming to play! See you at `epoch + 9h30m`
      </p>
      <ListBox
        items={items}
        setItems={setItems}
      >
        {items.map(id => <SortableItem key={id} id={id}>{id}</SortableItem>)}
      </ListBox>
      <CronBox />
    </div>
  )
}
