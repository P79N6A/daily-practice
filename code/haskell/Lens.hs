module Main(main) where

data Position = Position { positionX :: Double, positionY :: Double } deriving Show
data Line = LIne { lineStart :: Position, lineEnd :: Position } deriving Show
setPositionX x' p = p { positionX = x' }
xLens f p = fmap (\x' -> setPositionX x' p) $ f (positionX p)

view :: xLens Position Double -> Position -> Double
set :: xLens Position Double -> Double -> Position -> Position
over :: xLens Position Double -> (Double -> Double) -> Position -> Position

view = undefined
set = undefined
over = undefined

main = do
  print $ xLens (\x -> Just (x+1)) (Position 3 4)
  print $ xLens (\x -> [x+1, x+2, x+3]) (Position 3 4)
