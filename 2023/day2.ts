const fs = require("fs");

const elfCubes: Record<"red" | "green" | "blue", number> = {
  red: 12,
  green: 13,
  blue: 14,
};

const input = (fs.readFileSync("day2", "utf8") as string)
  .split("\n")
  .map((line) => {
    const [_, game] = line.split(":").map((str) => str.trim());
    return game.split(";").map((set) =>
      set
        .split(",")
        .map((str) => str.trim().split(" "))
        .reduce((acc, [amount, color]) => {
          acc[color] = +amount;
          return acc;
        }, {} as Record<string, number>)
    );
  })
  .reduce((acc, game, i) => {
    let isPossible = true;
    for (const set of game) {
      for (const [color, amount] of Object.entries(elfCubes)) {
        if (set[color] > amount) {
          isPossible = false;
        }
      }
      if (!isPossible) {
        break;
      }
    }

    return isPossible ? i + 1 + acc : acc;
  }, 0);

console.log(input);
