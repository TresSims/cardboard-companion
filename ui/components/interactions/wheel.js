'use client';

import { useState } from 'react';
import { ListBox, SortableItem } from '../list'

export default function Wheel() {
  const [items, setItems] = useState(["TwoHeaded Strix", "Commander", "Other Commander", "Third Commander"])

  return (
    <ListBox
      items={items}
      setItems={setItems}
    >
      {items.map(id => <SortableItem key={id} id={id}>{id}</SortableItem>)}
    </ListBox>
  )
}
