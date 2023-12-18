import fs from "fs";

type Tile = [row: number, col: number];

const pipesMap: Record<string, Tile[]> = {
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

const input = fs.readFile("day10", "utf8", (error, data) => {
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

  let cur: Tile = [animal[0], animal[1] + 1];
  const visited = new Set([animal.toString(), cur.toString()]);

  while (true) {
    const pipe = map[cur[0]][cur[1]];
    const [dir1, dir2] = pipesMap[pipe];

    const next1: Tile = [cur[0] + dir1[0], cur[1] + dir1[1]];
    if (!visited.has(next1.toString())) {
      visited.add(next1.toString());
      cur = next1;
      continue;
    }

    const next2: Tile = [cur[0] + dir2[0], cur[1] + dir2[1]];
    if (!visited.has(next2.toString())) {
      visited.add(next2.toString());
      cur = next2;
      continue;
    }

    if (map[next1[0]][next1[1]] === "S" || map[next2[0]][next2[1]] === "S") {
      break;
    }
  }

  console.log(visited.size / 2);
});
