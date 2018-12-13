module Test where

import           Data.Char (digitToInt, toUpper)
import           Data.List (isPrefixOf)

data Position = Double :+ Double
distance :: Position -> Position -> Double
distance p1 p2 =
  case p1 of
    x1 :+ y1 ->
      case p2 of
        x2 :+ y2 -> sqrt ((x1-x2)^2 + (y1-y2)^2)

distance2 :: Position -> Position -> Double
distance2 (x1 :+ y1) (x2 :+ y2) =
  sqrt ((x1-x2)^2 + (y1-y2)^2)

pointA :: Position
pointA = 0 :+ 0
pointB :: Position
pointB = 3 :+ 4

(!!!) :: [a] -> Int -> Maybe a
[] !!! _ = Nothing
(x:xs) !!! 0 = Just x
(x:xs) !!! n = xs !!! (n-1)

dup :: a -> Int -> [a]
dup _ 0 = []
dup x n = x : dup x (n-1)

myDropX n xs = if n <= 0 || null xs then xs else myDropX (n-1) (tail xs)
myTakeX n xs = if n <= 0 || null xs then [] else (head xs):(myTakeX (n-1) (tail xs))

-- record style
type BookId = Int
type BookTitle = String
type BookAuthor = String
data Book = Book {
  bookID      :: BookId,
  bookTitle   :: BookTitle,
  bookAuthors :: [BookAuthor]
} deriving (Show)


zipWithAdd = zipWith (+)

addOnePow x = (\x -> x^2) $ (\x -> x + 1) $ x

data Color = Red
             | Orange
             | Yellow
             | Green
             | Blue
             | Indigo
             | Violet
               deriving (Eq, Show)

sumList (x:xs) = x + sumList xs
sumList [] = 0

third (a, b, c) = c

complicated (c, a, x:xs, 5) = (a, xs)

data Tree = Tree {
  left  :: Tree,
  right :: Tree
} deriving (Show)

amIRich money
  | money < 0 = "Broken"
  | money < 10000 = "You're not rich"
  | money < 100000 = "You're a rich guy"
  | otherwise = "You're very rich"


asInt xs = loop 0 xs
loop :: Int -> String -> Int
loop acc [] = acc
loop acc (x:xs) = let acc' = acc * 10 + digitToInt x
                  in loop acc' xs

square :: [Double] -> [Double]
square [] = []
square (x:xs) = x*x : square xs

uppercase :: String -> String
uppercase [] = []
uppercase (x:xs) = toUpper x : uppercase xs

accl :: [Int] -> Int
accl xs = foldl (+) 0 xs

reverseString :: String -> String
reverseString xs = foldl (\acc c -> c:acc) [] xs
-- foldr(\c acc -> c:acc) [] xs

owl = (.) $ (.)
dot = (.) . (.)

giveMeFive :: [a] -> [a]
giveMeFive xs = map snd $ filter (\(i, x) -> i `rem` 5 == 0) $ zip [0..] xs

myMap :: (a -> b) -> [a] -> [b]
myMap f xs = foldr ((:) . f) [] xs
-- myMap f xs = foldr (\c acc -> f c : acc) [] xs

myLength :: [a] -> Int
myLength = foldr (\_ acc -> acc + 1) 0

myMax :: (Ord a) => [a] -> a
myMax [] = error "..."
myMax (x:xs) = foldr max x xs

-- mini sub-list sum
minSubList :: (Num a, Ord a) => [a] -> Int -> a
minSubList xs m = initSum + minDiff
  where
    (initXs, shifted) = splitAt m xs
    initSum = sum initXs
    minDiff = minimum $ scanl (+) 0 $ zipWith (-) shifted xs

fact :: Integer -> Integer
fact n = product [1..n]

swap :: (a, b) -> (b, a)
swap (x, y) = (y, x)
alpha :: ((a, b), c) -> (a, (b, c))
alpha ((x, y), z) = (x, (y, z))
alpha_inv :: (a, (b, c)) -> ((a, b), c)
alpha_inv  (x, (y, z)) = ((x, y), z)

data Pair a b = Pair a b deriving(Show)

tuple = (,) "Tuple" True
triple = (,,) "Triple" True False

startsWithSymbol :: (String, String, Int) -> Bool
startsWithSymbol (name, symbol, _) = isPrefixOf symbol name

-- record
data Element = Element { name         :: String
                       , symbol       :: String
                       , atomicNumber :: Int }
                       deriving(Show)

tupleToElem :: (String, String, Int) -> Element
tupleToElem (n, s, a) = Element { name = n
                                , symbol = s
                                , atomicNumber = a }
elemToTuple :: Element -> (String, String, Int)
elemToTuple e = (name e, symbol e, atomicNumber e)

prodToSum :: (a, Either b c) -> Either (a, b) (a, c)
prodToSum (x, e) =
    case e of
      Left  y -> Left  (x, y)
      Right z -> Right (x, z)

-- Functor Composition
-- square x = x * x
mis :: Maybe [Int]
mis = Just [1, 2, 3]
-- :t (fmap . fmap) square mis
-- (fmap . fmap) square mis :: Num b => Maybe [b]


-- Bifunctor
class Bifunctor f where
    bimap :: (a -> c) -> (b -> d) -> f a b -> f c d
    bimap g h = first g . second h
    first :: (a -> c) -> f a b -> f c b
    first g = bimap g id
    second :: (b -> d) -> f a b -> f a d
    second = bimap id

instance Bifunctor (,) where
  bimap f g (x, y) = (f x, g y)

instance Bifunctor Either where
  bimap f _ (Left x)  = Left (f x)
  bimap _ g (Right y) = Right (g y)

-- 找到直角三角形
rightTriangles = [(a, b, c) | a <- [1..10], b <- [1..10], c <- [1..10], a^2 + b^2 == c^2]
