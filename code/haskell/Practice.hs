module Practice where

import           Data.Char   (toUpper)
import           Data.String (words)

type Writer a = (a, String)

(>=>) :: (a -> Writer b) -> (b -> Writer c) -> (a -> Writer c)

m1 >=> m2 = \x -> 
    let (y, s1) = m1 x
        (z, s2) = m2 y
    in (z, s1 ++ s2)

upCase :: String -> Writer String
upCase s = (map toUpper s, "upCase ")

toWords :: String -> Writer [String]
toWords s = (words s, "toWords ")

process :: String -> Writer [String]
process = upCase >=> toWords