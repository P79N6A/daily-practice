module TestStateMonad where

type Stack = [Int]

newtype State s a = State { runState :: s -> (a, s) }

instance Monad (State s) where
  return x = State $ \s -> (x, s)
  (State h) >>= f = State $ \s -> let (a, newState) = h s
                                      (State g) = f a
                                  in g newState

pop :: State Stack Int
pop = State $ \(x:xs) -> (x, xs)

push :: Int -> State Stack ()
push a = State $ \xs -> ((), a:xs)

stackManip :: State Stack Int
stackManip = do
  push 3
  a <- pop
  pop

stackyStack :: State Stack ()
stackyStack = do
  stackNow <- get
  if stackNow == [1,2,3]
      then put [8,3,1]
      else put [9,2,1]
