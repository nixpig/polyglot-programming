import { parse } from "path";

function getInput(): string {
  return `forward 5
down 5
forward 8
up 3
down 8
forward 2`;
}

type Direction = "forward" | "up" | "down";

type Coordinates = [x: number, y: number];

function parseLine(line: string): Coordinates {
  const direction: Direction = line.split(" ")[0] as Direction;
  const amount: number = +line.split(" ")[1];

  if (direction === "forward") {
    return [amount, 0];
  } else if (direction === "up") {
    return [0, -amount];
  }

  return [0, amount];
}

const out = getInput()
  .split("\n")
  .map((x) => parseLine(x))
  .reduce(
    (acc, amount) => {
      acc[0] += amount[0];
      acc[1] += amount[1];
      return acc;
    },
    [0, 0]
  );

console.log(out, out[0] * out[1]);
