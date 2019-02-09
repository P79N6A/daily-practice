module CoMonad where

type Option = String
data Config = MakeConfig [Option] deriving (Show)

configBuilder :: [Option] -> Config
configBuilder options = MakeConfig options

defaultConfig :: [Option] -> Config
defaultConfig options = MakeConfig (["-Wall"] ++ options)

profile :: ([Option] -> Config) -> Config
profile builder = builder ["-prof", "-auto-all"]

goFaster :: ([Option] -> Config) -> Config
goFaster builder = builder ["-O2"]

profile'  :: ([Option] -> Config) -> ([Option] -> Config)
profile' builder =
    \options -> builder (["-prof", "-auto-all"] ++ options)

goFaster' :: ([Option] -> Config) -> ([Option] -> Config)
goFaster' builder =
    \options -> builder (["-O2"] ++ options)

extract :: ([Option] -> Config) -> Config
extract builder = builder []

(#) :: a -> (a -> b) -> b
x # f = f x

extend :: (([Option] -> Config) ->              Config )
       ->  ([Option] -> Config) -> ([Option] -> Config)
extend setter builder =
        \opts2 -> setter (\opts1 -> builder (opts1 ++ opts2))
