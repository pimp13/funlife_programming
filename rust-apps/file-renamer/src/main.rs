use std::fs;

fn main() {
    let dir = "./test-file";

    for entry in fs::read_dir(dir).unwrap() {
        let path = entry.unwrap().path();

        if let Some(file_name) = path.file_name().and_then(|n| n.to_str()) {
            let new_name = file_name.trim().to_lowercase().replace(" ", "_");

            println!("filename: {:?} new filename: {}", file_name, new_name);
        }
    }
}
