import fs from "fs";

type Hand = [cards: string, bid: number];

const input: Hand[] = fs
  .readFileSync("day7", "utf8")
  .split("\r\n")
  .map((line) => {
    const [cards, bid] = line.split(" ");
    return [cards, +bid];
  });

const handType: [
  high: Hand[],
  one: Hand[],
  two: Hand[],
  three: Hand[],
  full: Hand[],
  four: Hand[],
  five: Hand[]
] = [[], [], [], [], [], [], []];

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

input.forEach((hand) => {
  const repeated: Record<string, number> = {};
  const [cards, _] = hand;
  for (let i = 0; i < cards.length; i++) {
    const card = cards[i];
    if (cards.substring(i + 1).includes(card)) {
      repeated[card] = (repeated[card] ?? 1) + 1;
    }
  }

  const values = Object.values(repeated);

  if (values[0] === 5) {
    handType[6].push(hand);
  } else if (values[0] === 4) {
    handType[5].push(hand);
  } else if (values.includes(2) && values.includes(3)) {
    handType[4].push(hand);
  } else if (values[0] === 3) {
    handType[3].push(hand);
  } else if (values[0] === 2 && values[1] === 2) {
    handType[2].push(hand);
  } else if (values[0] === 2) {
    handType[1].push(hand);
  } else {
    handType[0].push(hand);
  }
});

handType.forEach((hands) =>
  hands.sort((a, b) => {
    for (let i = 0; i < a[0].length; i++) {
      const aIdx = cardType[a[0][i] as keyof typeof cardType];
      const bIdx = cardType[b[0][i] as keyof typeof cardType];
      if (aIdx !== bIdx) return aIdx - bIdx;
    }
    return 0;
  })
);

console.log(
  handType.flat().reduce((acc, cur, idx) => {
    return acc + cur[1] * (idx + 1);
  }, 0)
);
