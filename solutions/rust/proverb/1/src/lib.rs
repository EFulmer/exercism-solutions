pub fn build_proverb(list: &[&str]) -> String {
    let Some(first) = list.first() else {
        return String::new();
    };
    list.windows(2)
        .map(|pair| format!("For want of a {} the {} was lost.", pair[0], pair[1]))
        .chain(std::iter::once(format!("And all for the want of a {first}.")))
        .collect::<Vec<String>>()
        .join("\n")
}
