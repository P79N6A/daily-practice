module TestMonad where

import           Control.Monad (filterM, guard, (<=<))

justHead :: String -> Maybe Char
justHead str = do
  (x:xs) <- Just str
  return x

type KnightPos  = (Int, Int)

moveKnight :: KnightPos -> [KnightPos]
moveKnight (c, r) = do
  (c', r') <- [(c+2, r-1), (c+2, r+1), (c-2, r-1), (c-2, r+1)
              , (c+1, r-2), (c+1, r+2), (c-1, r-2), (c-1, r+2)]
  guard (c' `elem` [1..8] && r' `elem` [1..8])
  return (c', r')

isBigGang :: Int -> (Bool, String)
isBigGang x = (x > 9, "Compared gang size to 9.")

applyLog :: (a, String) -> (a -> (b, String)) -> (b, String)
applyLog (x, log) f = let (y ,newLog) = f x in (y, log ++ newLog)

powerset :: [a] -> [[a]]
powerset xs = filterM (\x -> [True, False]) xs

count :: Int
count = sum $ do
  x <- [1..10]
  y <- [1..10]
  let area = x * y
  return $ if area < 50 then 1 else 0

count2 :: Int
count2 = sum $ do
  x <- [1..10]
  y <- [x..10]
  _ <- [y..10]
  return 1
