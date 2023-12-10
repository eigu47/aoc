import fs from "fs";

const input = (fs.readFileSync("day3", "utf8") as string)
  .split("\n")
  .map((line, idx, arr) =>
    [...line.trim().matchAll(/(\d+)/g)].map((num) => {
      const top =
        arr[idx - 1]
          ?.substring(num.index! - 1, num.index! + num[0].length + 1)
          .trim() ?? "";
      const lft = line[num.index! - 1]?.trim() ?? "";
      const rgt = line[num.index! + num[0].length]?.trim() ?? "";
      const btm =
        arr[idx + 1]
          ?.substring(num.index! - 1, num.index! + num[0].length + 1)
          .trim() ?? "";

      if (/[^0-9.]/.test(top + lft + rgt + btm)) {
        return +num[0];
      }
    })
  )
  .flat()
  .reduce((acc, cur) => acc! + (cur ?? 0), 0);

console.log(input);
