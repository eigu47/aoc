import fs from "fs";

const [seed, ...maps] = fs
  .readFileSync("day5", "utf8")
  .split("\r\n\r\n")
  .map((line) =>
    line
      .split(":")[1]
      .split("\r\n")
      .filter(Boolean)
      .map((str) => str.match(/\d+/g)?.map(Number))
  );

function getDestination(seed: number[]) {
  return seed.map((num) => {
    return maps.reduce((acc, map) => {
      for (const line of map) {
        const [destination, origin, range] = line!;
        if (acc >= origin && acc < origin + range) {
          acc = acc - origin + destination;
          break;
        }
      }

      return acc;
    }, num);
  });
}

console.log(Math.min(...getDestination(seed[0]!)));
