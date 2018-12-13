(defun mapcar2 (fn arg)
    (if (null arg)
        nil
        (cons (funcall fn (car arg)) (mapcar2 fn (cdr arg)))))

(defun sq(x) (* x x))

(mapcar2 #'sq '(1 2 3 4 5))