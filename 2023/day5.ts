import fs from "fs";

const [seed, ...maps] = fs
  .readFileSync("test", "utf8")
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
    return maps.reduce((acc, map, idx) => {
      let result = acc;

      console.log({ acc, map, idx, num });
      
      map.forEach((line) => {
        const [destination, origin, range] = line!;
        if (result >= origin && result < origin + range) {
          result = result - origin + destination;
          return;
        }
      });

      return result;
    }, num);
  });
}

console.log(getDestination(seed[0]!));
