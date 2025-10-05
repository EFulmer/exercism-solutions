(import (rnrs))

(define (leap-year? year)
  (define (is-divisible-by? n)
    (zero? (remainder year n)))
  (cond
    ((is-divisible-by? 400) #t)
    ((is-divisible-by? 100) #f)
    ((is-divisible-by?   4) #t)
    ('else                  #f)))

