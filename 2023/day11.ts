import fs from "fs";

fs.readFile("day11", "utf8", (_, data) => {
  const rows = data.split("\r\n").map((row) => row.split(""));
  const expandedRows = rows.reduce((acc, row) => {
    acc.push(row);
    if (row.every((cell) => cell === ".")) {
      acc.push(row);
    }
    return acc;
  }, [] as string[][]);

  const expanded: string[][] = expandedRows[0].reduce((acc, _, y) => {
    const column = expandedRows.map((row) => row[y]);
    const isEmpty = column.every((cell) => cell === ".");

    column.forEach((cell, x) => {
      acc[x].push(cell);
      if (isEmpty) {
        acc[x].push(cell);
      }
    });

    return acc;
  }, expandedRows.map(() => []) as string[][]);

  const galaxies = expanded.reduce((acc, row, y) => {
    row.forEach((cell, x) => {
      if (cell === "#") {
        acc.push([x, y]);
      }
    });

    return acc;
  }, [] as number[][]);

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
