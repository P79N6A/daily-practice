module Main (main) where

-- import           SimpleJSON
--
-- main = print (JObject [("foo", JNumber 3.1415), ("bar", JBool False)])

main :: IO()
main = do
  x <- return 1
  y <- return 2
  if x + y < 10
    then print $ x * y
    else print $ ""
