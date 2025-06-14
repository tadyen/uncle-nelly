import React from 'react';

import { useAppContext } from '../AppContext';
import { type UncleNelly, UncleNellyTables } from '../unclenelly_types';
import until from '../helpers/until';

type StepProps = {
  progress: number;
  base: string;
  numIngredients?: number;
  onStepClick: (step: number) => void;
};
function Steps(props: StepProps){
  const app = useAppContext();
  const [baseIngredients, setBaseIngredients] = React.useState<UncleNellyTables['base_ingredients'] | null>(null);

  //progress starts at 1
  const { progress, base, numIngredients, onStepClick } = props;
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
    {
      label: (progress > 2 && numIngredients )
        ? `${numIngredients} Ingredients Selected`
        : "Select Ingredients"
    },
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

function EntryListItem({name, src, onClick}: { name: string, src?: string, onClick?: () => void }) {
  return (
    <div className='flex items-center gap-2 cursor-pointer select-none' onClick={onClick}>
      <Icon src={src} alt={name} />
      <span>{name}</span>
    </div>
  );
}

function CookingSim(){
  const app = useAppContext();

  const [baseIngredients, setBaseIngredients] = React.useState<UncleNellyTables['base_ingredients'] | null>(null);
  const [mixIngredients, setMixIngredients] = React.useState<UncleNellyTables['mix_ingredients'] | null>(null);

  const [progress, setProgress] = React.useState<number>(1);
  const [base, setBase] = React.useState<string>("");
  // const [ingredient, setIngredient] = React.useState<string>("");
  const [ingredients, setIngredients] = React.useState<string[]>([]);
  const [draggedIdx, setDraggedIdx] = React.useState<number|null>(null);

  const [unelly, setUnelly] = React.useState<UncleNelly|null>(null);
  // const [product, setProduct] = React.useState<any|null>(null); // TODO Unelly product type

  const baseDropdownRef = React.useRef<HTMLDetailsElement>(null);

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
    // setProduct(null);
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
      <Steps progress={progress} base={base} numIngredients={(ingredients).length} onStepClick={handleStepClick} />
      {/* Step 1: Select Base */}
      {progress === 1 && (
        <form
          className="card bg-base-100 shadow p-6 flex flex-col gap-4"
          onSubmit={(e) => {
            e.preventDefault();
            if (base) handleNext();
          }}
        >
          <details open className="menu menu-dropdown w-52" ref={baseDropdownRef}>
            <summary className="btn">
              { base
                ? <EntryListItem name={base} src={baseIngredients ? baseIngredients[base]?.Icon : undefined}/>
                : "Select Base"
              }
              ▼▲
              </summary>
            <ul className="menu menu-dropdown dropdown-content bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm">
              {baseIngredients && Object.entries(baseIngredients).map(([k,_]) => (
                <li key={k}>
                  <EntryListItem
                    name={k}
                    src={baseIngredients ? baseIngredients[k]?.Icon : undefined}
                    onClick={()=>{
                      setBase(k);
                      baseDropdownRef.current && (baseDropdownRef.current.open = false);
                    }}/>
                </li>
              ))}
            </ul>
          </details>
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
          <div className="flex gap-2">
            <div className="dropdown">
              <div tabIndex={0} role="button" className="btn m-1 w-xl">Select Ingredients (order matters)</div>
              <ul tabIndex={0} className="dropdown-content menu bg-base-100/90 rounded-box z-1 p-2 shadow-sm max-h-64 min-w-2xl overflow-y-auto">
              {mixIngredients && Object.entries(mixIngredients).map(([k,_]) => (
                <li key={k}>
                  <EntryListItem
                    name={k}
                    src={mixIngredients ? mixIngredients[k]?.Icon : undefined}
                    onClick={()=>{
                      setIngredients([...ingredients, k]);
                    }}/>
                </li>
              ))}
              </ul>
            </div>

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
              onClick={()=>{
                if (!unelly) return;
                var res;
                res = unelly.init_job();
                if (res.error){
                  console.error("Error initialising unelly", res.error);
                }
                res = unelly.reset_product();
                if (res.error){
                  console.error("Error resetting product:", res.error);
                }
                res = unelly.set_product_base(base);
                if (res.error){
                  console.error("Error setting product base:", res.error);
                }
                res = unelly.cook_with(...ingredients);
                if (res.error){
                  console.error("Error cooking with ingredients:", res.error);
                }
                handleNext();
              }}
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
          <div className="alert alert-success">
            {(()=>{
              if (!unelly) return "Uncle Nelly is not initialized.";
              const res = unelly.product_info()
              if (res.error) {
                console.error("Error getting product:", res.error);
                return "Error";
              }
              const result = res.response as {
                Base: string;
                Effects: string[];
                MixHistory: string[];
                MixQueue: string[];
                Multiplier: number;
                Price: number;
              };

              return (
                <div className="card bg-base-100 shadow-lg p-6 max-w-md mx-auto">
                  <div className="card-body flex flex-col gap-4">
                    {/* Base Name */}
                    <div className="flex items-center gap-3">
                      <span className="badge badge-primary badge-lg px-4 py-2 text-lg">
                        {result.Base}
                      </span>
                    </div>

                    {/* Effects */}
                    <div>
                      <div className="font-semibold mb-1 text-base-content/80">Effects</div>
                      <div className="flex flex-wrap gap-2">
                        {result.Effects.map((effect) => (
                          <span key={effect} className="badge badge-info badge-outline">
                            {effect}
                          </span>
                        ))}
                      </div>
                    </div>

                    {/* Mix History */}
                    <div>
                      <div className="font-semibold mb-1 text-base-content/80">Mix History</div>
                      <div className="flex flex-wrap gap-2">
                        {result.MixHistory.map((item) => (
                          <span key={item} className="badge badge-secondary badge-outline">
                            {item}
                          </span>
                        ))}
                      </div>
                    </div>

                    {/* Multiplier & Price */}
                    <div className="flex gap-4 mt-2">
                      <div className="flex flex-col items-center">
                        <span className="text-xs text-base-content/60">Multiplier</span>
                        <span className="font-bold text-lg text-success">{result.Multiplier}×</span>
                      </div>
                      <div className="flex flex-col items-center">
                        <span className="text-xs text-base-content/60">Price</span>
                        <span className="font-bold text-lg text-warning">${result.Price}</span>
                      </div>
                    </div>
                  </div>
                </div>
              )
            })()}

          </div>
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
