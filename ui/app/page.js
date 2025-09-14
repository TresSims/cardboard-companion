import Interaction from '../components/interaction'
import Wheel from '../components/interactions/wheel'
import Format from '../components/interactions/format'

export default function Home() {
  return (
    <>
      <Interaction title="The Wheel"><Wheel /></Interaction>
      <Interaction title="Format Poll"><Format /></Interaction>
    </>
  );
}
