module Main(main) where

import           System.IO

main = do
  putStrLn "Input file name:"
  inf <- getLine
  putStrLn "Output file name:"
  outf <- getLine
  putStrLn "PIPE..."
  inpStr <- readFile inf
  writeFile outf inpStr
  putStrLn "DONE."
