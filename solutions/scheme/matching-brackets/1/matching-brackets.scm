(import (rnrs))

(define (balanced? string)
  (define characters (string->list string))
  (define (process characters stack)
    ; This could be so much cleaner!
    (cond ((null? characters) (null? stack))
          ((eqv? (car characters) #\( ) (process
                                         (cdr characters)
                                         (cons #\( stack)))
          ((eqv? (car characters) #\) ) (if
                                         (or
                                          (null? stack)
                                          (not (eqv? (car stack) #\()))
                                          #f
                                          (process (cdr characters) (cdr stack))))
          ((eqv? (car characters) #\[ ) (process
                                         (cdr characters)
                                         (cons #\[ stack)))
          ((eqv? (car characters) #\] ) (if
                                         (or
                                          (null? stack)
                                          (not (eqv? (car stack) #\[)))
                                          #f
                                          (process (cdr characters) (cdr stack))))
          ((eqv? (car characters) #\{ ) (process
                                         (cdr characters)
                                         (cons #\{ stack)))
          ((eqv? (car characters) #\} ) (if
                                         (or
                                          (null? stack)
                                          (not (eqv? (car stack) #\{)))
                                          #f
                                          (process (cdr characters) (cdr stack))))
          ('else (process (cdr characters) stack))
     )
    )
  (process characters '()))