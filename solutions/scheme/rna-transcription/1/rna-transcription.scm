(import (rnrs))



(define (dna->rna dna)
  (define translations-chars
    '((#\C . #\G)
      (#\G . #\C)
      (#\T . #\A)
      (#\A . #\U)))
  (define (translate-one-char c)
    (cdr (assoc c translations-chars)))
  (string-map translate-one-char dna))

