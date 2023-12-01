const fs = require('fs');
const readline = require('readline');

const numberPatterns = [
    /zero/g,
    /one/g,
    /two/g,
    /three/g,
    /four/g,
    /five/g,
    /six/g,
    /seven/g,
    /eight/g,
    /nine/g
];

Number.prototype.pad = function (size) {
    var s = String(this);
    while (s.length < (size || 2)) { s = "0" + s; }
    return s;
}

async function parseInput() {

    // Not sure why, but VSCode's debugger didn't like the relative path
    const fileStream = fs.createReadStream(`${__dirname}/01-Trebuchet-input.txt`);

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });
    let total = 0;
    let lineNumber = 0;
    for await (const line of rl) {
        lineNumber++;
        const lineValue = numberFromLine(line);
        console.log(`Line ${lineNumber.pad(4)} value: ${lineValue}`);
        total += lineValue;
    }
    console.log(`Total: ${total}`);

}

function numberFromLine(input) {
    let firstDigit = null;
    let firstIndex = 9999999;
    let lastDigit = null;
    let lastIndex = -1;

    let re = /\d/g;
    let match;
    while ((match = re.exec(input)) != null) {

        let index = match.index;
        let digit = match[0];
        if (index < firstIndex) {
            firstDigit = digit;
            firstIndex = index;
        }
        if (index > lastIndex) {
            lastDigit = digit;
            lastIndex = index;
        }
    }

    for (let n = 1; n < numberPatterns.length; n++) {
        let re = numberPatterns[n];
        while ((match = re.exec(input)) != null) {
            let index = match.index;
            if (index < firstIndex) {
                firstDigit = n;
                firstIndex = index;
            }
            if (index > lastIndex) {
                lastDigit = n;
                lastIndex = index;
            }
        }
    }
    return parseInt(`${firstDigit}${lastDigit}`);
}

parseInput();
