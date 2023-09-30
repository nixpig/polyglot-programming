use std::path::PathBuf;

use anyhow::{anyhow, Context, Ok, Result};

use crate::opts::Opts;

#[derive(Debug)]
pub struct Config {
    pub operation: Operation,
    pub pwd: PathBuf,
    pub config: PathBuf,
}

#[derive(Debug)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String),
    Remove(String),
}

impl TryFrom<Opts> for Config {
    type Error = anyhow::Error;

    fn try_from(value: Opts) -> Result<Self> {
        let operation = value.args.try_into()?;
        let config = get_config(value.config)?;
        let pwd = get_pwd(value.pwd)?;

        return Ok(Config {
            operation,
            config,
            pwd,
        });
    }
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;

    fn try_from(value: Vec<String>) -> Result<Self, Self::Error> {
        let mut value = value;

        if value.is_empty() {
            return Ok(Operation::Print(None));
        }

        let term = value.get(0).expect("expect to exist");

        if term == "add" {
            if value.len() != 3 {
                let err = anyhow!(
                    "operation add expects 2 arguments but got {}",
                    value.len() - 1
                );

                return Err(err);
            }

            let mut drain = value.drain(1..=2);

            return Ok(Operation::Add(
                drain.next().expect("to exist"),
                drain.next().expect("to exist"),
            ));
        }

        if term == "remove" {
            if value.len() != 2 {
                let err = anyhow!(
                    "operation remove expects 1 argument but got {}",
                    value.len() - 1
                );

                return Err(err);
            }

            let arg = value.pop().expect("to exist");

            return Ok(Operation::Remove(arg));
        }

        if value.len() > 1 {
            let err = anyhow!(
                "operation print expects 0 or 1 arguments but got {}",
                value.len()
            );

            return Err(err);
        }

        let arg = value.pop().expect("to exist");

        Ok(Operation::Print(Some(arg)))
    }
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(config) = config {
        return Ok(config);
    }

    let home_config_path =
        std::env::var("HOME").context("unable to get HOME environment variable")?;
    let mut home_config_path = PathBuf::from(home_config_path);

    home_config_path.push(".config");
    home_config_path.push("projector");
    home_config_path.push("projector.json");

    return Ok(home_config_path);
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(pwd) = pwd {
        return Ok(pwd);
    }

    let working_dir = std::env::current_dir().context("error getting present working directory")?;

    return Ok(working_dir);
}
