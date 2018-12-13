module Applicative where

a :: Maybe Int
a = Just 3
b :: Char
b = 'x'
replicateB :: Int -> String
replicateB = \x -> replicate x b
-- fmap replicateB a
-- Just "xxx"

applyMaybe :: Maybe (a -> b) -> Maybe a -> Maybe b
applyMaybe (Just f) (Just x) = Just $ f x
applyMaybe _ _  = Nothing

addAll :: Int -> Int -> Int -> Int
addAll x y z = x + y + z

-- ['a'..'d'] <* [1..6]
-- "aaaaaabbbbbbccccccdddddd"
-- ['a'..'d'] *> [1..6]
-- [1,2,3,4,5,6,1,2,3,4,5,6,1,2,3,4,5,6,1,2,3,4,5,6]

-- :t fmap replicate (Just 3)
-- fmap replicate (Just 3) :: Maybe (a -> [a])

-- :t fmap replicate (Just 3) <*> (Just 'x')
-- fmap replicate (Just 3) <*> (Just 'x') :: Maybe [Char]
