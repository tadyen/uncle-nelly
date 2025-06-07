import React from 'react';
import { useAppContext } from '../AppContext';





// PLACEHOLDER ONLY
import breadImg from '../assets/baseIngredients/OGKush_Icon.png';
import riceImg from '../assets/baseIngredients/Cocaine_Icon.png';
import pastaImg from '../assets/baseIngredients/SourDiesel_Icon.png';
const BASE_IMAGES: Record<string, string> = {
  bread: breadImg,
  rice: riceImg,
  pasta: pastaImg,
};
const INGREDIENT_IMAGES: Record<string, string|undefined> = {
  egg: undefined,
  cheese: undefined,
  tomato: undefined,
  chicken: undefined,
};
const BASES = [
  { label: "Bread", value: "bread" },
  { label: "Rice", value: "rice" },
  { label: "Pasta", value: "pasta" },
];
const INGREDIENTS = [
  { label: "Egg", value: "egg" },
  { label: "Cheese", value: "cheese" },
  { label: "Tomato", value: "tomato" },
  { label: "Chicken", value: "chicken" },
];



// END PLACEHOLDERS


type StepProps = {
  progress: number;
  base: string;
  onStepClick: (step: number) => void;
};
function Steps(props: StepProps){
  //progress starts at 1
  const { progress, base, onStepClick } = props;
  const baseObj = BASES.find((b) => b.value === base);

  const steps = [
    {
      label: progress > 1 && baseObj
        ? (
          <span className="flex items-center gap-2">
            <IngredientIcon
              src={BASE_IMAGES[baseObj.value]}
              alt={baseObj.label}
            />
            {baseObj.label}
          </span>
        )
        : ( "Select Base" ),
    },
    { label: "Select Ingredients" },
    { label: "Crack-ulate" },
    { label: "Result" },
  ];

  return (
    <div className="mx-auto">
      <ul className="steps">
        {steps.map((step, idx) => (
          <li
            key={idx}
            className={
              (progress >= idx + 1 ? "step step-primary" : "step") +
              " cursor-pointer select-none"
            }
            onClick={() => onStepClick(idx + 1)}
            tabIndex={0}
            aria-label={`Go to step ${idx + 1}`}
          >
            {step.label}
          </li>
        ))}
      </ul>
    </div>
  );
}

// Helper for fallback placeholder (plain orange circle)
function IngredientIcon({ src, alt }: { src?: string; alt: string }) {
  const [imgError, setImgError] = React.useState(false);
  return src && !imgError ? (
    <span className="inline-block w-5 h-5 rounded-full bg-base-200 overflow-hidden">
      <img
        src={src}
        alt={alt}
        className="w-full h-full object-cover"
        onError={() => setImgError(true)}
      />
    </span>
  ) : (
    <span className="inline-block w-5 h-5 rounded-full bg-orange-400" />
  );
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
