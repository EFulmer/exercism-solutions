const LIMIT: u64 = 1_000_000_000;

pub fn collatz(n: u64) -> Option<u64> {
    if n == 0 {
        return None;
    }
    let mut steps = 0;
    let mut n = n;
    while n > 1 && steps < LIMIT {
        n = formula(n);
        steps += 1;
    }
    if steps >= LIMIT { // Just in case ... :)
        None
    } else {
        Some(steps)
    }
}

fn formula(n: u64) -> u64 {
    match n % 2 {
        0 => n / 2,
        _ => (n * 3) + 1
    }
}
