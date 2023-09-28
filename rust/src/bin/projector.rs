use clap::Parser;
// the `rust` on the line below is the name of the package, which in this case is `rust`
// because the project was created in a directory called `rust`
use rust::opts::Opts;

fn main() {
    let opts = Opts::parse();

    println!("options: {:?}", opts);
}
