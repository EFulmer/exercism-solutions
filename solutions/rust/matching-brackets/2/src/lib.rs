pub fn brackets_are_balanced(string: &str) -> bool {
    let mut stack: Vec<char> = Vec::new();

    for c in string.chars() {
        match c {
            '{' | '[' | '(' => stack.push(c),
            '}' | ']' | ')' => {
                if !(handle_closing_bracket(c, stack.pop())) {
                    return false;
                }
            }
            _ => {}
        };
    }
    stack.is_empty()
}

fn handle_closing_bracket(current_char: char, last_seen: Option<char>) -> bool {
    match (current_char, last_seen) {
        (_, None) => false,
        ('}', Some('{')) => true,
        (']', Some('[')) => true,
        (')', Some('(')) => true,
        _ => false,
    }
}
