use clap::Parser;

fn main() {
    // the `rust` on the line below is the name of the package, which in this case is `rust`
    // because the project was created in a directory called `rust`
    let opts = rust::opts::Opts::parse();

    println!("options: {:?}", opts);
}
