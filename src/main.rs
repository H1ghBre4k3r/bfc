mod lexing;
mod util;

use std::process;

use clap::{IntoApp, Parser};
use lexing::lexer;

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

    let content = match util::read_file(&args.file) {
        Ok(content) => content,
        Err(_) => {
            eprintln!("Could not open file '{}'!", &args.file);
            process::exit(-1);
        }
    };
    println!("{:?}", lexer::lex(&content, &args.file));
}
