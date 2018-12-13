module GetLine(main)where

import           Data.Char
import           Data.List

main = do
  putStrLn "Say hi be here:"
  line <- fmap (intersperse '-' . reverse . map toUpper) getLine
  putStrLn line
