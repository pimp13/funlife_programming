use clap::{Arg, ArgAction, ArgMatches, Command, arg};

fn generate_command() -> ArgMatches {
    Command::new("Simple Grep")
        .about("A simple grep")
        .version("v1.0.0")
        .author("pouya")
        .arg(
            Arg::new("file")
                .short('f')
                .long("file")
                .required(true)
                .action(ArgAction::Set),
        )
        .arg(
            Arg::new("pattern")
                .short('p')
                .long("pattern")
                .required(true),
        )
        .get_matches()
}

fn main() {
    let cmd = generate_command();

    let file = cmd.get_one::<String>("file").unwrap();
    let pattern = cmd.get_one::<String>("pattern").unwrap();
    println!("{} {}", file, pattern);
}
