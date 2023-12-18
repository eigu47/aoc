import fs from "fs";

fs.readFile("day8", "utf8", (_, data) => {
  const [directions, ...rest] = data.split("\r");
  const startNodes: string[] = [];

  const nodes = rest.reduce((acc, line) => {
    const [node, dir] = line.split("=").map((s) => s.match(/\w+/g));
    if (!node || !dir) return acc;

    acc[node[0]] = {
      L: dir[0],
      R: dir[1],
    };

    if (node[0].at(-1) === "A") {
      startNodes.push(node[0]);
    }

    return acc;
  }, {} as Record<string, { L: string; R: string }>);

  const minSteps: number[] = [];

  startNodes.forEach((current) => {
    let steps = 0;

    while (current.at(-1) !== "Z") {
      const node = nodes[current];
      const dir = directions[steps % directions.length] as "L" | "R";
      current = node[dir];
      steps++;
    }

    minSteps.push(steps);
  });

  console.log(lcm(minSteps));
});

function lcm(nums: number[]) {
  const gcd = (a: number, b: number): number => (b ? gcd(b, a % b) : a);
  return nums.reduce((acc, num) => (acc * num) / gcd(acc, num));
}
