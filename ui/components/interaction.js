export default function Interaction(props) {
  return (
    <div className="flex flex-col p-[10px] m-[10px] flex-grow rounded-sm bg-card-bg">
      <h2 className="text-fg-hilight text-center p-2">{props.title}</h2>
      {props.children}
    </div>
  )
}
