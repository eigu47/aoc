import fs from "fs";

const [seeds, ...maps] = fs
  .readFileSync("day5", "utf8")
  .split("\r\n\r\n")
  .map((line) =>
    line
      .split(":")[1]
      .split("\r\n")
      .filter(Boolean)
      .map((str) => str.match(/\d+/g)?.map(Number) ?? [])
  );

const closestSeed = seeds[0].reduce((closest, start, idx, arr) => {
  if (idx % 2 === 0) {
    const range = arr[idx + 1];

    for (let i = start; i < start + range; ) {
      let minRange = range;

      const destination = maps.reduce((acc, map) => {
        for (const [destination, origin, range] of map) {
          if (acc >= origin && acc < origin + range) {
            minRange = Math.min(minRange, origin + range - acc);
            acc = acc - origin + destination;
            break;
          }
        }
        return acc;
      }, i);

      if (destination < closest) {
        closest = destination;
      }

      i += minRange;
    }
  }
  return closest;
}, Infinity);

console.log(closestSeed);
