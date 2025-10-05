(import (rnrs))

(define (leap-year? year)
  (define (is-divisible-by? n)
    (= 0 (modulo year n)))
  (cond
    ((is-divisible-by? 400) #t)
    ((is-divisible-by? 100) #f)
    ((is-divisible-by? 4) #t)
    (#t #f)))

