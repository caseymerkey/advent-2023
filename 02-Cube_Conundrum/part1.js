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

readInputFile('input.txt');
