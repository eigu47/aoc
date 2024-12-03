import fs from "fs";

const [times, distances] = fs
  .readFileSync("day6", "utf8")
  .split("\n")
  .map((line) => [+line.split(":")[1].replace(/\D/g, "")]);

const result = times.reduce((result, time, idx) => {
  const distance = distances[idx];
  let wins = 0;

  for (let t = 1; t < time; t++) {
    if (t * (time - t) > distance) {
      wins++;
    }
  }

  return result * wins;
}, 1);

console.log(result);
