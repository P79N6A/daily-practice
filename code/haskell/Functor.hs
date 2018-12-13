module Functor where

class FunctorNew f where
  fmap :: (a -> b) -> f a -> f b

instance FunctorNew [] where
  fmap = map

instance FunctorNew Maybe where
  fmap f (Just a) = Just (f a)
  fmap _ Nothing = Nothing
