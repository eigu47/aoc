import fs from "fs";

const input = (fs.readFileSync("day3", "utf8") as string)
  .split("\n")
  .map((line, idx, arr) =>
    [...line.trim().matchAll(/(\d+)/g)].map((num) => {
      if (num.index === undefined) return;
      const adjacent: [idx: number, i: number][] = [];

      for (let i = num.index - 1; i < num.index + num[0].length + 1; i++) {
        adjacent.push([idx - 1, i]);
        adjacent.push([idx + 1, i]);
      }
      adjacent.push([idx, num.index - 1]);
      adjacent.push([idx, num.index + num[0].length + 1]);

      const asterisk = adjacent.filter(([idx, i]) => arr[idx]?.[i] === "*");
      if (asterisk) {
        return {
          asterisk,
          number: +num[0],
        };
      }
    })
  )
  .flat()
  .reduce((acc: Map<string, number[]>, cur) => {
    cur?.asterisk.forEach(([idx, i]) => {
      const key = `${idx},${i}`;
      acc.set(key, [...(acc.get(key) ?? []), cur.number]);
    });

    return acc;
  }, new Map());

console.log(input);
console.log(
  [...input.values()].reduce((acc, cur) => {
    if (cur.length === 2) {
      const gear = cur[0] * cur[1];
      return acc + gear;
    }
    return acc;
  }, 0)
);
