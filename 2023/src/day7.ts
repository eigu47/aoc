import fs from "fs";

const cardType = {
  A: 14,
  K: 13,
  Q: 12,
  T: 10,
  "9": 9,
  "8": 8,
  "7": 7,
  "6": 6,
  "5": 5,
  "4": 4,
  "3": 3,
  "2": 2,
  J: 1,
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

    const repeated: Record<string, number> = {};
    let j = 0;

    for (let i = 0; i < cards.length; i++) {
      const card = cards[i];
      if (card === "J") {
        j++;
      } else if (cards.substring(i + 1).includes(card)) {
        repeated[card] = (repeated[card] ?? 1) + 1;
      }
    }

    let [high = 0, low] = Object.values(repeated).sort((a, b) => b - a);
    high = high === 0 ? 1 + j : high + j;
    let rank = 0;

    if (high >= 5) {
      rank = handType.five;
    } else if (high === 4) {
      rank = handType.four;
    } else if (high === 3 && low === 2) {
      rank = handType.full;
    } else if (high === 3) {
      rank = handType.three;
    } else if (high === 2 && low === 2) {
      rank = handType.two;
    } else if (high === 2) {
      rank = handType.one;
    } else {
      rank = handType.high;
    }

    return {
      cards,
      bid: +bid,
      rank,
    };
  })
  .sort((a, b) => {
    const aRank = a.rank;
    const bRank = b.rank;

    if (aRank !== bRank) return aRank - bRank;

    for (let i = 0; i < a.cards.length; i++) {
      const aIdx = cardType[a.cards[i] as keyof typeof cardType];
      const bIdx = cardType[b.cards[i] as keyof typeof cardType];
      if (aIdx !== bIdx) return aIdx - bIdx;
    }

    return 0;
  })
  .reduce((acc, cur, idx) => {
    return acc + cur.bid * (idx + 1);
  }, 0);

console.log(input);
// console.log(input.filter((hand) => hand.rank === 5));
