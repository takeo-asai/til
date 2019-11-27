main = do
  putStrLn "Hello Haskell2"
  {-
  comment
  -}
  print (add 10 5)
  -- t@(a, b, c) = (123, 3.14, "hello") -- 中にはこのままではかけない
  print a
  print b
  print c
  print t
  print (absMaybe (Just 10))
  print (absMaybe (Just (-10)))
  print $absMaybe $Just (-111) -- $ で () を省略できる
  print (absMaybe Nothing)
  print (area 10.0)
  print ((add 10)  5)
  print (map (+ 3) [1, 2, 3, 4, 5])
  print (map (3 *) [1, 2, 3, 4, 5])
  print (map (add 4) [1, 2, 3, 4, 5])
  print $ map (\x -> x * x) [1, 2, 3, 4, 5]
  print $foldr (+) 0 [1, 2, 3] -- 1 + (2 + (3 + 0)) = 6
  print $foldr (-) 0 [1, 2, 3]  -- 1 - (2 - (3 - 0)) = 2
  print $foldl (+) 0 [1, 2, 3]  -- ((0 + 1) + 2) + 3 = 6
  print $foldl (-) 0 [1, 2, 3]  -- ((0 - 1) - 2) - 3 = -6
  print [1..5]
  print [(x, y) | x <- [1, 2], y <- [3, 4]] 

t@(a, b, c) = (123, 3.14, "hello")

add :: Int -> Int -> Int
add x y =
  x + y

absMaybe x =
  case x of
    Nothing            -> 0
    Just x | x < 0     -> -x   -- Just x にマッチし，x < 0 のとき
      | otherwise -> x    -- Just x にマッチし，x < 0 でないとき

area r = pi * square r
　where
    pi = 3.14
    square x = x * x