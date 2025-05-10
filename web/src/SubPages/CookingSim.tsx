import React from 'react';


type StepProps = {
  progress: number;
}
function Steps(props: StepProps){
  //progress starts at 1
  const progress = props.progress;
  return (
    <div className="mx-auto">
      <ul className="steps">
        <li className={(progress>=1) ? "step step-primary" : "step"}>Select Base</li>
        <li className={(progress>=2) ? "step step-primary" : "step"}>Select Ingredients</li>
        <li className={(progress>=3) ? "step step-primary" : "step"}>Crack-ulate</li>
        <li className={(progress>=4) ? "step step-primary" : "step"}>Result</li>
      </ul>
    </div>
  )
}

export default function CookingSim(){
  const [progress, setProgress] = React.useState<number>(1);

  React.useEffect(() => {
    setProgress(3);
  },[]);

  return(<>
    <div className="mx-auto flex">
      <Steps {...{ progress }} />
    </div>
  </>)
}
