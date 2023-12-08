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
  .reduce((acc, game) => {
    let min = {
      red: 0,
      green: 0,
      blue: 0,
    };

    game.forEach((set) => {
      Object.entries(set).forEach(([color, amount]) => {
        if (amount > min[color as keyof typeof min]) {
          min[color as keyof typeof min] = amount;
        }
      });
    });

    return acc + Object.values(min).reduce((acc, val) => acc * val, 1);
  }, 0);

console.log(input);
