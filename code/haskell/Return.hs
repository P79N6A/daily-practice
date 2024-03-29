module Return where

import           Data.Char (toUpper)

isYes :: String -> Bool
isYes inpStr = (toUpper . head $ inpStr) == 'Y'

isGreen :: IO Bool
isGreen = do
  putStrLn "Is green your favorite color?"
  inpStr <- getLine
  return (isYes inpStr)

main = do
  yn <- isGreen
  putStrLn $ show yn
