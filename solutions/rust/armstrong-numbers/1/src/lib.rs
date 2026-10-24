pub fn is_armstrong_number(num: u32) -> bool {
    if num < 10 {
        return true;
    }

    // Storing the digits as a vector
    let mut digits: Vec<u32> = Vec::new();
    let mut acc = num;
    while acc > 0 {
        // modulo-ing the whole thing by 10 to extract it
        digits.push((acc as u32) % 10);
        // then discard that last digit by dividing by 10.
        acc /= 10;
    }

    let num_digits = digits.len() as u32;

    let sum: u32 = digits.into_iter().map(|x| x.pow(num_digits)).sum();

    sum == num
}
