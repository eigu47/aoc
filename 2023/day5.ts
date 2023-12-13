import fs from "fs";

const input = fs
  .readFileSync("day5", "utf8")
  .split("\r\n\r\n")
  .map((line) =>
    line
      .split(":")[1]
      .split("\r\n")
      .filter(Boolean)
      .map((str) => str.match(/\d+/g)?.map(Number) ?? [])
  );

const result = input[0][0].map((num) =>
  input.slice(1).reduce((acc, map) => {
    for (const [destination, origin, range] of map) {
      if (acc >= origin && acc < origin + range) {
        acc = acc - origin + destination;
        break;
      }
    }
    return acc;
  }, num)
);

console.log(Math.min(...result));
