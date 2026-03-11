module Pangram (isPangram) where

import Data.Char

isPangram :: String -> Bool
isPangram text = containsAll lcText alphabet
  where
    containsAll txt alpha = and [ c `elem` txt | c <- alpha ]
    lcText = map toLower $ filter isLetter text

alphabet :: String
alphabet = "abcdefghijklmnopqrstuvwxyz"