module PrettyJSON
    (
      renderJValue
    ) where

import           Data.List  (intercalate)
import           Prettify   (Doc, char, compact, double, fsep, hcat, pretty,
                             punctuate, text, (<>))
import           SimpleJSON

renderJValue :: JValue -> Doc
renderJValue (JBool True)  = text "true"
renderJValue (JBool False) = text "false"
renderJValue JNull         = text "null"
renderJValue (JNumber num) = double num
renderJValue (JString str) = ptext str

renderJValue (JObject obj) = series '{' '}' field obj
    where
      field (name,val) = string name
      <> text ": "
      <> renderJValue val

renderJValue (JArray ary) = series '[' ']' renderJValue ary

putJValue :: JValue -> IO ()
putJValue v = putStrLn (render (renderJValue v))
