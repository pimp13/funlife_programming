use std::{
    collections::HashSet,
    fs::File,
    io::{BufRead, BufReader},
};

fn main() {
    let filepath = "persian_dict_19k.txt";
    let file = File::open(filepath).expect("file is not found!");
    let reader = BufReader::new(file);

    let mut words: Vec<String> = Vec::new();
    for line_result in reader.lines() {
        let line = line_result.expect("failed to read line file");

        if let Some((word, _)) = line.trim().split_once(':') {
            words.push(word.to_string());
        }
    }

    let word_set: HashSet<String> = words.iter().cloned().collect();
    for word in &words {
        for p in vec!["من".to_string(), "ما".to_string()] {
            let man_word = format!("{}{}", p, word);
            if word_set.contains(&man_word) {
                // println!("Word: {}, ManWord: {}", word, man_word);
                println!("به {} نگو {}, {} تو نیستم!", p, word, man_word);
            }
        }
    }
}
