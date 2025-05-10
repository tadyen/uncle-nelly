function Steps(){
  return (
    <div className="mx-auto">
      <ul className="steps">
        <li className="step step-primary">Select Base</li>
        <li className="step step-primary">Select Ingredients</li>
        <li className="step">Crackulate</li>
        <li className="step">Result</li>
      </ul>
    </div>
  )
}

export default function CookingSim(){
    return(<>
      <div className="mx-auto flex">
        <Steps/>
      </div>
    </>)
}
