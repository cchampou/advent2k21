const fs = require("fs");
const path = require("path");
const readline = require('readline');
const events = require("events");

console.log('Application started, measures array initialization');
const measures = [];

const getSum3Last = (input, index) => index >= 2 ? input[index - 2] + input[index - 1] + input[index] : null;

(async () => {
  try {
    const rlInterface = readline.createInterface({
      input: fs.createReadStream(path.join(__dirname, 'input'))
    })
    rlInterface.on('line', (line) => {
      measures.push(Number(line));
    })
    await events.once(rlInterface, 'close');
    console.log(`File analysis complete, have ${measures.length} values`)
    console.log('Counting increases')
    const result = measures.reduce((acc, current, index) => {
      if (getSum3Last(measures, index - 1) && getSum3Last(measures, index) && getSum3Last(measures, index) > getSum3Last(measures, index - 1)) return acc + 1;
      return acc;
    }, 0);
    console.log(`Found ${result} increases`);
  } catch (e) {
    console.error('Fail to read input file')
}
})()
