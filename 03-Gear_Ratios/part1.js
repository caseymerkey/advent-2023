const fs = require('fs');
const readline = require('readline');

const numberPattern = /\d+/g;
const symbolPattern = /[^\.\d]/;

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
    let lineSum = 0;
    let match;
    while((match = numberPattern.exec(center)) != null) {
        const numberString = match[0];
        console.log(`Evaluating ${numberString} on line ${currentCenterLine}`);
        let alreadyPassed = false;
        searchStart = Math.max(0, match.index - 1);
        searchEnd = Math.min((center.length - 1), (match.index + numberString.length + 1));
        if(above) {
            let testArea = above.substring(searchStart, searchEnd);
            if (symbolPattern.test(testArea)) {
                lineSum += parseInt(numberString);
                alreadyPassed = true;
            }
        }
        if(below && !alreadyPassed) {
            let testArea = below.substring(searchStart, searchEnd);
            if (symbolPattern.test(testArea)) {
                lineSum += parseInt(numberString);
                alreadyPassed = true;
            }
        }
        if (!alreadyPassed) {
            let testArea = '';
            if (match.index > 0) {
                testArea += center.substring(match.index - 1, match.index);
            }
            if ((match.index + numberString.length) < center.length) {
                testArea += center.substring((match.index + numberString.length ), (match.index + numberString.length + 1));
            }
            if (symbolPattern.test(testArea)) {
                lineSum += parseInt(numberString);
            } else {
                console.log(`Number ${numberString} on line ${currentCenterLine} does not pass`);
            }
        }
    }
    return lineSum;
}

function advance(line) {
    above = center;
    center = below;
    below = line;
    currentCenterLine++;
}

readInputFile('input.txt');
