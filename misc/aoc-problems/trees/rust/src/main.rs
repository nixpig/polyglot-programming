fn get_input() -> &'static str {
    "..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#"
}

fn main() {
    println!(
        "Number of trees hit: {}",
        get_input()
            .lines()
            .enumerate()
            .flat_map(|(idx, line)| line.chars().nth(idx * 3 % line.len()))
            .filter(|&x| x == '#')
            .count()
    );
}
