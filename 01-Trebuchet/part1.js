const fs = require('fs');
const readline = require('readline');

async function parseInput() {

    const fileStream = fs.createReadStream('01-Trebuchet-input.txt');
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });
    let total = 0;
    for await (const line of rl) {
        const lineValue = numberFromLine(line);
        // console.log(`Line value: ${lineValue}`);
        total += lineValue;
    }
    console.log(`Total: ${total}`);

}

function numberFromLine(line) {

    let firstDigit = null;
    let lastDigit = null;
    for (let i = 0; i < line.length; i++) {
        const ch = line.charAt(i);
        if ((ch >= '0') && (ch <= '9')) {
            if (firstDigit == null) {
                firstDigit = ch;
            }
            lastDigit = ch;
        }
    }
    return parseInt(`${firstDigit}${lastDigit}`);
}

parseInput();
