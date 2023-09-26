fn get_input() -> &'static str {
    "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn parse_line(line: &str) -> Point {
    let (dir, amount) = line
        .split_once(' ')
        .expect("This should always be a whitespace!");

    let amount = str::parse::<i32>(amount).expect("Second arg must be an integer");

    if dir == "forward" {
        Point { x: amount, y: 0 }
    } else if dir == "up" {
        Point { x: 0, y: -amount }
    } else {
        Point { x: 0, y: amount }
    }
}

fn main() {
    let result =
        get_input()
            .lines()
            .map(parse_line)
            .fold(Point { x: 0, y: 0 }, |mut acc, point| {
                acc.x += point.x;
                acc.y += point.y;

                acc
            });

    println!("Result: {:?}", result);
}
