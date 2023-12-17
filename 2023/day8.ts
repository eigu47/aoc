import fs from "fs";

const [instructions, rest] = fs.readFileSync("day8", "utf8").split("\r\n\r\n");

const nodes = rest.split("\r\n").reduce((acc, line) => {
  const [node, dir] = line.split("=").map((s) => s.match(/\w+/g));
  if (!node || !dir) return acc;

  acc[node + ""] = {
    L: dir[0],
    R: dir[1],
  };

  return acc;
}, {} as Record<string, { L: string; R: string }>);

function getNode(key = "AAA", steps = 0) {
  if (key === "ZZZ") return steps;

  const node = nodes[key];
  const dir = instructions[steps % instructions.length] as "L" | "R";
  const next = node[dir];

  console.log({ key, dir, next, steps });

  return getNode(next, steps + 1);
}

console.log(getNode());
