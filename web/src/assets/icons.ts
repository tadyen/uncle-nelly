// base ingredients
import Blank from "./react.svg";
import OGKush from "./baseIngredients/OGKush_Icon.png";
import SourDiesel from "./baseIngredients/SourDiesel_Icon.png";
import GreenCrack from "./baseIngredients/GreenCrack_Icon.png";
import GranddaddyPurple from "./baseIngredients/GranddaddyPurple_Icon.png";
import Meth from "./baseIngredients/Meth_Icon.png";
import Cocaine from "./baseIngredients/Cocaine_Icon.png";


// mix ingredients
import Addy from "./mixIngredients/Addy_Icon.png";
import Banana from "./mixIngredients/Banana_Icon.png";
import Battery from "./mixIngredients/Battery_Icon.png";
import Chili from "./mixIngredients/Chili_Icon.png";
import Cuke from "./mixIngredients/Cuke_Icon.png";
import Donut from "./mixIngredients/Donut_Icon.png";
import EnergyDrink from "./mixIngredients/Energy_Drink_Icon.png";
import FluMedicine from "./mixIngredients/Flu_Medicine_Icon.png";
import Gasoline from "./mixIngredients/Gasoline_Icon.png";
import HorseSemen from "./mixIngredients/Horse_Semen_Icon.png";
import Iodine from "./mixIngredients/Iodine_Icon.png";
import MegaBean from "./mixIngredients/Mega_Bean_Icon.png";
import MotorOil from "./mixIngredients/Motor_Oil_Icon.png";
import MouthWash from "./mixIngredients/Mouth_Wash_Icon.png";
import Paracetamol from "./mixIngredients/Paracetamol_Icon.png";
import Viagor from "./mixIngredients/Viagor_Icon.png";

// ENSURE keys match those from the go wasm

export const baseIngredientsIcons = {
    "Blank": Blank,
    "OG Kush": OGKush,
    "Sour Diesel": SourDiesel,
    "Green Crack": GreenCrack,
    "Granddaddy Purple": GranddaddyPurple,
    "Meth": Meth,
    "Cocaine": Cocaine,
} as {[key: string]: string};

export const mixIngredientsIcons = {
    "Addy": Addy,
    "Banana": Banana,
    "Battery": Battery,
    "Chili": Chili,
    "Cuke": Cuke,
    "Donut": Donut,
    "Energy Drink": EnergyDrink,
    "Flu Medicine": FluMedicine,
    "Gasoline": Gasoline,
    "Horse Semen": HorseSemen,
    "Iodine": Iodine,
    "Mega Bean": MegaBean,
    "Motor Oil": MotorOil,
    "Mouth Wash": MouthWash,
    "Paracetamol": Paracetamol,
    "Viagor": Viagor,
} as {[key: string]: string};

// relative path strings instead
/*
export const baseIngredientsIcons = {
    "Blank": "assets/react.svg",
    "OGKush": "assets/baseIngredientes/OGKush_Icon.png",
    "SourDiesel": "assets/baseIngredientes/SourDiesel_Icon.png",
    "GreenCrack": "assets/baseIngredientes/GreenCrack_Icon.png",
    "GranddaddyPurple": "assets/baseIngredientes/GranddaddyPurple_Icon.png",
    "Meth": "assets/baseIngredientes/Meth_Icon.png",
    "Cocaine": "assets/baseIngredients/Cocaine_Icon.png",
} as {[key: string]: string};

export const mixIngredientsIcons = {
    "Addy": "assets/mixIngredients/Addy_Icon.png",
    "Banana": "assets/mixIngredients/Banana_Icon.png",
    "Battery": "assets/mixIngredients/Battery_Icon.png",
    "Chili": "assets/mixIngredients/Chili_Icon.png",
    "Cuke": "assets/mixIngredients/Cuke_Icon.png",
    "Donut": "assets/mixIngredients/Donut_Icon.png",
    "Energy Drink": "assets/mixIngredients/Energy_Drink_Icon.png",
    "Flu Medicine": "assets/mixIngredients/Flu_Medicine_Icon.png",
    "Gasoline": "assets/mixIngredients/Gasoline_Icon.png",
    "Horse Semen": "assets/mixIngredients/Horse_Semen_Icon.png",
    "Iodine": "assets/mixIngredients/Iodine_Icon.png",
    "Mega Bean": "assets/mixIngredients/Mega_Bean_Icon.png",
    "Motor Oil": "assets/mixIngredients/Motor_Oil_Icon.png",
    "Mouth Wash": "assets/mixIngredients/Mouth_Wash_Icon.png",
    "Paracetamol": "assets/mixIngredients/Paracetamol_Icon.png",
    "Viagor": "assets/mixIngredients/Viagor_Icon.png",
} as {[key: string]: string};
*/

export const effectsColors = {
    "None": "#FFFFFF",
    "Anti-gravity": "#0051ff",
    "Athletic": "#00a2ff",
    "Balding": "#ffee00",
    "Bright-Eyed": "#00ffff",
    "Calming": "#fff09c",
    "Calorie-Dense": "#e100ff",
    "Cyclopean": "#ffd153",
    "Disorienting": "#ff6a65",
    "Electrifying": "#00aeff",
    "Energizing": "#1eff00",
    "Euphoric": "#ffc14f",
    "Explosive": "#ff0000",
    "Focused": "#00fff2",
    "Foggy": "#b9b9b9",
    "Gingeritis": "#ff7300",
    "Glowing": "#00ff62",
    "Jennerising": "#ff00ff",
    "Laxative": "#7F0000",
    "Long faced": "#ffd86d",
    "Munchies": "#e90000",
    "Paranoia": "#c54545",
    "Refreshing": "#88ff7d",
    "Schizophrenic": "#0000a0",
    "Sedating": "#8b5ada",
    "Shrinking": "#6cc9e0",
    "Siezure-Inducing": "#FFBF00",
    "Slippery": "#7edad5",
    "Smelly": "#368b3b",
    "Sneaky": "#a3a3a3",
    "Spicy": "#ff0000",
    "Thought-Provoking": "#ffa6ff",
    "Toxic": "#1a8100",
    "Tropic Thunder": "#ff4800",
    "Zombifying": "#328800",
} as {[key: string]: string};
