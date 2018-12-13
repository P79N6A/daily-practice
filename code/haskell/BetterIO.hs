module Main(main) where

import           Control.Exception (SomeException, handle)
import           System.IO         (IOMode (..), hClose, hFileSize, openFile)

handler :: SomeException -> IO (Maybe Integer)
handler _ = return Nothing

simpleFileSize :: FilePath -> IO (Maybe Integer)
simpleFileSize path = handle handler $ do
  h <- openFile path ReadMode
  size <- hFileSize h
  hClose h
  return (Just size)

main = do
  putStrLn "Input file name:"
  path <- getLine
  size <- simpleFileSize path
  print size
