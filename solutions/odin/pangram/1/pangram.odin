package pangram

import "core:fmt"
import "core:os"
import "core:strings"

ALPHABET : string : "abcdefghijklmnopqrstuvwxyz"

is_pangram :: proc(str: string) -> bool {
    contained, ok := strings.ascii_set_make(strings.to_lower(str))
    if !ok {
        fmt.printfln("Error parsing string %s into an ASCII set", str)
        os.exit(1)
    }
    for letter in ALPHABET {
        // TODO there should be a better way of doing this than a cast
        if !strings.ascii_set_contains(contained, u8(letter)) {
            return false
        }
    }
    return true
}
