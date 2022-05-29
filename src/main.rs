mod interpreter;
mod lexing;
mod parsing;
mod util;

use clap::{IntoApp, Parser};
use interpreter::interpret;

#[derive(Parser, Debug)]
#[clap(name = "BFC")]
#[clap(author, about, version)]
struct Args {
    /// input file for interpreter/compiler
    #[clap(short, long)]
    file: String,

    /// flag for compiling instead of interpreting
    #[clap(short, long)]
    compile: bool,

    /// output directory for compilation files
    #[clap(short, long)]
    output: Option<String>,
}

fn main() {
    let args = Args::parse();

    if args.compile {
        if let None = args.output.as_deref() {
            let mut cmd = Args::command();
            cmd.error(
                clap::ErrorKind::ArgumentConflict,
                "output path must be present when trying to compile!",
            )
            .exit();
        }
    }

    if args.compile {
    } else {
        interpret(&args.file);
    }
}
