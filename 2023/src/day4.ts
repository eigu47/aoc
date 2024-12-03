import fs from "fs";

const input = (fs.readFileSync("day4", "utf8") as string)
  .split("\n")
  .map((line) => {
    const [winners, card] = line
      .split(":")[1]
      .split("|")
      .map((str) => str.match(/\d+/g));
    return { winners, card, count: 1 };
  })
  .reduce((acc, { winners, card, count }, idx, arr) => {
    let wins = 0;
    winners?.forEach((winner) => {
      if (card?.includes(winner)) {
        wins++;
        const next = arr[idx + wins];
        if (next) {
          next.count += count;
        }
      }
    });

    return acc + count;
  }, 0);

console.log(input);
