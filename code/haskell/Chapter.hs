module Chapter where

safeHead :: [x] -> Maybe x
safeHead (x:_) = Just x
safeHead [] = Nothing

suffixes :: [a] -> [[a]]
suffixes xs@(_:xs') = xs : suffixes xs'
suffixes [] = []

myFoldl :: (a -> b -> a) -> a -> [b] -> a
myFoldl f a xs = foldr (\c acc x -> acc (f x c)) id xs a

-- instance of Eq
data HolyCrap = Holy Double | Crap Int
instance Eq HolyCrap where
  Holy x == Crap y = truncate(x) == y

-- quick sort
qsort :: Ord a => [a] -> [a]
qsort [] = []
qsort (x:xs) = qsort (filter (<x) xs) ++ [x] ++ (filter (>x) xs)

-- coerce compute
foldl2 :: (b -> a -> b) -> b -> [a] -> b
foldl2 _ acc [] = acc
foldl2 f acc (x:xs) =
  let acc' = f acc x
  in seq acc' (foldl2 f acc' xs)
