module Queen8 where

safe :: Int -> [Int] -> Int -> Bool

safe _ [] _ = True
safe x (x1:xs) n =
  x /= x1
  && x /= x1 + n && x /= x1 - n
  && safe x xs (n+1)

queens :: Int -> [[Int]]
queens 0 = [[]]
queens n = [ x : y | y <- queens (n-1), x <- [1..8], safe x y 1 ]

queensN :: Int -> [[Int]]
queensN n = queens n
  where
    queens 0 = [[]]
    queens m = [ x : y | y <- queens (m-1), x <- [1..n], safe x y 1 ]
