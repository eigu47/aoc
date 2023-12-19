import fs from "fs";

const dirMap = {
  U: [1, 0],
  D: [-1, 0],
  L: [0, -1],
  R: [0, 1],
};

const input = fs
  .readFileSync("day18", "utf8")
  .split("\r\n")
  .map((line) => line.split(" "))
  .reduce(
    (acc, line) => {
      const [dir, length] = line;
      const [dirY, dirX] = dirMap[dir as keyof typeof dirMap];

      for (let i = 0; i < +length; i++) {
        const [y, x] = acc.at(-1)!;
        const newDir = [y + dirY, x + dirX, dir] as [number, number, string];
        acc.push(newDir);
      }

      return acc;
    },
    [[0, 0, "S"]] as [y: number, x: number, dir: string][]
  );

const minValues = input.reduce(
  (acc, [y, x]) => {
    acc[0] = Math.min(acc[0], y);
    acc[1] = Math.min(acc[1], x);
    return acc;
  },
  [0, 0] as [minY: number, minX: number]
);

let volume = -1;

const map = input.reduce((acc, [y, x, dir]) => {
  const adjustedY = y - minValues[0];
  if (!acc[adjustedY]) {
    acc[adjustedY] = [];
  }

  acc[adjustedY][x - minValues[1]] = dir;
  volume++;
  return acc;
}, [] as string[][]);

for (let y = 0; y < map.length; y++) {
  const row = map[y];
  let cross = 0;

  for (let x = 0; x < row.length; x++) {
    if (map[y][x] === "U" || map[y][x] === "D") {
      cross++;
      continue;
    }

    if (map[y][x] === undefined && cross % 2 === 1) {
      volume++;
    }
  }
}

console.log({ volume });
