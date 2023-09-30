use anyhow::Result;
use clap::Parser;
// the `rust` on the line below is the name of the package, which in this case is `rust`
// because the project was created in a directory called `rust`
use rust::{config::Config, opts::Opts};

fn main() -> Result<()> {
    let opts: Config = Opts::parse().try_into()?;

    println!("options: {:?}", opts);

    return Ok(());
}
