const fs = require("fs");

const map = {
  "0": "0",
  "1": "1",
  "2": "2",
  "3": "3",
  "4": "4",
  "5": "5",
  "6": "6",
  "7": "7",
  "8": "8",
  "9": "9",
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight: "8",
  nine: "9",
};

const input = (fs.readFileSync("day1", "utf8") as string)
  .split("\n")
  .map((line) => {
    let nums: string[] = [];

    Object.entries(map).forEach(([key, value]) => {
      const first = line.indexOf(key);
      if (first > -1) {
        nums[first] = value;
      }

      const last = line.lastIndexOf(key);
      if (last > -1) {
        nums[last] = value;
      }
    });

    nums = nums.filter((num) => num !== undefined);
    return (nums.at(0) ?? "") + (nums.at(-1) ?? "");
  })
  .reduce((acc, cur) => acc + +cur, 0);

console.log(input);

export {};
