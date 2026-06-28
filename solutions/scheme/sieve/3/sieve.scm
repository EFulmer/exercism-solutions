(import (rnrs)
        (rnrs arithmetic bitwise))

(define (inc n)
  (+ 1 n))

(define (make-vector-range n)
  ; Make a vector like a range [0..n], inclusive.
  (let ((v (make-vector (inc n))))
    (do
      ((i 0 (+ i 1))) ((> i n) v)
      (vector-set! v i i))))

(define (sieve n)
  (define sieve-vector (make-vector-range n))
  (let outer-loop ((outer-index 2))
    (when (<= outer-index n)
      (let inner-loop ((multiple 2))
        (when (<= (* multiple outer-index) n)
          (vector-set! sieve-vector (* multiple outer-index) #f)
          (inner-loop (inc multiple))))
      (outer-loop (inc outer-index))))
  (filter
    (lambda (n) (not (boolean? n)))
    (cddr (vector->list sieve-vector))))
