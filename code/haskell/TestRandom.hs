module TestRandom where

import           Control.Monad.State
import           System.Random

randomStr :: (RandomGen g, Random a) => State g a
randomStr = state random
