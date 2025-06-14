export default function colorInvert(hex: string): string {
    // expects a hex color string, e.g. #ffffff
    // returns a hex color string,
    const hexColorVal = hex.replace('#', '0x');
    const mask = "0xFFFFFF";
    const colorAsInt = parseInt(hexColorVal, 16);
    const maskAsInt = parseInt(mask, 16);
    const invertedAsInt = colorAsInt ^ maskAsInt;
    const invertedHex = invertedAsInt.toString(16).toUpperCase();
    const invertedHexPadded = invertedHex.padStart(6, '0');
    const invertedColorString = `#${invertedHexPadded}`;
    return invertedColorString
}
