import fs from "fs";

const [seed, ...maps] = fs
  .readFileSync("day5", "utf8")
  .split("\r\n\r\n")
  .map((line) =>
    line
      .split(":")[1]
      .split("\r\n")
      .filter(Boolean)
      .map((str) => str.match(/\d+/g)?.map(Number) ?? [])
  );

const seedRanges = seed[0]
  .reduce((acc, num, idx, arr) => {
    if (idx % 2 === 0) {
      acc.push([num, arr[idx + 1]]);
    }
    return acc;
  }, [] as number[][])
  .map(([start, range]) => {
    const results: [destination: number, origin: number][] = [];

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

      results.push([destination, i]);
      i += minRange;
    }

    return results;
  })
  .flat()
  .reduce(
    (acc, result) => {
      if (result[0] < acc[0]) {
        acc = result;
      }
      return acc;
    },
    [Infinity, 0] as [destination: number, origin: number]
  );

console.log(seedRanges);
