module BasicIO where

import           System.IO

main = do
  putStrLn "Input file name:"
  inf <- getLine
  putStrLn "Output file name:"
  outf <- getLine
  inh <- openFile inf ReadMode
  outh <- openFile outf WriteMode
  mainloop inh outh
  hClose inh
  hClose outh

mainloop :: Handle -> Handle -> IO ()
mainloop inh outh = do
  ineof <- hIsEOF inh
  if ineof then return ()
  else do inpStr <- hGetLine inh
          hPutStrLn outh inpStr
          mainloop inh outh
