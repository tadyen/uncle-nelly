import React from 'react';
// import { useAppContext } from '../AppContext';

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

function CookingSim(){
  // const appContext = useAppContext();
  // const un = appContext.uncleNelly;

  const [progress, setProgress] = React.useState<number>(1);
  // const [tables, setTables] = React.useState<any>(null);

  React.useEffect(() => {
    // Default progress to 1
    setProgress(1);

    // if (un) {
    //   setTables(un.get_tables());
    // }
    //
    // console.log(tables);

  },[]);

  return(<>
    <div className="mx-auto flex">
      <Steps {...{ progress }} />
    </div>
  </>)
}

export default CookingSim;
