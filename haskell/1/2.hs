data Shape =
    Rect Double Double
    | Tri Double Double

area :: Shape -> Double
area (Rect x y) = x * y
area (Tri x y) = x * y / 2

data Person =
    Person { name :: String, age :: Int }

inc :: Person -> Person
inc p = p { age = age p + 1 }

data Point = Pt Double Double

instance Eq Point where
  (Pt x y) == (Pt x' y') = x == x' && y == y'

main = do
    print (area (Rect 3.0 2.0))
    print (area (Tri 3 2))
    let taro = Person { name = "Taro", age = 25 }
    print (age taro)
    print (age (inc taro))
    print (age taro)
    print (name taro)
    print $ (Pt 1 2) == (Pt 2 3)
    print $ (Pt 1 2) /= (Pt 2 3)
    print $ (Pt 1 2) == (Pt 1 2)

