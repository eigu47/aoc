const fs = require("fs");

const input = (fs.readFileSync("day1-1", "utf8") as string)
  .split("\n")
  .reduce((sum, line) => {
    const nums = line
      .trim()
      .split("")
      .filter((c) => !isNaN(c as any));
    return sum + +((nums.at(0) ?? "") + (nums.at(-1) ?? ""));
  }, 0);

console.log(input);
