(import (rnrs)
        (rnrs arithmetic bitwise))

(define (inc n)
  (+ 1 n))

(define (make-vector-range n)
  ; Make a vector like a range [0..n], inclusive.
  (define v (make-vector (inc n)))
  (define (acc i)
             (if
               (> i n)
               v
               (begin
                 (vector-set! v i i)
                 (acc (inc i)))))
  (acc 0))

(define (sieve n)
  (define sieve-vector (make-vector-range n))
  (let outer-loop ((outer-index 2))
    (when (<= outer-index n)
      (let inner-loop ((multiple 2))
        (when (<= (* multiple outer-index) n)
          (display multiple)
          (display "\n")
          (vector-set! sieve-vector (* multiple outer-index) #f)
          (inner-loop (inc multiple))))
      (outer-loop (inc outer-index))))
  (filter
    number?
    (cddr (vector->list sieve-vector))))
