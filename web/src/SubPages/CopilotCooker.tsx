import React from 'react';

// Import your actual ingredient images
import breadImg from '../assets/baseIngredients/OGKush_Icon.png';
import riceImg from '../assets/baseIngredients/Cocaine_Icon.webp';
import pastaImg from '../assets/baseIngredients/SourDiesel_Icon.png';


// Map ingredient/base values to their images
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

type StepProps = {
  progress: number;
  base: string;
  onStepClick: (step: number) => void;
};
function Steps(props: StepProps) {
  const { progress, base, onStepClick } = props;
  const baseObj = BASES.find((b) => b.value === base);

  const steps = [
    {
      label:
        progress > 1 && baseObj ? (
          <span className="flex items-center gap-2">
            <IngredientIcon
              src={BASE_IMAGES[baseObj.value]}
              alt={baseObj.label}
            />
            {baseObj.label}
          </span>
        ) : (
          "Select Base"
        ),
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

function CopilotCooker() {
  const [progress, setProgress] = React.useState<number>(1);

  // Step 1: Base
  const [base, setBase] = React.useState<string>("");

  // Step 2: Ingredients
  const [ingredient, setIngredient] = React.useState<string>("");
  const [ingredients, setIngredients] = React.useState<string[]>([]);

  // Drag state for reordering
  const [draggedIdx, setDraggedIdx] = React.useState<number | null>(null);

  // Step 3: Calculation (dummy)
  const [result, setResult] = React.useState<string>("");

  // Reset on mount
  React.useEffect(() => {
    setProgress(1);
    setBase("");
    setIngredients([]);
    setIngredient("");
    setResult("");
  }, []);

  // Step handlers
  const handleNext = () => {
    if (progress === 3) {
      // Fake calculation
      setResult(
        `You made ${base} with ${ingredients.length ? ingredients.join(", ") : "nothing"}!`
      );
    }
    setProgress((p) => Math.min(4, p + 1));
  };
  const handleBack = () => setProgress((p) => Math.max(1, p - 1));

  // Ingredient add
  const handleAddIngredient = () => {
    if (ingredient) {
      setIngredients([...ingredients, ingredient]);
      setIngredient("");
    }
  };
  const handleRemoveIngredient = (ing: string, idx: number) => {
    setIngredients(ingredients.filter((_, i) => i !== idx));
  };

  // Drag and drop handlers
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

  // Step jump logic
  const handleStepClick = (step: number) => {
    // Prevent jumping to step 2+ without a base
    if (step > 1 && !base) return;
    // Prevent jumping to step 3+ without at least one ingredient
    if (step > 2 && ingredients.length === 0) return;
    // When jumping to step 4, recalculate result
    if (step === 4) {
      setResult(
        `You made ${base} with ${ingredients.length ? ingredients.join(", ") : "nothing"}!`
      );
    }
    setProgress(step);
  };

  return (
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
            {BASES.map((b) => (
              <option key={b.value} value={b.value}>
                {b.label}
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
              onChange={e => {
                const selected = e.target.value;
                if (selected) {
                  setIngredients([...ingredients, selected]);
                  // Reset select to placeholder after adding
                  e.target.value = "";
                }
              }}
            >
              <option value="">Select ingredient</option>
              {INGREDIENTS.map((i) => (
                <option key={i.value} value={i.value}>
                  {i.label}
                </option>
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
                    rounded-full px-3 py-1
                    bg-info text-info-content
                    shadow
                  `}
                >
                  {/* Ingredient icon */}
                  <IngredientIcon
                    src={INGREDIENT_IMAGES[ing]}
                    alt={ing}
                  />
                  {/* Ingredient label */}
                  <span>
                    {INGREDIENTS.find((i) => i.value === ing)?.label || ing}
                  </span>
                  {/* Remove button */}
                  <button
                    type="button"
                    className="btn btn-xs btn-circle btn-ghost ml-1"
                    onClick={() => handleRemoveIngredient(ing, idx)}
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
            <span className="font-semibold">Base:</span> {BASES.find((b) => b.value === base)?.label}
          </div>
          <div>
            <span className="font-semibold">Ingredients:</span>{" "}
            {ingredients.length
              ? ingredients.map((ing) => INGREDIENTS.find((i) => i.value === ing)?.label).join(", ")
              : "None"}
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
          <div className="alert alert-success">{result}</div>
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
  );
}

export default CopilotCooker;
