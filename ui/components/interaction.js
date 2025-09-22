export default function Interaction(props) {
  return (
    <div className="flex flex-col p-[10px] m-[10px] flex-grow rounded-sm bg-card-bg max-w-[50%]">
      <h2 className="text-fg-hilight text-center p-2 text-2xl font-black">{props.title}</h2>
      {props.children}
    </div>
  )
}
