use std::fs;

use clap::{Arg, Command};

fn main() {
    let matches = Command::new("File Renamer")
        .about("Rename to standard name for your files")
        .version("v1.0.0")
        .author("pimp.puma.13@gmail.com")
        .arg(Arg::new("dir").alias("d").required(true).index(1))
        .get_matches();

    let dir = matches.get_one::<String>("dir").unwrap();
    for entry in fs::read_dir(dir).unwrap() {
        let path = entry.unwrap().path();

        if let Some(file_name) = path.file_name().and_then(|n| n.to_str()) {
            let new_name = file_name.trim().to_lowercase().replace(" ", "_");

            let new_path = path.with_file_name(new_name);

            fs::rename(&path, &new_path).unwrap();
        }
    }
}
