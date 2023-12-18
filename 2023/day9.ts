import fs from "fs";

const input = fs
  .readFileSync("day9", "utf8")
  .split("\r\n")
  .map((line) => line.split(" ").map(Number))
  .reduce((acc, nums) => {
    const diffs = [nums];

    while (!diffs.at(-1)!.every((n) => n === 0)) {
      const cur = diffs.at(-1)!;
      const diff = cur.slice(0, -1).map((num, i) => cur[i + 1] - num);

      diffs.push(diff);
    }

    const extrapolated = diffs.reverse().reduce((acc, cur) => cur[0] - acc, 0);

    return acc + extrapolated;
  }, 0);

console.log(input);
