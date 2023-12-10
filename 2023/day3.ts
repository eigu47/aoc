import fs from "fs";

const input = (fs.readFileSync("day3", "utf8") as string)
  .split("\n")
  .map((line, idx, arr) => {
    line = line.trim();

    const nums: number[] = [];

    for (let i = 0; i < line.length; i++) {
      if (!isNaN(+line[i])) {
        let num = line[i];

        while (++i < line.length && !isNaN(+line[i])) {
          num += line[i];
        }

        const top =
          arr[idx - 1]?.substring(i - num.length - 1, i + 1).trim() ?? "";
        const lft = line[i - num.length - 1] ?? "";
        const rgt = line[i] ?? "";
        const btm =
          arr[idx + 1]?.substring(i - num.length - 1, i + 1).trim() ?? "";

        const adjacent = top + lft + rgt + btm;

        if (/[^0-9.]/.test(adjacent)) {
          console.log({ num, adjacent });
          nums.push(+num);
        }
      }
    }

    return nums;
  })
  .flat()
  .reduce((acc, cur) => acc + cur, 0);

console.log(input);
