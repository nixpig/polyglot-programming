import path from "path";
import { Opts } from "./opts";

export enum Operation {
  Add,
  Remove,
  Print,
}

export type Config = {
  args: string[];
  operation: Operation;
  config: string;
  pwd: string;
};

function getArgs(opts: Opts): string[] {
  if (!opts.args || opts.args.length) {
    return [];
  }

  const operation = getOperation(opts);

  if (operation === Operation.Print) {
    // Print accepts either:
    // - 0 arguments (prints all)
    // - 1 argument (prints item for specificed argument)
    if (opts.args.length > 1) {
      throw new Error(`Expected 0 or 1 arguments but got ${opts.args.length}`);
    }

    return opts.args;
  }

  if (operation === Operation.Add) {
    // Add accepts 2 arguments - the name and the value
    // 3 arguments are required because the first will be the `add`
    // E.g. command add <name> <value>
    if (opts.args.length !== 3) {
      throw new Error(`Expected 2 arguments but got ${opts.args.length - 1}`);
    }
  }

  if (operation === Operation.Remove) {
    // Remove accepts 1 arguments - the name and the value
    // 2 arguments are required because the first will be the `remove`
    // E.g. command remove <name>
    if (opts.args.length !== 2) {
      throw new Error(`Expected 1 argument but got ${opts.args.length - 1}`);
    }
  }

  // Slice after the `add` or `remove` command argument
  return opts.args.slice(1);
}

function getOperation(opts: Opts): Operation {
  if (opts[0] === "add") {
    return Operation.Add;
  }

  if (opts[0] === "remove") {
    return Operation.Remove;
  }

  return Operation.Print;
}

function getConfig(opts: Opts): string {
  if (opts.config) {
    return opts.config;
  }

  const homeConfigDir =
    process.env["XDG_CONFIG_HOME"] ?? `${process.env["HOME"]}/.config`;

  if (!homeConfigDir) {
    throw new Error("Could not determine config location");
  }

  return path.join(homeConfigDir, "projector", "projector.json");
}

function getPwd(opts: Opts): string {
  return opts.pwd ?? process.cwd();
}

export default function config(opts: Opts): Config {
  return {
    args: getArgs(opts),
    operation: getOperation(opts),
    config: getConfig(opts),
    pwd: getPwd(opts),
  };
}
