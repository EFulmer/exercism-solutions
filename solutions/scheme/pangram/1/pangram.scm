(import (rnrs))

(define (pangram? phrase)
  (define alphabet
    (string->list "abcdefghijklmnopqrstuvwxyz"))
  
  (define phrase-letters 
    (map char-downcase (string->list phrase)))
  
  (define (phrase-contains? v)
    ; (memv v lst) returns the rest of lst, starting with the first occurence of v, if v is in lst, and false/#f otherwise.
    (list? (memv v phrase-letters)))
  
  (for-all phrase-contains? alphabet))