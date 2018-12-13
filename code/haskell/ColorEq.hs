module ColorEq where

class BasicEq a where
  isEqual    :: a -> a -> Bool
  isEqual x y = not (isNotEqual x y)

  isNotEqual :: a -> a -> Bool
  isNotEqual x y = not (isEqual x y )

data Color = Red | Green |Blue
  deriving (Read, Show, Eq, Ord)

newtype UniqueID = UniqueID Int
  deriving (Eq)


myShow value = show value
myShow2 :: (Show a) => a -> String
myShow2 = show
