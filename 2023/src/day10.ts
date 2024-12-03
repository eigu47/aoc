import fs from "fs";

type Tile = [y: number, x: number];

const pipesMap: Record<string, Tile[]> = {
  S: [
    [0, 1],
    [1, 0],
  ],
  "|": [
    [-1, 0],
    [1, 0],
  ],
  "-": [
    [0, -1],
    [0, 1],
  ],
  L: [
    [-1, 0],
    [0, 1],
  ],
  J: [
    [-1, 0],
    [0, -1],
  ],
  "7": [
    [1, 0],
    [0, -1],
  ],
  F: [
    [1, 0],
    [0, 1],
  ],
};

fs.readFile("day10", "utf8", (_, data) => {
  const map = data.split("\r\n");

  const animal: Tile = [0, 0];
  for (let y = 0; y < map.length; y++) {
    const x = map[y].indexOf("S");
    if (x !== -1) {
      animal[0] = y;
      animal[1] = x;
      break;
    }
  }

  const path = new Set([animal.toString()]);
  let cur = animal;

  while (true) {
    const pipe = map[cur[0]][cur[1]];
    const [dir1, dir2] = pipesMap[pipe];

    const next1: Tile = [cur[0] + dir1[0], cur[1] + dir1[1]];
    if (!path.has(next1.toString())) {
      path.add(next1.toString());
      cur = next1;
      continue;
    }

    const next2: Tile = [cur[0] + dir2[0], cur[1] + dir2[1]];
    if (!path.has(next2.toString())) {
      path.add(next2.toString());
      cur = next2;
      continue;
    }

    // console.log(path.size / 2);
    break;
  }

  const enclosed = map.reduce((acc, line, y) => {
    let crosses = 0;

    for (let x = 0; x < map[y].length; x++) {
      if (path.has([y, x].toString())) {
        if (line[x] === "|" || line[x] === "7" || line[x] === "F") {
          crosses++;
        }
        continue;
      }

      if (crosses % 2 === 1) {
        acc++;
      }
    }

    return acc;
  }, 0);

  console.log(enclosed);
});
