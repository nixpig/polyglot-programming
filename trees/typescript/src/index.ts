function getInput(): string {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

let treesHit = 0;

getInput()
  .split("\n")
  .forEach((row, i) => {
    if (row[(i * 3) % row.length] === "#") {
      treesHit++;
    }
  });

console.log(`Number of trees hit: ${treesHit}`);
