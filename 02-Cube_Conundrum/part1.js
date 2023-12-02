const fs = require('fs');
const readline = require('readline');

const bag = {
    "red": 12,
    "green": 13,
    "blue": 14
}

async function readInputFile(inputFile) {

    const fileStream = fs.createReadStream(inputFile);
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });
    
    let total = 0;
    for await (const line of rl) {
        let game = parseGame(line);
        total += evaluateGame(game);
    }
    console.log(`Total: ${total}`);

}

function parseGame(gameString) {
    
    let str = gameString.split(':')

    let game = {
        "gameNumber": parseInt(str[0].substring(5)),
        "draws": []
    };
    
    drawResults = str[1].trim().split(';');
    drawResults.forEach((dr) => {
        let countsDict = parseDraw(dr);
        game["draws"].push(countsDict);
    });
    return game;
}

function parseDraw(drawResult){
    let counts = drawResult.trim().split(',');
    let countsDict = {};
    counts.forEach((cnt) => {
        c = cnt.trim().split(' ');
        countsDict[c[1]] = parseInt(c[0]);
    });
    return countsDict;
}

function evaluateGame(game) {

    let valid = true;
    game.draws.forEach((draw) => {
        for (const [key, value] of Object.entries(draw)) {
            if(value > bag[key]) {
                valid = false;
                break;
            }
        }
    });
    return valid ? game.gameNumber : 0;
}

/*
const testInput = [
    'Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green',
    'Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue',
    'Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red',
    'Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red',
    'Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green'
];

let total = 0;
testInput.forEach((gameString) => {
    let game = parseGame(gameString);
    total += evaluateGame(game);
});
console.log(`Total: ${total}`);
*/
readInputFile('input.txt')
