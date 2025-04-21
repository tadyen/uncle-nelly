import '../App.css'
import nelly_image from '../assets/uncle_nelly.png'

export default function Navbar(){
    return (<>
<div className="navbar bg-base-100 shadow-sm">
  <div className="navbar-start">
    <img src={nelly_image} className="max-h-24"/>
    <a className="btn btn-ghost text-xl">Uncle Nelly</a>
  </div>
  <div className="navbar-center hidden lg:flex">

    <div role="tablist" className="tabs tabs-border">
      <a role="tab" className="tab tab-active">Cooking Simulator</a>
      <a role="tab" className="tab">Recipe Reverse (tbd)</a>
      <a role="tab" className="tab">Optimiser (tbd)</a>
    </div>
  </div>
  <div className="navbar-end">
  </div>
</div>
    </>)
}
