const fs = require('fs');
const readline = require('readline');

// distance = -(pressTime*pressTime) + (raceTime*pressTime)
// Shift that graph down by the record and find the roots of that equation
// 0 =  -(pressTime*pressTime) + (raceTime*pressTime) - record
//
// Any whole numbers between those roots will be winning numbers

async function readInputFile(inputFile) {

    const fileStream = fs.createReadStream(inputFile);
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let times = [];
    let distances = [];
    for await (const line of rl) {
        if (line.startsWith('Time:')) {
            times = line.match(/[\d]+/g);
        } else if (line.startsWith('Distance:')) {
            distances = line.match(/[\d]+/g);
        }
    }

    let marginProduct = 1;
    for (let i=0; i<times.length; i++) {
        let a = -1
        let b = times[i];
        let c = (0 - distances[i]);

        let roots = findRoots(a, b, c);
        let winStart =  Math.floor(roots[0] + 1);
        let winEnd = Math.ceil(roots[1] - 1);
        let m = winEnd - winStart + 1;

        console.log(`Wins begin at ${winStart} and end at ${winEnd}, for a margin of ${m}`);


        marginProduct = marginProduct * m;
    }
    console.log(`Margin product is ${marginProduct}`);

}

function findRoots(a, b, c) {

    let d = b * b - 4 * a * c;
    let sqrt_val = Math.sqrt(Math.abs(d));

    if (d <= 0) {
        console.log('Unable to solve it this way');
        process.exit(1);
    }

    let root1 = (-b + sqrt_val) / (2 * a);
    let root2 = (-b - sqrt_val) / (2 * a);

    return [root1, root2];
}

readInputFile('input.txt');
