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
  },[app.UNtables])

  const steps = [
    {
      label: progress > 1 && baseIngredient?.Name
        ? (
          <div className="flex flex-col items-center gap-2 mt-2">
            <Icon
              src={baseIngredient.Icon}
              alt={baseIngredient.Name}
            />
            <div>{baseIngredient.Name}</div>
          </div>
        )
        : ( "Select Base" ),
    },
    { label: "Select Ingredients" },
    { label: "Crack-ulate" },
    { label: "Result" },
  ];

  return (
    <div className="flex flex-col items-center w-4xl h-30 mx-auto">
        <ul className="steps w-full">
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
    <div className="inline-block w-10 h-10 rounded-full bg-transparent overflow-hidden">
      <img
        src={src}
        alt={alt}
        className="w-full h-full object-cover"
        onError={() => setImgError(true)}
      />
    </div>
  ) : (
    // Fallback icon if absolutely no src provided / image err
    <div className="inline-block w-10 h-10 rounded-full bg-info" />
  );
}


function CookingSim(){
  const app = useAppContext();

  const [baseIngredients, setBaseIngredients] = React.useState<UncleNellyTables['base_ingredients'] | null>(null);
  const [mixIngredients, setMixIngredients] = React.useState<UncleNellyTables['mix_ingredients'] | null>(null);

  const [progress, setProgress] = React.useState<number>(1);
  const [base, setBase] = React.useState<string>("");
  const [ingredient, setIngredient] = React.useState<string>("");
  const [ingredients, setIngredients] = React.useState<string[]>([]);
  const [draggedIdx, setDraggedIdx] = React.useState<number|null>(null);

  const [unelly, setUnelly] = React.useState<UncleNelly|null>(null);
  const [product, setProduct] = React.useState<any|null>(null); // TODO Unelly product type

  const getUN = async () => {
    await until(() => app.uncleNelly != null);
    return app.uncleNelly;
  }
    React.useEffect(() => {
      (async () => {
        await until(() => app.UNtables != null);
        const _baseIngredients: UncleNellyTables['base_ingredients'] = app.UNtables?.base_ingredients;
        const _mixIngredients: UncleNellyTables['mix_ingredients'] = app.UNtables?.mix_ingredients;
        setBaseIngredients(_baseIngredients);
        setMixIngredients(_mixIngredients);
      })()
    },[])


  // Initialise on mount
  React.useEffect(() => {
    const initUnelly = async () => {
      const unelly = await getUN();
      setUnelly(unelly);
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
    <div className="mx-auto flex flex-col gap-8 max-w-4xl">
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
            defaultValue={"Select base"}
            value={base}
            onChange={(e) => setBase(e.target.value)}
            required
          >
            <option value="" disabled={true}>Select base</option>
            {baseIngredients && Object.entries(baseIngredients).map(([k,v]) => (
              <option key={k} value={k}>{k}</option>
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


      {/* Step 2: Select Ingredients */}
      {progress === 2 && (
        <form
          className="card bg-base-100 shadow p-6 flex flex-col gap-4"
          onSubmit={(e) => {
            e.preventDefault();
            handleNext();
          }}
        >
          <label className="label">Add ingredients (order matters):</label>
          <div className="flex gap-2">
            <select
              className="select select-accent"
              value=""
              defaultValue="Select ingredient"
              onChange={e => {
                const selected = e.target.value;
                if (selected) {
                  setIngredients([...ingredients, selected]);
                  // Reset select to placeholder after adding
                  e.target.value = "";
                }
              }}
            >
              <option value="" disabled={true}>Select ingredient</option>
              {mixIngredients && Object.entries(mixIngredients).map(([k,v]) => (
                <option key={k} value={k}>{k}</option>
              ))}
            </select>
            <button
              type="button"
              className="btn btn-outline btn-error"
              onClick={() => setIngredients([])}
              disabled={ingredients.length === 0}
            >
              Clear
            </button>
          </div>
          <div className="flex flex-wrap gap-2 py-2">
            {ingredients.map((ing, idx) => (
              <div
                key={idx}
                className={`flex items-center gap-1 cursor-move ${
                  draggedIdx === idx ? "opacity-50" : ""
                }`}
                draggable
                onDragStart={() => handleDragStart(idx)}
                onDragOver={e => handleDragOver(e, idx)}
                onDragEnd={handleDragEnd}
                onDrop={handleDragEnd}
              >
                <span
                  className={`
                    inline-flex items-center gap-1
                    rounded-lg px-3 py-1
                    border-info bg-secondary text-secondary-content
                    shadow
                  `}
                >
                  {/* Ingredient icon */}
                  <Icon
                    src={mixIngredients ? mixIngredients[ing].Icon : ''}
                    alt={ing}
                  />
                  { mixIngredients ? mixIngredients[ing].Name : '' }
                  {/* Ingredient label */}
                  {/* Remove button */}
                  <button
                    type="button"
                    className="btn btn-xs btn-circle btn-ghost ml-1"
                    onClick={() => handleRemoveIngredient(idx)}
                    aria-label="Remove"
                  >
                    ×
                  </button>
                </span>
                {idx < ingredients.length - 1 && (
                  <span className="text-xl text-base-content">→</span>
                )}
              </div>
            ))}
          </div>
          <div className="flex gap-2 mt-4">
            <button
              type="button"
              className="btn"
              onClick={handleBack}
            >
              Back
            </button>
            <button
              type="submit"
              className="btn btn-primary"
            >
              Next
            </button>
          </div>
        </form>
      )}

      {/* Step 3: Crack-ulate */}
      {progress === 3 && (
        <div className="card bg-base-100 shadow p-6 flex flex-col gap-4">
          <div className="text-lg font-bold">Ready to crack-ulate?</div>
          <div>
            <span className="font-semibold">Base: </span>
            { baseIngredients && baseIngredients[base].Name }
          </div>
          <div>
            <span className="font-semibold">Ingredients: </span>{" "}
            { (ingredients.length && mixIngredients)
              ? ingredients.map((ing) => mixIngredients[ing].Name).join(", ")
              : "None"
            }
          </div>
          <div className="flex gap-2 mt-4">
            <button
              type="button"
              className="btn"
              onClick={handleBack}
            >
              Back
            </button>
            <button
              type="button"
              className="btn btn-primary"
              onClick={handleNext}
            >
              Crack-ulate!
            </button>
          </div>
        </div>
      )}


      {/* Step 4: Result */}
      {progress === 4 && (
        <div className="card bg-base-100 shadow p-6 flex flex-col gap-4">
          <div className="text-lg font-bold">Result</div>
          <div className="alert alert-success">{"Result value"}</div>
          <div className="flex gap-2 mt-4">
            <button
              type="button"
              className="btn"
              onClick={() => setProgress(1)}
            >
              Start Over
            </button>
          </div>
        </div>
      )}


    </div>
  </>)
}

export default CookingSim;
