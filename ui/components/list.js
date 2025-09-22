'use client';

import {
  SortableContext,
  verticalListSortingStrategy,
  arrayMove,
  useSortable
} from "@dnd-kit/sortable";
import { DndContext } from "@dnd-kit/core";
import { CSS } from '@dnd-kit/utilities'

export function ListBox(props) {
  function handleDragEnd(event) {
    const { active, over } = event;

    if (active.id !== over.id) {
      props.setItems((item) => {
        const oldIndex = item.indexOf(active.id)
        const newIndex = props.items.indexOf(over.id)

        return arrayMove(props.items, oldIndex, newIndex)
      })
    }
  }

  return (
    <div className="bg-list-bg flex-grow rounded-sm p-2">
      <DndContext onDragEnd={handleDragEnd}>
        <SortableContext items={props.items} strategy={verticalListSortingStrategy}>
          {props.children}
        </SortableContext>
      </DndContext>
    </div>
  )
}

export function SortableItem(props) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition
  } = useSortable({ id: props.id })

  const style = { transform: CSS.Transform.toString(transform), transition }

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      className="bg-list-item-bg text-list-item-fg p-[5px] m-[5px] rounded-sm cursor-pointer"
    >
      {props.children}
    </div >
  )
}
