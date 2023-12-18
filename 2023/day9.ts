import fs from "fs";

const input = fs
  .readFileSync("day9", "utf8")
  .split("\r\n")
  .map((line) => line.split(" ")?.map(Number) ?? [])
  .reduce((acc, nums) => {
    const diffs: number[][] = [nums];

    while (!diffs.at(-1)?.every((n) => n === 0)) {
      const cur = diffs.at(-1)!;
      const diff = cur.slice(0, -1).map((num, i) => cur[i + 1] - num);

      diffs.push(diff);
    }

    return acc + diffs.reduce((acc, cur) => acc + (cur.at(-1) ?? 0), 0);
  }, 0);

console.log(input);
