import os

fn main() {
	
	input := os.args[1]
	number := input.int()

	if number %2 == 0 {
		println('${number} is zooj')
	} else {
		println('${number} is faard')
	}
}