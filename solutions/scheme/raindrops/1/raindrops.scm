(import (rnrs))

(define (convert number)
  (define number-sounds (get-sounds number))
  (if (not (null? number-sounds))
      (sounds->answer number-sounds)
      (number->string number)))
      

(define (divides? m n) (zero? (modulo m n)))
(define (divides-3? m) (divides? m 3))
(define (divides-5? m) (divides? m 5))
(define (divides-7? m) (divides? m 7))

(define (take-true bools vals)
  ; This should be tail-recursive, but I haven't gotten around to making it so yet.
    (if (null? bools)
        '()
        (if (car bools)
            (cons
             (car vals) (take-true (cdr bools) (cdr vals)))
            (take-true (cdr bools) (cdr vals)))))

(define sounds '("Pling" "Plang" "Plong"))

(define (get-sounds n)
  (take-true
   `(,(divides-3? n) ,(divides-5? n) ,(divides-7? n))
   sounds))

(define (sounds->answer lst)
      (apply string-append lst))