'use client';

import { useState } from 'react';
import { ListBox, SortableItem } from '../list'

export default function Format() {
  const [items, setItems] = useState(["Commander", "Can't Play Today", "The Wheel"])

  return (
    <ListBox
      items={items}
      setItems={setItems}
    >
      {items.map(id => <SortableItem key={id} id={id}>{id}</SortableItem>)}
    </ListBox>
  )
}
