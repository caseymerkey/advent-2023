const fs = require('fs');
const readline = require('readline');

const numberPattern = /\d+/g;
const gearPattern = /[\*]/g;

let above = null;
let center = null;
let below = null;
let currentCenterLine = -1;

async function readInputFile(inputFile) {

    const fileStream = fs.createReadStream(inputFile);
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let total = 0;
    for await (const line of rl) {
        advance(line);

        if (center) {
            total += evaluate();
        }
    }
    advance(null);
    total += evaluate();
    console.log(`Total: ${total}`);

}

function evaluate() {

    let match;
    let gearRatio = 0;

    while ((match = gearPattern.exec(center)) != null) {
        const gearIndex = match.index;
        const gears = [];
        console.log(`Searching for numbers near '*' found at index ${gearIndex} on line ${currentCenterLine}`);
        if (above) {
            let numbers = searchLineForNumbers(gearIndex, above);
            numbers.forEach((n) => {gears.push(n)});
        }
        if (below) {
            let numbers = searchLineForNumbers(gearIndex, below);
            numbers.forEach((n) => {gears.push(n)});
        }
        let numbers = searchLineForNumbers(gearIndex, center);
        numbers.forEach((n) => {gears.push(n)});

        console.log(`Found ${gears.length} numbers near '*' found at index ${gearIndex} on line ${currentCenterLine}`);
        if (gears.length === 2) {
            gearRatio += (gears[0] * gears[1]);
        }
    }

    return gearRatio;
}

function searchLineForNumbers(index, line) {
    const numbers = [];
    while ((match = numberPattern.exec(line)) != null) {
        const numberString = match[0];
        const numberStartIndex = match.index
        const numberEndIndex = numberStartIndex + numberString.length - 1;
        if (((index + 1) >= numberStartIndex) && ((index - 1) <= numberEndIndex)) {
            numbers.push(parseInt(numberString));
        }
    }
    return numbers;
}

function advance(line) {
    above = center;
    center = below;
    below = line;
    currentCenterLine++;
}

readInputFile('input.txt');
