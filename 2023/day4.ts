import fs from "fs";

const input = (fs.readFileSync("day4", "utf8") as string)
  .split("\n")
  .map((line) => [...line.split(":")[1].split("|"), "1"])
  .reduce((acc, cur) => {
    const winners = cur[0].match(/\d+/g);
    const card = cur[1].match(/\d+/g);

    const wins = winners?.reduce((acc, cur) => {
      if (card?.some((c) => c == cur)) {
        return !acc ? 1 : (acc *= 2);
      }
      return acc;
    }, 0);

    return acc + (wins ?? 0);
  }, 0);

console.log(input);
