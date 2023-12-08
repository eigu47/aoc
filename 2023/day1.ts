const fs = require("fs");

const map = {
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight: "8",
  nine: "9",
};

const input = (fs.readFileSync("day1-1", "utf8") as string)
  .split("\n")
  .map((line) => {
    const first = line
      .trim()
      .split("")
      .find((c) => !isNaN(+c));

    const last = line
      .trim()
      .split("")
      .findLast((c) => !isNaN(+c));

    return (first ?? "") + (last ?? "");
  })
  .reduce((acc, cur) => acc + +cur, 0);

console.log(input);
