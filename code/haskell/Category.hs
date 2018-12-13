module Main(main) where

-- compare ordinary list within monadic type
main = do
   let res1 = fmap (odd . (^2)) [1..10]
   let res2 = (fmap odd . fmap (^2))[1..10]
   print $ res1 == res2
