#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist(first_list: &[i32], second_list: &[i32]) -> Comparison {
    // cases:
    // 1. only one list is empty => sublist of the other
    // 2. both lists have the same length => check each item, or use a library method
    // 3. one list is longer than the other => shorter list, slide it along the longer until items match, then check each item from the zip starting at index i of the longer list
    if first_list.len() > 0 && second_list.len() == 0 {
        return Comparison::Superlist;
    } else if first_list.len() == 0 && second_list.len() > 0 {
        return Comparison::Sublist;
    } else if first_list.len() == 0 && second_list.len() == 0 {
        return Comparison::Equal;
    } else {
        if first_list.len() == second_list.len() {
            if first_list == second_list {
                Comparison::Equal
            } else {
                Comparison::Unequal
            }
        } else if first_list.len() < second_list.len() {
            if is_list_sublist(first_list, second_list) {
                Comparison::Sublist
            } else {
                Comparison::Unequal
            }
        } else {
            if is_list_sublist(second_list, first_list) {
                Comparison::Superlist
            } else {
                Comparison::Unequal
            }
        }
    }
}

fn is_list_sublist(sublist: &[i32], superlist: &[i32]) -> bool {
    let m = sublist.len();
    let n = superlist.len();
    for i in 0..=n-m {
        if &superlist[i..i+m] == sublist {
            return true;
        }
    }
    false
}