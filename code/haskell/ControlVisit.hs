module Main(main) where

import           System.Directory (Permissions)
import           System.IO        (FilePath)

data Info = Info {
      infoPath    :: FilePath
    , infoPerms   :: Maybe Permissions
    , infoSize    :: Maybe Integer
    , infoModTime :: Maybe ClockTime
    } deriving (Eq, Ord, Show)
