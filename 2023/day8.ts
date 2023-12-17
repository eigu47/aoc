import fs from "fs";

fs.readFile("day8", "utf8", (_, data) => {
  const [directions, ...rest] = data.split("\r");

  const nodes = rest.reduce((acc, line) => {
    const [node, dir] = line.split("=").map((s) => s.match(/\w+/g));
    if (!node || !dir) return acc;

    acc[node + ""] = {
      L: dir[0],
      R: dir[1],
    };

    return acc;
  }, {} as Record<string, { L: string; R: string }>);

  let current = "AAA";
  let steps = 0;

  while (current !== "ZZZ") {
    const node = nodes[current];
    const dir = directions[steps % directions.length] as "L" | "R";
    current = node[dir];
    steps++;
  }

  console.log(steps);
});
