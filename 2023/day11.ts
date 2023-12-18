import fs from "fs";

fs.readFile("day11", "utf8", (_, data) => {
  const universe = data.split("\r\n").map((row) => row.split(""));

  const galaxies = universe.reduce((acc, row, y) => {
    row.forEach((cell, x) => {
      if (cell === "#") {
        acc.push([x, y]);
      }
    });

    return acc;
  }, [] as number[][]);

  [...universe].reverse().forEach((row, idx) => {
    idx = universe[0].length - 1 - idx;
    if (row.every((cell) => cell === ".")) {
      galaxies.forEach((galaxy) => {
        if (galaxy[1] > idx) {
          galaxy[1] += 999_999;
        }
      });
    }
  });

  [...universe[0]].reverse().forEach((_, idx) => {
    idx = universe[0].length - 1 - idx;
    const column = universe.map((row) => row[idx]);
    if (column.every((cell) => cell === ".")) {
      galaxies.forEach((galaxy) => {
        if (galaxy[0] > idx) {
          galaxy[0] += 999_999;
        }
      });
    }
  });

  let total = 0;
  for (let i = 0; i < galaxies.length; i++) {
    for (let j = i + 1; j < galaxies.length; j++) {
      total +=
        Math.abs(galaxies[i][0] - galaxies[j][0]) +
        Math.abs(galaxies[i][1] - galaxies[j][1]);
    }
  }

  console.log(total);
});
