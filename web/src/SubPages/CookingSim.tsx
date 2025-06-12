import React from 'react';

import { useAppContext } from '../AppContext';
import { type UncleNelly, UncleNellyTables } from '../unclenelly_types';
import until from '../helpers/until';

type StepProps = {
  progress: number;
  base: string;
  onStepClick: (step: number) => void;
};
function Steps(props: StepProps){
  const app = useAppContext();
  const [baseIngredients, setBaseIngredients] = React.useState<UncleNellyTables['base_ingredients'] | null>(null);

  //progress starts at 1
  const { progress, base, onStepClick } = props;
  const baseIngredient = baseIngredients ? baseIngredients[base] : undefined;

  React.useEffect(()=>{
    const _baseIngredients: UncleNellyTables['base_ingredients'] = app.UNtables?.base_ingredients;
    setBaseIngredients(_baseIngredients);
    console.log(baseIngredients);
  },[app.UNtables])


  const steps = [
    {
      label: progress > 1 && baseIngredient?.Name
        ? (
          <div className="flex items-center gap-2">
            <Icon
              src={baseIngredient.Icon}
              alt={baseIngredient.Name}
            />
            {baseIngredient.Name}
          </div>
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

function Icon({ src, alt }: { src?: string; alt: string }){
  const [imgError, setImgError] = React.useState(false);
  return src && !imgError ? (
    <div className="inline-block w-5 h-5 rounded-full bg-base-200 overflow-hidden">
      <img
        src={src}
        alt={alt}
        className="w-full h-full object-cover"
        onError={() => setImgError(true)}
      />
    </div>
  ) : (
    // Fallback icon if absolutely no src provided / image err
    <div className="inline-block w-5 h-5 rounded-full bg-orange-400" />
  );
}


function CookingSim(){
  const app = useAppContext();

  // const baseIngredients: UncleNellyTables['base_ingredients'] = appContext.UNtables?.base_ingredients;
  // const mixIngredients: UncleNellyTables['mix_ingredients'] = appContext.UNtables?.mix_ingredients;
  // const effects: UncleNellyTables['effects'] = appContext.UNtables?.effects;
  const [baseIngredients, setBaseIngredients] = React.useState<UncleNellyTables['base_ingredients'] | null>(null);

  const [progress, setProgress] = React.useState<number>(1);
  const [base, setBase] = React.useState<string>("");
  const [ingredient, setIngredient] = React.useState<string>("");
  const [ingredients, setIngredients] = React.useState<string[]>([]);
  const [draggedIdx, setDraggedIdx] = React.useState<number|null>(null);

  const [unelly, setUnelly] = React.useState<UncleNelly|null>(null);
  const [product, setProduct] = React.useState<any|null>(null); // TODO Unelly product type

  const newProduct = async () => {
    await until(() => app.UNellyLoader != null)
    const unellyLoader = app.UNellyLoader;
    if (unellyLoader) {
      return unellyLoader();
    }
    throw new Error("Uncle Nelly Loader is not initialized");
  }
    React.useEffect(() => {
      (async () => {
        await until(() => app.UNtables != null);
        const _baseIngredients: UncleNellyTables['base_ingredients'] = app.UNtables?.base_ingredients;
        setBaseIngredients(_baseIngredients);
      })()
    },[])


  // Initialise on mount
  React.useEffect(() => {
    const initUnelly = async () => {
      const un = await newProduct()
      setUnelly(un);
      un.init_job();
    }
    initUnelly();
    setProduct(null);
    setProgress(1);
    setBase("");
    setIngredients([]);
  },[]);

  // Step handlers
  const handleNext = () => {
    setProgress((p) => Math.min(4, p + 1));
  }
  const handleBack = () => setProgress((p) => Math.max(1, p - 1));

  // jump-to-step handler
  const handleStepClick = (step: number) => {
    // Prevent jumping to step 2+ without a base
    if (step > 1 && !base) return;
    // Prevent jumping to step 3+ without at least one ingredient
    if (step > 2 && ingredients.length === 0) return;
    // Prevent jumping to step 4 without a product
    //
    setProgress(step);
  }

  // ingredient add/remove/draggable handlers
  // add ingredient does not need a handler
  const handleRemoveIngredient = (idx: number) => {
    setIngredients(ingredients.filter((_, i) => i !== idx));
  }
  const handleDragStart = (idx: number) => setDraggedIdx(idx);
  const handleDragOver = (e: React.DragEvent<HTMLDivElement>, idx: number) => {
    e.preventDefault();
    if (draggedIdx === null || draggedIdx === idx) return;
    const newIngredients = [...ingredients];
    const [removed] = newIngredients.splice(draggedIdx, 1);
    newIngredients.splice(idx, 0, removed);
    setIngredients(newIngredients);
    setDraggedIdx(idx);
  };
  const handleDragEnd = () => setDraggedIdx(null);


  return(<>
    <div className="mx-auto flex flex-col gap-8 max-w-xl">
      <Steps progress={progress} base={base} onStepClick={handleStepClick} />

      {/* Step 1: Select Base */}
      {progress === 1 && (
        <form
          className="card bg-base-100 shadow p-6 flex flex-col gap-4"
          onSubmit={(e) => {
            e.preventDefault();
            if (base) handleNext();
          }}
        >
          <label className="label">Choose a base:</label>
          <select
            className="select select-primary"
            value={base}
            onChange={(e) => setBase(e.target.value)}
            required
          >
            <option value="">Select base</option>
            {baseIngredients && Object.entries(baseIngredients).map(([k,v]) => (
              <option key={k} value={k}>
                <Icon
                  src={v.Icon}
                  alt={k}
                />
                {k}
              </option>
            ))}
          </select>
          <div className="flex gap-2 mt-4">
            <button
              type="submit"
              className="btn btn-primary"
              disabled={!base}
            >
              Next
            </button>
          </div>
        </form>
      )}




    </div>
  </>)
}

export default CookingSim;
