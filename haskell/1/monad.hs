f x = 2 * x
g x = return (2 * x)

-- x が 2 で割り切れない場合には失敗する
div2 :: Int -> Maybe Int
div2 x =
    if even x then Just (x `div` 2)
    else Nothing

div8 :: Int -> Maybe Int
div8 x = return x >>= div2 >>= div2 >>= div2



main = do
    putStrLn "--f--"
    print $fmap f (Just 5)
    print $fmap f Nothing
    print $fmap f [1, 2, 3]
    putStrLn "--g--"
    print $ Just 5 >>= g
    print $ Nothing >>= g
    print $ [1, 2, 3] >>= g
    print $ [] >>= g
    putStrLn "-----"
    print (div8 8)
    print (div8 64)
    print (div8 9)