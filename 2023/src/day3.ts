import fs from "fs";

const gears = new Map<string, number[]>();

(fs.readFileSync("day3", "utf8") as string)
  .split("\n")
  .forEach((line, idx, arr) =>
    [...line.trim().matchAll(/\d+/g)].forEach((num) => {
      if (num.index === undefined) return;
      const adjacent: [idx: number, i: number][] = [];

      for (let i = num.index - 1; i < num.index + num[0].length + 1; i++) {
        adjacent.push([idx - 1, i]);
        adjacent.push([idx + 1, i]);
      }
      adjacent.push([idx, num.index - 1]);
      adjacent.push([idx, num.index + num[0].length]);

      adjacent
        .filter(([idx, i]) => arr[idx]?.[i] === "*")
        .forEach(([idx, i]) => {
          const key = `${idx},${i}`;
          const gearValue = gears.get(key) ?? [];
          gearValue.push(+num[0]);
          gears.set(key, gearValue);
        });
    })
  );

console.log(
  [...gears.values()].reduce((acc, cur) => {
    return cur.length === 2 ? acc + cur[0] * cur[1] : acc;
  }, 0)
);
