import fs from "fs";

type Hand = [cards: string, bid: number];

const cardType = {
  A: 14,
  K: 13,
  Q: 12,
  J: 11,
  T: 10,
  "9": 9,
  "8": 8,
  "7": 7,
  "6": 6,
  "5": 5,
  "4": 4,
  "3": 3,
  "2": 2,
};

const handType = {
  five: 6,
  four: 5,
  full: 4,
  three: 3,
  two: 2,
  one: 1,
  high: 0,
};

const input = fs
  .readFileSync("day7", "utf8")
  .split("\r\n")
  .map((line) => {
    const [cards, bid] = line.split(" ");
    return [cards, +bid] as Hand;
  })
  .sort((a, b) => {
    const aRank = getRank(a[0]);
    const bRank = getRank(b[0]);

    if (aRank !== bRank) return aRank - bRank;

    for (let i = 0; i < a[0].length; i++) {
      const aIdx = cardType[a[0][i] as keyof typeof cardType];
      const bIdx = cardType[b[0][i] as keyof typeof cardType];
      if (aIdx !== bIdx) return aIdx - bIdx;
    }
    return 0;
  })
  .reduce((acc, cur, idx) => {
    return acc + cur[1] * (idx + 1);
  }, 0);

function getRank(cards: string) {
  const repeated: Record<string, number> = {};

  for (let i = 0; i < cards.length; i++) {
    const card = cards[i];
    if (cards.substring(i + 1).includes(card)) {
      repeated[card] = (repeated[card] ?? 1) + 1;
    }
  }

  const [high, low] = Object.values(repeated).sort((a, b) => b - a);

  if (high === 5) {
    return handType.five;
  } else if (high === 4) {
    return handType.four;
  } else if (high === 3 && low === 2) {
    return handType.full;
  } else if (high === 3) {
    return handType.three;
  } else if (high === 2 && low === 2) {
    return handType.two;
  } else if (high === 2) {
    return handType.one;
  } else {
    return handType.high;
  }
}

console.log(input);
